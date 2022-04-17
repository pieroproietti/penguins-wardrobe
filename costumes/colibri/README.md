# colibri

Una remix pensata per sviluppare in typescript con node-16x, code e quanto occorre per eggs

# Problemi

## background
Determinare la procedura per il cambio del desktop

```
xfce4_background_set.sh
```

Comandi utili:

## list
```
xfconf-query --channel xfce4-desktop -l
```

## vedere i cambiamenti
```
xfconf-query -c xfce4-desktop -m
```

## cambiare i desktop

```
s#!/bin/sh
xfconf-query --channel xfce4-desktop \
-p /backdrop/screen0/monitor0/last-image \
-s "$1"

xfconf-query --channel xfce4-desktop \
-p /backdrop/screen0/monitor1/last-image \
-s "$1"

xfconf-query --channel xfce4-desktop \
-p /backdrop/screen0/monitorVirtual-1/workspace0/last-image \
-s "$1"

xfconf-query --channel xfce4-desktop \
-p /backdrop/screen0/monitorVirtual-1/workspace1/last-image \
-s "$1"

xfconf-query --channel xfce4-desktop \
-p /backdrop/screen0/monitorVirtual-1/workspace2/last-image \
-s "$1"

xfconf-query --channel xfce4-desktop \
-p /backdrop/screen0/monitorVirtual-1/workspace3/last-image \
-s "$1"
```

# background
![colibri](./dirs/usr/share/backgrounds/colibri/3794764350_2839ca0b26_b.jpg)

