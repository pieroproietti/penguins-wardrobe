#!/bin/bash
clear
echo "arch_educational"
echo ""

if [[ $EUID -ne 0 ]]; then
   printf "%s need to run with root privileges. \nPlease, prefix it with sudo\n", "$0"
   exit 1
fi

pacman -Syu --noconfirm \
    klavaro \
    solfege 

# lacks:

# yay --noconfirm gelemental
# yay --noconfirm gcompris 
# yay --noconfirm granule 
# yay --noconfirm littlewizard 
# yay --noconfirm stellarium 
# yay --noconfirm tuxtype
# yay --noconfirm wordplay