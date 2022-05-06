#!/bin/sh

#USER=$1
#DESKTOP=$2
USER=${SUDO_USER}
DESKTOP=${EGGS_DESKTOP}
BACKGROUND=${EGGS_BACKGROUND}

# set autologin
# We need some different to just append 
echo "# wardrobe\n[Seat:*]\nautologin-user=${USER}" >> /etc/lightdm/lightdm.conf

# set background
echo "# wardrobe\nbackground=${BACKGROUND}" >> /etc/lightdm/lightdm-gtk-greeter.conf
