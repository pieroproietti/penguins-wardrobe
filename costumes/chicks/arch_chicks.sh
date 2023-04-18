#!/bin/bash
clear
echo "arch_chicks"
echo ""

if [[ $EUID -ne 0 ]]; then
   printf "%s need to run with root privileges. \nPlease, prefix it with sudo\n", "$0"
   exit 1
fi

COSTUME="chicks"
MY_USERNAME=$(logname)
MY_USERHOME="/home/${MY_USERNAME}"
# update

# wait to start
printf "wardrobe: Prepare your costume: %s?\n", "${COSTUME}"
echo "Press enter to continue or CTRL-C to abort"
read -r

# install costume
pacman -Syu --noconfirm \
firefox \
lightdm \
lightdm-gtk-greeter \
man \
network-manager-applet \
polkit-gnome \
shellcheck \
spice-vdagent \
unzip \
xarchiver \
xdg-user-dirs \
xfce4 \
xfce4-goodies \
xorg-apps \
xorg-server \
zenity

# add user on autologin group
groupadd -r autologin
gpasswd -a "${MY_USERNAME}" autologin

# accessories
source ../../accessories/educational/arch_educational.sh
source ../../accessories/office/arch_office.sh

# enabling services
systemctl enable NetworkManager
systemctl enable lightdm

# copy configuration from dirs to / and MY_USERHOME
cp ./dirs/* / -R
rsync -avx ./dirs/etc/skel/.config "${MY_USERHOME}"/ 
chown "${MY_USERNAME}:${MY_USERNAME}" "${MY_USERHOME}" -R

# /etc/hostname
../../scripts/hostname.sh

# config lightdm $COSTUME
../../scripts/config_lightdm.sh "${COSTUME}"
../../scripts/config_desktop_link.sh

#reboot
reboot
