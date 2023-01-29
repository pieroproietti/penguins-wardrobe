#!/bin/sh

# Here we are always root and always are working on a virgin system
tmp=`mktemp -d`
cd $tmp

#git clone https://github.com/electrikjesus/waydroid-package-manager
#rm ./waydroid-package-manager/.git
#mv ./waydroid-package-manager /etc/skel
scp -r artisan@192.168.1.2:/opt/waydroid/waydroid-package-manager /etc/skel





