#!/bin/bash

CHROOT=$(mount | grep proc | grep calamares | awk '{print $3}' | sed -e "s#/proc##g")

tee $CHROOT/etc/mkinitcpio.conf << EOF
MODULES=""
BINARIES=""
FILES=""
HOOKS="base udev akshara plymouth autodetect modconf block keyboard keymap consolefont filesystems fsck"
COMPRESSION="zstd"

EOF

chroot $CHROOT mkinitcpio /etc/mkinitcpio.conf -g /boot/initramfs-linux-zen.img