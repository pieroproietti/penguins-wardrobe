#!/bin/bash
#
# temporary script to get colibri in manjaro
# 

clear
COSTUME="colibri"
MY_USERNAME=$(logname)
echo "MY_USERNAME: ${MY_USERNAME}"

MY_USERHOME="/home/${MY_USERNAME}"
echo "MY_USERHOME: ${MY_USERHOME}"

pacman -Syyu
pacman -S xorg-server xorg-apps
pacman -S lightdm lightdm-gtk-greeter
pacman -S xfce4 xfce4-goodies
pacman -S spice-vdagent shellcheck xdg-user-dirs
# tools
pacman -S xarchiver unzip firefox network-manager-applet polkit-gnome

# pongo utente nel gruppo autologin
gpasswd -a ${MY_USERNAME} autologin

# eggs-dev
pacman -Syu vscode nodejs npm 
npm install pnpm -g

systemctl enable NetworkManager
systemctl enable lightdm


# copy configuration from dirs
cp ./dirs/* / -R

# copy configuration from dirs to MY_USERHOME
# `rsync -avx  ${this.costume}/dirs/etc/skel/.config /home/${user}/`
rsync -avx ./dirs/etc/skel/.config "${MY_USERHOME}"/

# config ligghtdn $COSTUME $MY_USERHOME
../../scripts/config_lightdm.sh "${COSTUME}" "${MY_USERHOME}"

# config desktop links $MY_USERHOME
../../scripts/config_desktop_link.sh "${MY_USERHOME}"

# Reimpostazione diritti
chown "${MY_USERNAME}:${MY_USERNAME}" "${MY_USERHOME}" -R
