#!/bin/bash
#
# script to get colibri on manjaro
# 

clear
COSTUME="colibri"
MY_USERNAME=$(logname)
echo "MY_USERNAME: ${MY_USERNAME}"

MY_USERHOME="/home/${MY_USERNAME}"
echo "MY_USERHOME: ${MY_USERHOME}"

pacman -Syu \
libxfce4ui-gtk3 \
lightdm \
lightdm-gtk-greeter \
qt5ct \
shellcheck \
spice-vdagent \
thunar \
xarchiver \
xdg-user-dirs \
xfce4-appfinder \
xfce4-notifyd \
xfce4-panel \
xfce4-pulseaudio-plugin \
xfce4-session \
xfce4-settings \
xfce4-terminal \
xfce4-whiskermenu-plugin \
xfconf \
xfdesktop-gtk3 \
xfwm4 \
xorg-server 

# 
pacman -Syu \
network-manager-applet \
spice-vdagent 

# eggs-dev
pacman -Syu \
nodejs \
npm \
vscode 
#
npm install pnpm -g

# enable services 
systemctl enable NetworkManager
systemctl enable lightdm


# copy configuration from dirs
cp ./dirs/* / -R

# copy configuration from dirs to MY_USERHOME
rsync -avx ./dirs/etc/skel/.config "${MY_USERHOME}"/

# config ligghtdn $COSTUME $MY_USERHOME
../../scripts/config_lightdm.sh "${COSTUME}" "${MY_USERHOME}"

# config desktop links $MY_USERHOME
../../scripts/config_desktop_link.sh "${MY_USERHOME}"

# reimpostazione diritti
chown "${MY_USERNAME}:${MY_USERNAME}" "${MY_USERHOME}" -R
