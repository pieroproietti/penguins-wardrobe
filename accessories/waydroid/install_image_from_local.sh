#!/bin/bash
clear
echo "wardrobe: Installing system.img and vendor.img from local"
read -t180 -rp "Press enter to continue..."
mkdir -p /usr/share/waydroid-extra/images
scp artisan@192.168.1.2:/opt/waydroid/lineage-18.1-20230121/*.img /usr/share/waydroid-extra/images
waydroid init -f
