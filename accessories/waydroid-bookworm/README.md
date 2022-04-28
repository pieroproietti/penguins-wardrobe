# Method 1: build and install from source

sudo apt install python3-pip

For development, you will also need Cython:

pip install cython


Make a new directory and gather all the dependencies
```
mkdir Waydroid-build && cd $_
git clone https://github.com/sailfishos/libglibutil.git
git clone https://github.com/mer-hybris/libgbinder.git
git clone --single-branch --branch bullseye https://github.com/erfanoabdi/gbinder-python.git
```

Install the necessary Debian build tools

```
sudo apt install devscripts dh-make
```
Build and install Waydroid dependencies
```
cd libglibutil/
sudo apt build-dep .
```
edit debian/compat and place 7
```
debuild -us -uc
cd ..
sudo dpkg -i libglibutil*.deb
cd libgbinder/
sudo apt build-dep .
```
edit debian/compat and place 7
```
debuild -us -uc
cd ..
sudo dpkg -i libgbinder*.deb
```
gbinder-python has incomplete Debian packaging and needs a little more work, remember to update the git version number next time you build:


```
rm gbinder-python -rf
git clone --single-branch --branch bullseye https://github.com/erfanoabdi/gbinder-python.git
cd gbinder-python/
sudo apt build-dep .
dch --create --package "gbinder-python" --newversion "1.0.0~git20220426-1" foo bar
dh_make --createorig -p "gbinder-python_1.0.0~git20220426"
```
# Select "p" when prompted for the package type, leave the rest at the defaults
```
debuild -us -uc
cd ..
sudo dpkg -i python3-gbinder_1.0.0~git20220426-1_amd64.deb
```
Build and install Waydroid

```
git clone --single-branch --branch bullseye https://github.com/waydroid/waydroid.git
cd waydroid/
sudo apt build-dep .
debuild -us -uc
cd ..
sudo dpkg -i waydroid_1.2.1_all.deb 
```