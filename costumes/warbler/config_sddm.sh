#!/bin/sh

USER=$1
DESKTOP=$2
BACKGROUND="/usr/share/backgrounds/helenae/25759889846_911ba6c1fd_o.jpg"

# set autologin
# We need some different to just append 
echo "# wardrobe\n[Autologin]\nUser=${USER}\nSession=plasma-wayland\n` > /etc/sddm.conf
# set background
echo "# wardrobe\nbackground=${BACKGROUND}" >> /etc/lightdm/lightdm-gtk-greeter.conf
