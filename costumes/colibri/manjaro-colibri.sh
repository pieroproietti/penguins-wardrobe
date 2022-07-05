# here we define a colibri for manjaro
# I need experience here
pamac install \
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

pamac spice-vdagent \
network-manager-applet

# eggs-dev
pamac install vscode nodejs npm pnpm

systemctl enable NetworkManager
systemctl enable lightdm

clear

# SUDO_USER=esiste
COSTUME="colibri"
MY_USERNAME=`logname`
echo "MY_USERNAME: ${MY_USERNAME}"

MY_USERHOME=/home/${MY_USERNAME}
echo "MY_USERHOME: ${MY_USERHOME}"

# copy configuration from dirs
cp ./dirs/* / -R
echo "rsync -avx ./dirs/etc/skel/.config ${MY_USERHOME}/.config"
rsync -avx ./dirs/etc/skel/.config ${MY_USERHOME}/.config

echo "../../scripts/config_lightdm.sh colibri ${MY_USERHOME}"
../../scripts/config_lightdm.sh colibri ${MY_USERHOME}

echo "../../scripts/config_desktop_link.sh ${MY_USERHOME}"
../../scripts/config_desktop_link.sh ${MY_USERHOME}

echo "chown ${MY_USERNAME}:${MY_USERNAME} ${MY_USERHOME} -R"
chown ${MY_USERNAME}:${MY_USERNAME} ${MY_USERHOME} -R
