#
# temporary script to get colibri in manjaro
# 

clear
COSTUME="colibri"
MY_USERNAME=`logname`
echo "MY_USERNAME: ${MY_USERNAME}"

MY_USERHOME=/home/${MY_USERNAME}
echo "MY_USERHOME: ${MY_USERHOME}"

pacman -Syu \
libxfce4ui-gtk3 \
lightdm \
lightdm-gtk-greeter \
qt5ct \
spice-vdagent \
thunar \
xarchiver \
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
xdg-user-dirs \
xorg-server 

pacman -Syu  spice-vdagent network-manager-applet

# eggs-dev
pacman -Syu vscode nodejs npm 
npm install pnpm -g

systemctl enable NetworkManager
systemctl enable lightdm


# copy configuration from dirs
cp ./dirs/* / -R

# copy configuration from dirs to MY_USERHOME
rsync -avx ./dirs/etc/skel/.config ${MY_USERHOME}/.config

# config ligghtdn $COSTUME $MY_USERHOME
../../scripts/config_lightdm.sh ${COSTUME} ${MY_USERHOME}

# config desktop links $MY_USERHOME
../../scripts/config_desktop_link.sh ${MY_USERHOME}

# Reimpostazione diritti
chown ${MY_USERNAME}:${MY_USERNAME} ${MY_USERHOME} -R
