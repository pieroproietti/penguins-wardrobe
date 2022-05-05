#!/bin/sh

USER=$1
DESKTOP=$2
BACKGROUND="none"

# set autologin
echo "[daemon]\n# Enabling automatic login\nAutomaticLoginEnable=true\nAutomaticLogin=${USER}" >> /etc/gdm3/custom.conf
