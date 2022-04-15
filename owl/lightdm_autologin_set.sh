#!/bin/sh

USER=$1
DESKTOP=$2
BACKGROUND="/usr/share/backgrounds/owl/37161836845_af56317fda_h.jpg"


# set autologin
echo autologin-user=${USER} >> /etc/lightdm/lightdm.conf

# set backgroud
echo backgroud=${BACKGROUND} >> /etc/lightdm/lightdm-gtk-greeter.conf
