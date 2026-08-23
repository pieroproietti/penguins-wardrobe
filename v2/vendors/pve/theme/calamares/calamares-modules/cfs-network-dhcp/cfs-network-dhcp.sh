#!/bin/bash

CHROOT=$(mount | grep proc | grep calamares | awk '{print $3}' | sed -e "s#/proc##g")

cp /etc/network/interfaces $CHROOT/etc/network/
