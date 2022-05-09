#!/bin/bash

COSTUME=$1
BACKGROUND=$(ls /usr/share/backgrounds/"${COSTUME}"/*.jpg)
USER=${SUDO_USER}

# autologin
echo -e "# wardrobe\n[Seat:*]\nautologin-user=${USER}" >> /etc/lightdm/lightdm.conf

# background
echo -e "# wardrobe\nbackground=${BACKGROUND}" >> /etc/lightdm/lightdm-gtk-greeter.conf
