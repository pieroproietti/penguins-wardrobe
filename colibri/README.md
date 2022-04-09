# colibri


#!/bin/bash
# https://forum.xfce.org/viewtopic.php?id=10894
#sudo rsync -avx  .penguins-eggs/wardrobe.d/$1/dirs/etc/skel/.config /home/artisan/
#sudo chown artisan:artisan /home/artisan/ -R
# To determine the correct location of the backdrop entry, run:
# xfconf-query -c xfce4-desktop -m

# list
# xfconf-query --channel xfce4-desktop -l

xfconf-query --channel xfce4-desktop \
  -p backdrop/screen0/monitorVirtual-1/workspace0/last-image \
  -s "$1"
