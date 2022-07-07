#!/bin/bash

COSTUME=$1
BACKGROUND=$(ls /usr/share/backgrounds/"${COSTUME}"/*.jpg)

# autologin
LIGHTDM_CONF=/etc/lightdm/lightdm.conf
if [ -f /etc/lightdm/lightdm.conf.d/lightdm-autologin-greeter.conf ]; then
    LIGHTDM_CONF=/etc/lightdm/lightdm.conf.d/lightdm-autologin-greeter.conf
fi
echo -e "# wardrobe\n[Seat:*]\nautologin-user=${SUDO_USER}" >> ${LIGHTDM_CONF}

# se abbiamo xfce4-session
if [ "dpkg -l | grep -qw xfce4-session" ]; then
    # greeter-session
    echo -e "greeter-session=lightdm-gtk-greeter" >> ${LIGHTDM_CONF}

    # session
    echo -e "user-session=xfce" >> ${LIGHTDM_CONF}
fi

# background
echo -e "wardrobe\nbackground=${BACKGROUND}" >> /etc/lightdm/lightdm-gtk-greeter.conf

