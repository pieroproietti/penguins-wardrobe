#!/bin/bash
DEST=~/WAYDROID-DEBS
if [ ! $# -eq 0 ]; then
    DEST=$1
fi

rm -f $DEST 
if [ ! -d "$DEST" ]; then
    mkdir $DEST
fi


# Here we are always root and always are working on a virgin system
tmp=`mktemp -d`
RELEASE=git$(date '+%Y%m%d')
export EMAIL="piero.proietti@gmail.com"

# install waydroid prerequisites
sudo apt install ca-certificates git iptables lxc python3 wget
# install waydroid prerequisites 2
sudo apt install gir1.2-gtk-3.0 gir1.2-vte-2.91 gir1.2-webkit2-4.0 libvte-2.91-dev python3-gi python3-gi-cairo python3-pip xclip

# install prerequisites for build
sudo apt install devscripts dh-make bison flex cython3 dh-python libpkgconf3 libpython3-all-dev pkgconf python3-all python3-all-dev
# sudo pip install cython
read -p "Press enter to continue"
clear

cd $tmp
mkdir waydroid-build && cd $_

###################################################################################################
# libglibutil
###################################################################################################
clear
echo libgllibutil
git clone https://github.com/sailfishos/libglibutil.git
cd libglibutil/
sudo apt build-dep .
echo 7 > debian/compat
debuild -us -uc
cd ..
sudo dpkg -i libglibutil*.deb
# read -p "Press enter to continue"

###################################################################################################
# libgbinder
###################################################################################################
clear
libgbinder
git clone https://github.com/mer-hybris/libgbinder.git
cd libgbinder/
sudo apt build-dep .
echo 7 > debian/compat
debuild -us -uc
cd ..
sudo dpkg -i libgbinder*.deb
read -p "Press enter to continue"

###################################################################################################
# gbinder-python
###################################################################################################
clear
echo gbinder-python
git clone --single-branch --branch bullseye https://github.com/erfanoabdi/gbinder-python.git
cd gbinder-python
sudo apt build-dep .
dch --create --package "gbinder-python" --newversion "1.0.0~${RELEASE}-1" foo bar
dh_make --createorig -p "gbinder-python_1.0.0~${RELEASE}"
# select p
debuild -us -uc
cd ..
sudo dpkg -i python3-gbinder_1.0.0~${RELEASE}-1_amd64.deb
read -p "Press enter to continue"

###################################################################################################
# waydroid
###################################################################################################
clear
echo waydroid
git clone --single-branch --branch bullseye https://github.com/waydroid/waydroid.git
cd waydroid/
sudo apt build-dep .
debuild -us -uc
cd ..
sudo dpkg -i waydroid_1.2.1_all.deb
read -p "Press enter to continue"
cp $tmp/waydroid-build/*.deb $DEST

