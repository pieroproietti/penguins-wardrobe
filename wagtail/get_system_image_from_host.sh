#!/bin/bash
read -p "Now I will copy system.img and vendor.img /usr/share/waydroid-extra/images. to Press enter to continue"
mkdir -p /usr/share/waydroid-extra/images
scp artisan@192.168.61.2:/eggs/*.img /usr/share/waydroid-extra/images
