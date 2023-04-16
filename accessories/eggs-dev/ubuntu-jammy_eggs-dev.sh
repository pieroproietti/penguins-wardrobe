#!/bin/bash
clear
echo "ubuntu-jammy_eggs-dev"
echo ""

if [[ $EUID -ne 0 ]]; then
   printf "%s need to run with root privileges. \nPlease, prefix it with sudo\n", "$0"
   exit 1
fi

# install 
sudo apt install --force-yes \
    code \
    nodejs \
    npm

# install pnpm with npm
npm install pnpm -g

../../scripts/add_g4.sh

# lacks: