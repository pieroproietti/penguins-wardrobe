#!/bin/bash
clear
echo "arch_naked"
echo ""

if [[ $EUID -ne 0 ]]; then
   printf "%s need to run with root privileges. \nPlease, prefix it with sudo\n", "$0"
   exit 1
fi

COSTUME="colibri"
MY_USERNAME=$(logname)
MY_USERHOME="/home/${MY_USERNAME}"
# update

# wait to start
printf "wardrobe: Prepare your costume: %s?\n", "${COSTUME}"
echo "Press enter to continue or CTRL-C to abort"
read -r

# install costume
pacman -Syu --noconfirm \
   man 
   
# /etc/hostname
../../scripts/hostname.sh

#reboot
reboot
