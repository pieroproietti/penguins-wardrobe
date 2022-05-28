#!/bin/bash

DESKTOP=$(su "${SUDO_USER}" xdg-user-dir DESKTOP)

# create DESKTOP
if [ ! -d "$DESKTOP" ]; then
    mkdir "$DESKTOP" -p
fi

# copy icons

cp /usr/lib/penguins-eggs/assets/eggs.png /usr/share/icons

# copy links
cp /usr/lib/penguins-eggs/addons/eggs/adapt/applications/eggs-adapt.desktop "${DESKTOP}"
cp /usr/lib/penguins-eggs/addons/eggs/adapt/bin/adapt /usr/local/bin/adapt
cp /usr/lib/penguins-eggs/assets/penguins-eggs.desktop "${DESKTOP}"
chown "${SUDO_USER}":"${SUDO_USER}" "${DESKTOP}" -R
