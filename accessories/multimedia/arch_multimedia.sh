#!/bin/bash
clear
echo "arch_multimedia"
echo ""

if [[ $EUID -ne 0 ]]; then
   printf "%s need to run with root privileges. \nPlease, prefix it with sudo\n", "$0"
   exit 1
fi

pacman -Syu --noconfirm \
    audacity \
    brasero \
    kdenlive \
    mediainfo \
    quodlibet \
    vlc \
    xfburn

# lacks: exfalso, kazam

