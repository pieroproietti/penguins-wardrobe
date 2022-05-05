#!/bin/sh

USER=$1
DESKTOP=$2
BACKGROUND="/usr/share/backgrounds/eagle/38064032662_83c2250d7f_h.jpg"


# set autologin
# We need some different to just append 
echo "# wardrobe\n[Seat:*]\nautologin-user=${USER}" >> /etc/lightdm/lightdm.conf

# set background
echo "# wardrobe\nbackground=${BACKGROUND}" >> /etc/lightdm/lightdm-gtk-greeter.conf
