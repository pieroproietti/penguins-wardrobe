#!/bin/sh

# $1 user
# $2 desktop

# copy icons
cp /usr/lib/penguins-eggs/assets/eggs.png /usr/share/icons
cp /usr/lib/penguins-eggs/addons/eggs/pve/artwork/eggs-pve.png /usr/share/icons

# copy links
cp /usr/lib/penguins-eggs/addons/eggs/pve/applications/eggs-pve.desktop $2
cp /usr/lib/penguins-eggs/addons/eggs/adapt/applications/eggs-adapt.desktop $2
cp /usr/lib/penguins-eggs/assets/penguins-eggs.desktop $2

chown $1:$1 $2 -R

