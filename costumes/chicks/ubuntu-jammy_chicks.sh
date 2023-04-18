#!/bin/bash
clear
echo "ubuntu-jammy_chicks"
echo ""

if [[ $EUID -ne 0 ]]; then
   printf "%s need to run with root privileges. \nPlease, prefix it with sudo\n", "$0"
   exit 1
fi

COSTUME="chicks"
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

# accessories:
source ../../accessories/office/ubuntu-jammy_office.sh

sudo apt install -y \
    epoptes \
    libreoffice-l10n-it \
    firefox-l10n-it

# enabling services
systemctl enable NetworkManager
systemctl enable lightdm

# copy configuration from dirs to / and MY_USERHOME
cp ./dirs/* / -R
rsync -avx ./dirs/etc/skel/.config "${MY_USERHOME}"/ 
chown "${MY_USERNAME}:${MY_USERNAME}" "${MY_USERHOME}" -R

# config lightdm $COSTUME
../../scripts/config_lightdm.sh "${COSTUME}"
../../scripts/config_desktop_link.sh

# /etc/hostname
echo ${COSTUME} > /etc/hostname

# /etc/hosts
cat << 'EOF' > /etc/hosts
127.0.0.1 localhost localhost.localdomain
127.0.1.1 chicks chicks.localhost 
# The following lines are desirable for IPv6 capable hosts
:: 1     ip6 - localhost ip6 - loopback
fe00:: 0 ip6 - localnet
ff00:: 0 ip6 - mcastprefix
ff02:: 1 ip6 - allnodes
ff02:: 2 ip6 - allrouters
ff02:: 3 ip6 - allhosts
EOF

# reboot
reboot
