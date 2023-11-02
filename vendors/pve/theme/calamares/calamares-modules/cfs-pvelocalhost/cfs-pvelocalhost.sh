#!/bin/bash

CHROOT=$(mount | grep proc | grep calamares | awk '{print $3}' | sed -e "s#/proc##g")
IP=$(hostname -I)
HOSTNAME=$(cat $CHROOT/etc/hostname)

cat <<EOF >$CHROOT/etc/hosts
127.0.0.1 localhost localhost.localdomain
${IP} ${HOSTNAME} pvelocalhost
# The following lines are desirable for IPv6 capable hosts
::1     ip6-localhost ip6-loopback
fe00::0 ip6-localnet
ff00::0 ip6-mcastprefix
ff02::1 ip6-allnodes
ff02::2 ip6-allrouters
ff02::3 ip6-allhosts
EOF
