#!/bin/bash

# ----------------------------------------------------------------------
# Instalacion tolerante de modulos DKMS para Quirinux
# Ejecutado por tailor en sequence.cmds, justo despues de
# config_kernel_backports.sh y ANTES de los accesorios.
#
# Por que: un paquete *-dkms cuyo modulo no compila contra el kernel
# instalado devuelve error en su postinst y deja dpkg a medias. A partir de
# ahi apt se niega a instalar nada y tailor reintenta cada paquete sin
# exito. Con el kernel 6.12 pasa con modulos que solo soportaban el 6.1.
#
# Este script instala de uno en uno los modulos cuya compatibilidad con
# kernels nuevos no esta garantizada. Si uno falla:
#   - guarda el make.log de DKMS en el informe,
#   - purga el paquete para dejar dpkg limpio,
#   - sigue con el resto.
# Los modulos que compilan se instalan con normalidad.
#
# Ademas bloquea (pin -1 en /etc/apt/preferences.d/quirinux-dkms-blocked) los
# paquetes que no compilan, para que quirinux-firmware, que los lista en
# Recommends, no los vuelva a traer y rompa dpkg. Para permitirlos de nuevo,
# borrar ese fichero.
#
# Informe: /var/log/tailor/quirinux-dkms-report.txt
# Ignora el argumento que tailor le pasa (nombre del costume).
# ----------------------------------------------------------------------

export PATH="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
export DEBIAN_FRONTEND=noninteractive

# Modulos DKMS con compatibilidad dudosa en kernels nuevos
DKMS_PACKAGES="broadcom-sta-dkms bcmwl-kernel-source digimend-dkms"

# Paquetes que se sabe que no compilan con kernel >= 6.12 (los cubre r8169)
BLOCKED=""
NEWEST_KERNEL=$(dpkg -l | awk '/^ii[[:space:]]+linux-image-[0-9]/ { sub("linux-image-", "", $2); print $2 }' | sort -V | tail -n1)
if [ -n "$NEWEST_KERNEL" ] && dpkg --compare-versions "$NEWEST_KERNEL" ge 6.12; then
    BLOCKED="r8125-dkms r8168-dkms"
fi

REPORT_DIR="/var/log/tailor"
REPORT="${REPORT_DIR}/quirinux-dkms-report.txt"
mkdir -p "$REPORT_DIR"
{
    echo "Informe DKMS de Quirinux - $(date '+%Y-%m-%d %H:%M:%S')"
    echo "Kernels instalados:"
    dpkg -l | awk '/^ii[[:space:]]+linux-image-[0-9]/ { print "  " $2 " " $3 }'
    echo
} > "$REPORT"

SKIPPED=""

is_installed() {
    dpkg -s "$1" 2>/dev/null | grep -q '^Status: install ok installed'
}

save_build_logs() {
    # Cola de cada make.log de DKMS para poder diagnosticar sin acceso a la VM
    local pkg=$1 log
    {
        echo "---- FALLO al instalar ${pkg} ----"
        for log in /var/lib/dkms/*/*/build/make.log; do
            [ -f "$log" ] || continue
            echo "== ${log}"
            tail -n 25 "$log"
        done
        echo
    } >> "$REPORT"
}

for pkg in $DKMS_PACKAGES; do
    if is_installed "$pkg"; then
        echo "[config_dkms_safe] ${pkg} ya esta instalado."
        continue
    fi

    if ! apt-cache show "$pkg" > /dev/null 2>&1; then
        echo "[config_dkms_safe] ${pkg} no esta disponible en los repositorios; se omite."
        continue
    fi

    echo "[config_dkms_safe] Instalando ${pkg}."
    if apt-get install -y "$pkg"; then
        echo "[config_dkms_safe] ${pkg}: OK."
        continue
    fi

    echo "[config_dkms_safe] ${pkg} NO se pudo instalar con este kernel; se omite."
    save_build_logs "$pkg"
    apt-get purge -y "$pkg" \
        || dpkg --purge --force-remove-reinstreq "$pkg" \
        || true
    dpkg --configure -a || true
    apt-get install -y --fix-broken || true
    SKIPPED="${SKIPPED} ${pkg}"
    BLOCKED="${BLOCKED} ${pkg}"
done

# Bloquear los paquetes que no compilan para que Recommends de otros
# paquetes (quirinux-firmware) no los vuelvan a instalar
BLOCKED=$(echo $BLOCKED)
if [ -n "$BLOCKED" ]; then
    {
        echo "# Generado por config_dkms_safe.sh: modulos DKMS que no compilan con el"
        echo "# kernel instalado. Borrar este fichero para permitirlos de nuevo."
        echo "Package: ${BLOCKED}"
        echo "Pin: version *"
        echo "Pin-Priority: -1"
    } > /etc/apt/preferences.d/quirinux-dkms-blocked
    echo "Paquetes bloqueados en apt (pin -1): ${BLOCKED}" >> "$REPORT"
    apt-get update > /dev/null 2>&1 || true
fi

{
    echo "---- Estado final de DKMS ----"
    dkms status 2>&1
    echo
    if [ -n "$SKIPPED" ]; then
        echo "Paquetes omitidos por no compilar:${SKIPPED}"
    else
        echo "Todos los modulos DKMS dudosos se instalaron correctamente."
    fi
} >> "$REPORT"

echo "[config_dkms_safe] Informe guardado en ${REPORT}"
[ -n "$SKIPPED" ] && echo "[config_dkms_safe] Omitidos:${SKIPPED}"

exit 0
