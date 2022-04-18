#!/bin/sh

USER=$1
DESKTOP=$2
BACKGROUND="/usr/share/backgrounds/duck/6819323284_e1ee501d4e_b.jpg"

# set autologin
# We need some different to just append 
echo "# wardrobe\n[Seat:*]\nautologin-user=${USER}" >> /etc/lightdm/lightdm.conf

# set background
echo "# wardrobe\nbackground=${BACKGROUND}" >> /etc/lightdm/lightdm-gtk-greeter.conf
