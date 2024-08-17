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
cp /usr/lib/penguins-eggs/assets/eggs.png /usr/share/icons

# copy links
cp -f /usr/lib/penguins-eggs/addons/eggs/adapt/applications/eggs-adapt.desktop "${DESKTOP}"
cp -f /usr/lib/penguins-eggs/addons/eggs/adapt/bin/adapt /usr/bin/adapt
chmod +x /usr/bin/adapt
cp -f /usr/lib/penguins-eggs/assets/penguins-eggs.desktop "${DESKTOP}"
chmod +x "${DESKTOP}"/*.desktop

chown "${SUDO_USER}":"${SUDO_USER}" "${DESKTOP}" -R
