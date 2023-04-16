#!/bin/bash
clear
echo "arch-wear-duck"
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
source ../../accessories/eggs-dev/arch-wear-eggs-dev.sh
source ../../accessories/grafica/arch-wear-grafica.sh
source ../../accessories/multimedia/arch-wear-multimedia.sh
source ../../accessories/office/arch-wear-office.sh

# copy configuration from dirs to / and MY_USERHOME
cp ./dirs/* / -R
rsync -avx ./dirs/etc/skel/.config "${MY_USERHOME}"/ 
chown "${MY_USERNAME}:${MY_USERNAME}" "${MY_USERHOME}" -R

# config lightdm $COSTUME
../../scripts/config_lightdm.sh "${COSTUME}"

# config desktop links
../../scripts/config_desktop_link.sh

# /etc/hostname
echo ${COSTUME} > /etc/hostname

# /etc/hosts
cat << 'EOF' > /etc/hosts
127.0.0.1 localhost localhost.localdomain
127.0.1.1 duck duck.localhost 
# The following lines are desirable for IPv6 capable hosts
:: 1     ip6 - localhost ip6 - loopback
fe00:: 0 ip6 - localnet
ff00:: 0 ip6 - mcastprefix
ff02:: 1 ip6 - allnodes
ff02:: 2 ip6 - allrouters
ff02:: 3 ip6 - allhosts
EOF

#reboot
reboot

