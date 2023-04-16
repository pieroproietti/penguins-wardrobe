#!/bin/bash
clear
echo "arch_grafica"
echo ""

if [[ $EUID -ne 0 ]]; then
   printf "%s need to run with root privileges. \nPlease, prefix it with sudo\n", "$0"
   exit 1
fi

pacman -Syu --noconfirm \
    converseen \
    gimp \
    imagemagick \
    okular \
    pinta \
    ristretto \
    shotwell 

# lacks: dia, kde-config-tablet