#!/bin/sh

BACKGROUND=$1
#USER=$1
#DESKTOP=$2

## FROM ENV
USER=${SUDO_USER}
DESKTOP=${EGGS_DESKTOP}


# set autologin
# We need some different to just append 
echo "# wardrobe\n[Seat:*]\nautologin-user=${USER}" >> /etc/lightdm/lightdm.conf

# set background
echo "# wardrobe\nbackground=${BACKGROUND}" >> /etc/lightdm/lightdm-gtk-greeter.conf
