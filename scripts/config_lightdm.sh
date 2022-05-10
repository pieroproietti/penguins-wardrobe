#!/bin/bash

COSTUME=$1
BACKGROUND=$(ls /usr/share/backgrounds/"${COSTUME}"/*.jpg)

# autologin
echo -e '# wardrobe\n[Seat:*]\nautologin-user="${SUDO_USER}"` >> /etc/lightdm/lightdm.conf

# background
echo -e "# wardrobe\nbackground=${BACKGROUND}" >> /etc/lightdm/lightdm-gtk-greeter.conf
