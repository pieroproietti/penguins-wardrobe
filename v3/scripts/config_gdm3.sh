#!/bin/bash

DISTRO=$(lsb_release -is)
USER=${SUDO_USER}
GDM_CONF="/etc/gdm3/custom.conf"
if [ "${DISTRO}" = "Debian" ]; then
    GDM_CONF="/etc/gdm3/daemon.conf"
fi

# set autologin
echo -e "[daemon]\n# Enabling automatic login\nAutomaticLoginEnable=true\nAutomaticLogin=${USER}" >> ${GDM_CONF}
