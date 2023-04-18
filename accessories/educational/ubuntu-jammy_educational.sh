#!/bin/bash
clear
echo "ubuntu-jammy_educational"
echo ""

if [[ $EUID -ne 0 ]]; then
   printf "%s need to run with root privileges. \nPlease, prefix it with sudo\n", "$0"
   exit 1
fi

# install 
apt-get install -y \
    gelemental \
    granule \
    klavaro \
    littlewizard \
    solfege \
    stellarium \
    tuxtype \
    wordplay

# lacks: gcompris gnome-chemistry-utils