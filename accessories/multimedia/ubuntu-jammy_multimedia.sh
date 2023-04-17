#!/bin/bash
clear
echo "ubuntu-jammy_multimedia"
echo ""

if [[ $EUID -ne 0 ]]; then
   printf "%s need to run with root privileges. \nPlease, prefix it with sudo\n", "$0"
   exit 1
fi

#


# install 
apt-get install -y \
    audacity \
    brasero \
    exfalso \
    kazam \
    kdenlive \
    mediainfo \
    quodlibet \
    vlc \
    xfburn

# lacks: