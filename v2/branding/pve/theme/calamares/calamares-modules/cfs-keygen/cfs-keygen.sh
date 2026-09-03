#!/bin/bash

CHROOT=$(mount | grep proc | grep calamares | awk '{print $3}' | sed -e "s#/proc##g")

# recreate keys
chroot $CHROOT ssh-keygen -A
