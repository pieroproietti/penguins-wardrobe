#!/bin/bash
clear
echo "arch-wear-colibri"
echo ""

if [[ $EUID -ne 0 ]]; then
   printf "%s need to run with root privileges. \nPlease, prefix it with sudo\n", "$0"
   exit 1
fi

COSTUME="colibri"
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
source ../../accessories/grafica/arch-wear-eggs-dev.sh

# enabling services
systemctl enable NetworkManager
systemctl enable lightdm

# copy configuration from dirs to / and MY_USERHOME
cp ./dirs/* / -R
rsync -avx ./dirs/etc/skel/.config "${MY_USERHOME}"/ 
chown "${MY_USERNAME}:${MY_USERNAME}" "${MY_USERHOME}" -R

# config lightdm $COSTUME $MY_USERHOME
../../scripts/config_lightdm.sh "${COSTUME}" "${MY_USERHOME}"

# config desktop links $MY_USERHOME
../../scripts/config_desktop_link.sh "${MY_USERHOME}"

# /etc/hostname
echo ${COSTUME} > /etc/hostname

# /etc/hosts
cat << 'EOF' > /etc/hosts
127.0.0.1 localhost localhost.localdomain
127.0.1.1 colibri colibri.localhost 
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
