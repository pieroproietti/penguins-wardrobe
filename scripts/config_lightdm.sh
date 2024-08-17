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

