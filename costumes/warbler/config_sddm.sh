#!/bin/sh

COSTUME=$1
BACKGROUND=$(ls /usr/share/backgrounds/"${COSTUME}"/*.jpg)

# select plasma, not plasma-wayland becouse don't work autologin
echo "# wardrobe\n[Autologin]\nUser=${SUDO_USER}\nSession=plasma\n" > /etc/sddm.conf
