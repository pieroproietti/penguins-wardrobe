#!/bin/bash

# ----------------------------------------------------------------------
# Kernel 6.12 (backports) para Quirinux
# Ejecutado por tailor en sequence.cmds, ANTES de los accesorios, para que
# los paquetes DKMS de quirinux-firmware y quirinux-drivers se compilen
# contra el kernel nuevo y no hagan fallar su postinst.
#
# Por qué: los runtimes de Flatpak 25.08 y 26.08 traen Mesa 25/26, que
# exige kernel >= 6.6. Con el 6.1 de Daedalus/Bookworm el driver amdgpu
# de Mesa se niega a funcionar dentro del sandbox, así que OBS (Flatpak) no
# detecta VAAPI y otras apps Flatpak pierden la aceleración por hardware.
#
# Instala desde <codename>-backports:
#   - linux-image-amd64 y linux-headers-amd64 (6.12.x)
#   - v4l2loopback-dkms y v4l2loopback-utils (0.13.x): la 0.12.7 de la
#     rama estable no compila con el kernel 6.12.
#
# No elimina kernels antiguos: quirinux-kernel-notice avisa de eso.
# Ignora el argumento que tailor le pasa (nombre del costume).
# ----------------------------------------------------------------------

export PATH="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
export DEBIAN_FRONTEND=noninteractive

# ----------------------------------------------------------------------
# 1. Comprobaciones previas (todas salen sin error: el costume continúa)
# ----------------------------------------------------------------------

if [ "$(dpkg --print-architecture)" != "amd64" ]; then
    echo "[config_kernel_backports] Arquitectura no amd64: se omite."
    exit 0
fi

# shellcheck disable=SC1091
. /etc/os-release
CODENAME="${VERSION_CODENAME:-}"

case "$CODENAME" in
    daedalus|bookworm) ;;
    *)
        echo "[config_kernel_backports] Codename '$CODENAME' no soportado: se omite."
        exit 0
        ;;
esac

SUITE="${CODENAME}-backports"

# ----------------------------------------------------------------------
# 2. Repositorio de backports (solo si no está ya configurado)
# ----------------------------------------------------------------------

if [ -f /etc/devuan_version ]; then
    MIRROR="http://deb.devuan.org/merged"
else
    MIRROR="http://deb.debian.org/debian"
fi

if grep -rqsE "^[[:space:]]*deb[[:space:]].*[[:space:]]${SUITE}[[:space:]]" \
        /etc/apt/sources.list /etc/apt/sources.list.d/ 2>/dev/null; then
    echo "[config_kernel_backports] ${SUITE} ya está configurado."
else
    echo "[config_kernel_backports] Añadiendo ${SUITE} (${MIRROR})."
    echo "deb ${MIRROR} ${SUITE} main contrib non-free non-free-firmware" \
        > /etc/apt/sources.list.d/quirinux-backports.list
fi

apt-get update || echo "[config_kernel_backports] apt-get update devolvió error; se continúa."

# ----------------------------------------------------------------------
# 3. Comprobar que el kernel está disponible en backports
# ----------------------------------------------------------------------

if ! apt-cache madison linux-image-amd64 | grep -q "${SUITE}"; then
    echo "[config_kernel_backports] linux-image-amd64 no está en ${SUITE}: se omite."
    exit 0
fi

# ----------------------------------------------------------------------
# 4. Retirar r8125-dkms y r8168-dkms si ya estan instalados
#    No compilan con el kernel 6.12 y hacen fallar el postinst del kernel
#    (deja dpkg a medias y el kernel sin initramfs). El driver r8169 del
#    propio kernel cubre esos chips Realtek.
# ----------------------------------------------------------------------

for pkg in r8125-dkms r8168-dkms; do
    st=$(dpkg-query -W -f='${db:Status-Abbrev}' "$pkg" 2>/dev/null)
    case "$st" in
        ""|un*|rc*) continue ;;
    esac
    echo "[config_kernel_backports] Eliminando ${pkg}: no compila en 6.12 (lo sustituye r8169)."
    apt-get purge -y "$pkg" || dpkg --purge --force-remove-reinstreq "$pkg" || true
done

# ----------------------------------------------------------------------
# 5. Instalar kernel, cabeceras y v4l2loopback desde backports
# ----------------------------------------------------------------------

echo "[config_kernel_backports] Instalando kernel y v4l2loopback desde ${SUITE}."
if ! apt-get install -y -t "${SUITE}" \
        linux-image-amd64 linux-headers-amd64 \
        v4l2loopback-dkms v4l2loopback-utils; then
    echo "[config_kernel_backports] La instalación devolvió error; se intenta reparar."
    dpkg --configure -a || true
    apt-get install -y --fix-broken || true
fi

echo "[config_kernel_backports] Kernels instalados:"
dpkg -l | awk '/^ii[[:space:]]+linux-image-[0-9]/ { print "  " $2 " " $3 }'

exit 0
