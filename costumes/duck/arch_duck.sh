#!/bin/bash
clear
echo "arch_duck"
echo ""

if [[ $EUID -ne 0 ]]; then
   printf "%s need to run with root privileges. \nPlease, prefix it with sudo\n", "$0"
   exit 1
fi

COSTUME="duck"
MY_USERNAME=$(logname)
MY_USERHOME="/home/${MY_USERNAME}"
# update

# wait to start
printf "wardrobe: Prepare your costume: %s?\n", "${COSTUME}"
echo "Press enter to continue or CTRL-C to abort"
read -r

# install costume
pacman -Syu --noconfirm \
cinnamon \
firefox \
gnome-terminal \
lightdm \
lightdm-gtk-greeter \
network-manager-applet \
polkit-gnome \
shellcheck \
spice-vdagent \
unzip \
xarchiver \
xdg-user-dirs \
xorg-apps \
xorg-server \
zenity

# others-stuff
pacman -Syu --noconfirm \
   gedit \
   gnome-system-monitor \
   pavucontrol \
   pulseaudio \
   pulseaudio-alsa \
   thunderbird 

# icons
pacman -Syu --noconfirm \
   mate-icon-theme-faenza

# add user on autologin group
groupadd -r autologin
gpasswd -a "${MY_USERNAME}" autologin

# enabling services
systemctl enable NetworkManager
systemctl enable lightdm

# accessories: 
source ../../accessories/eggs-dev/arch_eggs-dev.sh
source ../../accessories/grafica/arch_grafica.sh
source ../../accessories/multimedia/arch_multimedia.sh
source ../../accessories/office/arch_office.sh

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

