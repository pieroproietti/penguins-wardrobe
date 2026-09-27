#!/bin/bash

# doas compatibile
if [[ -n "$DOAS_USER" ]]; then
    SUDO_USER="$DOAS_USER"
fi

DESKTOP=$(su "${SUDO_USER}" xdg-user-dir DESKTOP)

# create DESKTOP
if [ ! -d "$DESKTOP" ]; then
    mkdir "$DESKTOP" -p
fi

# copy icons
cp /usr/lib/oa-tools/assets/coa.png /usr/share/icons

# copy links
cp -f /usr/lib/oa-tools/addons/coa/adapt/applications/coa-adapt.desktop "${DESKTOP}"
cp -f /usr/lib/oa-tools/assets/oa-tools.desktop "${DESKTOP}"
chmod +x "${DESKTOP}"/*.desktop

chown "${SUDO_USER}":"${SUDO_USER}" "${DESKTOP}" -R
