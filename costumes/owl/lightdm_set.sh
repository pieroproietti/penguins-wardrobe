#!/bin/sh

USER=$1
DESKTOP=$2
BACKGROUND="/usr/share/backgrounds/owl/37161836845_af56317fda_h.jpg"


# set autologin
# We need some different to just append 
echo "# wardrobe\n[Seat:*]\nautologin-user=${USER}" >> /etc/lightdm/lightdm.conf

# set background
echo "# wardrobe\nbackground=${BACKGROUND}" >> /etc/lightdm/lightdm-gtk-greeter.conf
