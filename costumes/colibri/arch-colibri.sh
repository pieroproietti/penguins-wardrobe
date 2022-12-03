#!/bin/bash
#
# temporary script to get colibri in manjaro
# 

clear
COSTUME="colibri"
MY_USERNAME=$(logname)
echo "MY_USERNAME: ${MY_USERNAME}"

MY_USERHOME="/home/${MY_USERNAME}"
echo "MY_USERHOME: ${MY_USERHOME}"

pacman -Syyu
pacman -Syy xorg-server \
            xorg-apps \
            lightdm \
            lightdm-gtk-greeter \
            xfce4 xfce4-goodies \
            spice-vdagent \
            shellcheck \
            xdg-user-dirs \
            xarchiver \
            unzip \
            firefox \
            network-manager-applet \
            polkit-gnome

# pongo utente nel gruppo autologin
groupadd -r autologin
gpasswd -a ${MY_USERNAME} autologin

# eggs-dev
pacman -Syyu vscode nodejs npm 
npm install pnpm -g

systemctl enable NetworkManager
systemctl enable lightdm


# copy configuration from dirs
cp ./dirs/* / -R

# copy configuration from dirs to MY_USERHOME
# `rsync -avx  ${this.costume}/dirs/etc/skel/.config /home/${user}/`
rsync -avx ./dirs/etc/skel/.config "${MY_USERHOME}"/

# config ligghtdn $COSTUME $MY_USERHOME
../../scripts/config_lightdm.sh "${COSTUME}" "${MY_USERHOME}"

# config desktop links $MY_USERHOME
../../scripts/config_desktop_link.sh "${MY_USERHOME}"

# Reimpostazione diritti
chown "${MY_USERNAME}:${MY_USERNAME}" "${MY_USERHOME}" -R

echo ${COSTUME} > /etc/hostname

cat << 'EOF' > /etc/hosts
127.0.0.1 localhost localhost.localdomain
127.0.1.1 "${COSTUME}".localhost "${COSTUME}"
# The following lines are desirable for IPv6 capable hosts
:: 1     ip6 - localhost ip6 - loopback
fe00:: 0 ip6 - localnet
ff00:: 0 ip6 - mcastprefix
ff02:: 1 ip6 - allnodes
ff02:: 2 ip6 - allrouters
ff02:: 3 ip6 - allhosts
EOF

