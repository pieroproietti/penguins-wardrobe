#!/bin/sh
DESKTOP=$(xdg-user-dir DESKTOP)
if [ ! -f ~/.linuxfx-eggs ]; then
    xdg-user-dirs-update > /dev/null
    xdg-user-dirs-update --force > /dev/null
    penguins-links-add.sh > /dev/null
    linuxfx-links-add.sh > /dev/null
    cp /usr/share/applications/penguins-eggs.desktop $DESKTOP > /dev/null
    cp /usr/share/applications/devicemanager.desktop $DESKTOP > /dev/null
    cp /usr/share/applications/microsoft-edge.desktop $DESKTOP > /dev/null
    cp /opt/Linuxfx/helloa/trash.desktop $DESKTOP > /dev/null
    cp /usr/share/applications/windowsfx-android.desktop $DESKTOP > /dev/null
    chmod +x $DESKTOP/*.desktop > /dev/null
    touch ~/.linuxfx-eggs > /dev/null
fi
