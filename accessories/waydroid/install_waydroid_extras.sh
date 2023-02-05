#!/bin/sh

# Here we are always root and always are working on a virgin system
tmp=`mktemp -d`
cd $tmp

git clone https://github.com/casualsnek/waydroid_script
cd waydroid_script
python3 -m pip install -r requirements.txt
python3 waydroid_extras.py -i # [-i/-g/-n/-h]