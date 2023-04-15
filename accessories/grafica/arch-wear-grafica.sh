#!/bin/bash
clear
echo "arch-wear-grafica"
echo ""

if [[ $EUID -ne 0 ]]; then
   printf "%s need to run with root privileges. \nPlease, prefix it with sudo\n", "$0"
   exit 1
fi

pacman -Syu --noconfirm \
    converseen \
    dia \
    gimp \
    imagemagick \
    kde-config-tablet \
    okular \
    pinta \
    ristretto \
    shotwell 

