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
xorg-xserver \
spice-vdagent \
network-manager-applet

# eggs-dev
pamac install vscode nodejs npm pnpm

systemctl enable NetworkManager
systemctl enable lightdm
