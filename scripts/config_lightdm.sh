#!/bin/sh

#BACKGROUND=$1
BACKGROUND=`ls /usr/share/backgrounds/halenae/*.jpg`
USER=${SUDO_USER}

# set autologin
# We need some different to just append 
echo "# wardrobe\n[Seat:*]\nautologin-user=${USER}" >> /etc/lightdm/lightdm.conf

# set background
echo "# wardrobe\nbackground=${BACKGROUND}" >> /etc/lightdm/lightdm-gtk-greeter.conf
