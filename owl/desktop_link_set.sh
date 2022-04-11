#!/bin/sh

USER=$1
DESKTOP=$2

# create DESKTOP
if [ ! -d "$DESKTOP" ]; then
    mkdir $DESKTOP -p
fi

# copy icons

cp /usr/lib/penguins-eggs/assets/eggs.png /usr/share/icons

# copy links
cp /usr/lib/penguins-eggs/addons/eggs/adapt/applications/eggs-adapt.desktop $DESKTOP
cp /usr/lib/penguins-eggs/assets/penguins-eggs.desktop $DESKTOP
chown $USER:$USER $DESKTOP -R
