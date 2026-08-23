#!/bin/bash

# doas compatibile
if [[ -n "$DOAS_USER" ]]; then
    SUDO_USER="$DOAS_USER"
fi

COSTUME=$1
BACKGROUND=$(ls /usr/share/backgrounds/"${COSTUME}"/*.jpg)

# autologin
LIGHTDM_CONF=/etc/lightdm/lightdm.conf
if [ -f /etc/lightdm/lightdm.conf.d/lightdm-autologin-greeter.conf ]; then
    LIGHTDM_CONF=/etc/lightdm/lightdm.conf.d/lightdm-autologin-greeter.conf
fi
echo -e "# wardrobe\n[Seat:*]\nautologin-user=${SUDO_USER}" >> ${LIGHTDM_CONF}

# background
echo -e "wardrobe\nbackground=${BACKGROUND}" >> /etc/lightdm/lightdm-gtk-greeter.conf


# Asegurar que lightdm sea el display manager que realmente arranca.
#
# Elegir "lightdm" en la pregunta de debconf durante la instalación de
# paquetes no escribe /etc/X11/default-display-manager ni gestiona los
# scripts de arranque en un sistema sysvinit (Devuan): si el display
# manager que antes tenía el arranque habilitado (p.ej. slim) se
# desinstala como parte del proceso, el sistema se queda sin ningún
# display manager configurado para arrancar y el equipo termina
# iniciando en modo texto en vez de mostrar la pantalla de login
# gráfica. Lo fijamos explícitamente aquí, ya con lightdm instalado.
echo /usr/sbin/lightdm > /etc/X11/default-display-manager

# sysvinit / Devuan: deshabilita cualquier otro display manager que
# pudiera haber quedado con el arranque activo y habilita lightdm.
if [ -x /usr/sbin/update-rc.d ]; then
    for dm in slim gdm3 gdm sddm; do
        if [ -f "/etc/init.d/${dm}" ]; then
            update-rc.d -f "${dm}" remove >/dev/null 2>&1
        fi
    done
    update-rc.d lightdm defaults >/dev/null 2>&1
    update-rc.d lightdm enable >/dev/null 2>&1
fi

# systemd (por si este wardrobe se usa alguna vez sobre una base con
# systemd en lugar de sysvinit): habilita el servicio y el target
# gráfico por defecto.
if command -v systemctl >/dev/null 2>&1 && systemctl list-unit-files 2>/dev/null | grep -q '^lightdm\.service'; then
    systemctl set-default graphical.target >/dev/null 2>&1
    systemctl enable lightdm.service >/dev/null 2>&1
fi
