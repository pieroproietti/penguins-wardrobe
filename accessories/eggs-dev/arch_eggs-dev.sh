#!/bin/bash
clear
echo "arch_eggs-dev"
echo ""

if [[ $EUID -ne 0 ]]; then
   printf "%s need to run with root privileges. \nPlease, prefix it with sudo\n", "$0"
   exit 1
fi

pacman -Syu --noconfirm \
    nodejs \
    npm \
    pacman-contrib \
    vscode 

# install pnpm with npm
npm install pnpm -g

../../scripts/config_g4.sh

# lacks: