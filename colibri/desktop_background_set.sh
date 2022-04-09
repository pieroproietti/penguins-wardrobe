#!/bin/bash
BACKGROUND="/usr/share/backgrounds/colibri-wallpapers/3794764350_2839ca0b26_b.jpg"
xfconf-query --channel xfce4-desktop \
-p /backdrop/screen0/monitor0/last-image \
-s "$BACKGROUND"

xfconf-query --channel xfce4-desktop \
-p /backdrop/screen0/monitor1/last-image \
-s "$BACKGROUND"

xfconf-query --channel xfce4-desktop \
-p /backdrop/screen0/monitorVirtual-1/workspace0/last-image \
-s "$BACKGROUND"

xfconf-query --channel xfce4-desktop \
-p /backdrop/screen0/monitorVirtual-1/workspace1/last-image \
-s "$BACKGROUND"

xfconf-query --channel xfce4-desktop \
-p /backdrop/screen0/monitorVirtual-1/workspace2/last-image \
-s "$BACKGROUND"

xfconf-query --channel xfce4-desktop \
-p /backdrop/screen0/monitorVirtual-1/workspace3/last-image \
-s "$BACKGROUND"
