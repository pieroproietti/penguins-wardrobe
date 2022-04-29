#!/bin/bash
RELEASE=git20220429
export EMAIL="piero.proietti@gmail.com"

# sudo apt install devscripts dh-make
sudo rmdir waydroid-build -rf

mkdir waydroid-build && cd $_

# libglibutil
git clone https://github.com/sailfishos/libglibutil.git
cd libglibutil/
sudo apt build-dep .
echo 7 > debian/compat
debuild -us -uc
cd ..
sudo dpkg -i libglibutil*.deb

# libgbinder
git clone https://github.com/mer-hybris/libgbinder.git
cd libgbinder/
sudo apt build-dep .
echo 7 > debian/compat
debuild -us -uc
cd ..
sudo dpkg -i libgbinder*.deb

# gbinder-python
git clone --single-branch --branch bullseye https://github.com/erfanoabdi/gbinder-python.git
cd gbinder-python
sudo apt build-dep .
dch --create --package "gbinder-python" --newversion "1.0.0~${RELEASE}-1" foo bar
dh_make --createorig -p "gbinder-python_1.0.0~${RELEASE}"
# select p
debuild -us -uc
cd ..
sudo dpkg -i python3-gbinder_1.0.0~${RELEASE}-1_amd64.deb

# waydroid
git clone --single-branch --branch bullseye https://github.com/waydroid/waydroid.git
cd waydroid/
sudo apt build-dep .
debuild -us -uc
cd ..
sudo dpkg -i waydroid_1.2.1_all.deb 

