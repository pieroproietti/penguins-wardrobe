#!/bin/bash
# ==============================================================================
# finish.sh - pasos finales específicos de Quirinux para Calamares/coa
# Adaptado de fin-instalacion (Javier Obregón / Charlie Martínez)
#
# Este script corre DENTRO del chroot del sistema recién instalado, vía el
# módulo shellprocess de Calamares (dontChroot: false) -- no hace falta
# detectar ni prefijar ningún $CHROOT como en el fin-instalacion original.
#
# Se ejecuta ANTES de shellprocess@oa-chroot-runner en la secuencia, así el
# update-grub genérico que ese paso ya corre para toda la familia Debian/
# Devuan recoge el cambio de GRUB_DISABLE_OS_PROBER de más abajo sin
# necesidad de llamarlo de nuevo acá.
# ==============================================================================
set -e

# --- Renombrar referencias al usuario plantilla "quirinux" -----------------
# Los archivos de skel de Quirinux traen paths/config hardcodeados para un
# usuario "quirinux"; hay que reemplazarlos por el usuario real que
# Calamares acaba de crear. Se excluye "live" porque a esta altura de la
# secuencia todavía no corrió removeuser y sigue existiendo.
NEWUSER=$(awk -F: '$3>=1000 && $3<60000 && $1!="live" {print $1; exit}' /etc/passwd)

if [ -n "$NEWUSER" ] && [ -d "/home/$NEWUSER" ]; then
    for dir in .config .local; do
        [ -d "/home/$NEWUSER/$dir" ] || continue
        find "/home/$NEWUSER/$dir" -maxdepth 1 -type f -print0 |
            xargs -0 -r sed -i "s/quirinux/$NEWUSER/g"
    done
fi

# --- Habilitar detección de otros sistemas operativos (dual-boot) ----------
# Debian/Devuan traen esto deshabilitado por defecto. No hace falta correr
# update-grub acá: lo hace el paso genérico de bootloader que sigue.
GRUB_FILE="/etc/default/grub"
if [ -f "$GRUB_FILE" ]; then
    if grep -q "^#GRUB_DISABLE_OS_PROBER=false" "$GRUB_FILE"; then
        sed -i 's/^#GRUB_DISABLE_OS_PROBER=false/GRUB_DISABLE_OS_PROBER=false/' "$GRUB_FILE"
    elif ! grep -q "^GRUB_DISABLE_OS_PROBER=false" "$GRUB_FILE"; then
        echo "GRUB_DISABLE_OS_PROBER=false" >> "$GRUB_FILE"
    fi
fi

# --- NO incluido a propósito (ya lo hace oa-chroot-runner.sh) ---------------
#   - update-initramfs -u -k all  -> ya corre "update-initramfs -u"
#   - update-grub / update-grub2  -> ya corre "update-grub"
#   - grub-install                -> ya lo maneja install_distro_bootloader
#   - rm /etc/sudoers.d/10-installer -> ese archivo no existe en coa; el
#     equivalente actual (/etc/sudoers.d/00-live) ya lo limpia
#     shellprocess_live-cleanup.conf, que coa aplica siempre.
