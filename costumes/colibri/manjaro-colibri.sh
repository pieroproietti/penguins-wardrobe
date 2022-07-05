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

# copy configuration from dirs
cp dirs/* / -R
rsync -avx ./dirs/etc/skel/.config ${HOME}


# SUDO_USER=esiste
COSTUME="colibri"

../../scripts/config_lightdm.sh colibri /home/${SUDO_USER}
chown ${SUDO_USER}:${SUDO_USER} /home/${SUDO_USER} -R
../../scripts/config_desktop_link.sh