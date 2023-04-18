#!/bin/bash
clear
echo "ubuntu-jammy_owl"
echo ""

if [[ $EUID -ne 0 ]]; then
   printf "%s need to run with root privileges. \nPlease, prefix it with sudo\n", "$0"
   exit 1
fi

COSTUME="owl"
MY_USERNAME=$(logname)
MY_USERHOME="/home/${MY_USERNAME}"

# wait to start
printf "wardrobe: Prepare your costume: %s?\n", "${COSTUME}"
echo "Press enter to continue or CTRL-C to abort"
read -r

# unpdate
apt-get update

# install costume
apt-get install -y \
    adwaita-qt \
    firefox \
    libxfce4ui-utils  \
    lightdm  \
    lightdm-gtk-greeter  \
    network-manager-gnome \
    spice-vdagent  \
    tango-icon-theme  \
    thunar  \
    xarchiver  \
    xfce4-appfinder \
    xfce4-notifyd \
    xfce4-panel \
    xfce4-pulseaudio-plugin \
    xfce4-session  \
    xfce4-settings \
    xfce4-terminal \
    xfce4-whiskermenu-plugin \
    xfconf \
    xfdesktop4 \
    xfwm4 \
    zenity

# add user on autologin group
#groupadd -r autologin
#gpasswd -a "${MY_USERNAME}" autologin

# accessories
source ../../accessories/grafica/ubuntu-jammy_grafica.sh
source ../../accessories/multimedia/ubuntu-jammy_multimedia.sh
source ../../accessories/office/ubuntu-jammy_office.sh

# enabling services
#systemctl enable NetworkManager
#systemctl enable lightdm

# copy configuration from dirs to / and MY_USERHOME
cp ./dirs/* / -R
rsync -avx ./dirs/etc/skel/.config "${MY_USERHOME}"/ 
chown "${MY_USERNAME}:${MY_USERNAME}" "${MY_USERHOME}" -R

# /etc/hostname
../../scripts/hostname.sh "${COSTUME}"

# config lightdm $COSTUME
../../scripts/config_lightdm.sh "${COSTUME}"
../../scripts/config_desktop_link.sh


# reboot
reboot
