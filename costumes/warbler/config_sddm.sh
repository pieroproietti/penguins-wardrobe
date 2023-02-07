#!/bin/sh

COSTUME=$1
BACKGROUND=$(ls /usr/share/backgrounds/"${COSTUME}"/*.jpg)

# set autologin
# We need some different to just append 
echo "# wardrobe\n[Autologin]\nUser=${SUDO_USER}\nSession=plasma\n" > /etc/sddm.conf
# set background
#echo "# wardrobe\nbackground=${BACKGROUND}" >> /etc/lightdm/lightdm-gtk-greeter.conf