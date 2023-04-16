#!/bin/bash
clear
echo "ubuntu-jammy_eggs-dev"
echo ""

if [[ $EUID -ne 0 ]]; then
   printf "%s need to run with root privileges. \nPlease, prefix it with sudo\n", "$0"
   exit 1
fi

# sources_list_d
curl -fsSL https://deb.nodesource.com/gpgkey/nodesource.gpg.key | gpg --dearmor -o /usr/share/keyrings/nodesource.gpg > /dev/null
echo "deb [signed-by=/usr/share/keyrings/nodesource.gpg] https://deb.nodesource.com/node_16.x bullseye main" | tee /etc/apt/sources.list.d/nodesource.list > /dev/null

# update
sudo apt update 

# install 
sudo apt install --force-yes \
    nodejs \
    npm

# snap install
sudo snap install \
    code --classic

# install pnpm with npm
npm install pnpm -g

../../scripts/add_g4.sh

# lacks: