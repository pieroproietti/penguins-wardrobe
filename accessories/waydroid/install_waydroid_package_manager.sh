#!/bin/sh
#
# Here we are always root and always are working on a virgin system
#
clear
echo "wardrobe: installing waydroid-package-manager"
read -t180 -rp "Press enter to continue..."
scp -r artisan@192.168.1.2:/opt/waydroid/waydroid-package-manager /etc/skel
