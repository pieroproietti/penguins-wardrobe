#!/bin/bash

CHROOT=$(mount | grep proc | grep calamares | awk '{print $3}' | sed -e "s#/proc##g")

SCRIPT_PATH=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
SOURCE_NAME="blissos"
# First, we check for a .iso file in /updates/blissos 
# If found we use that as the path to the .zip file

if [ -f ${CHROOT}/updates/blissos/update.zip ]; then
    ZIP_PATH=${CHROOT}/updates/blissos/update.zip
else
	echo -e $(ls ${CHROOT})
	echo -e $(ls ${CHROOT}/updates)
	echo -e $(ls ${CHROOT}/updates/blissos)
	exit 0 # don't abort it's installed!
fi

# Check to see if we have a /blissos/ folder already
if [ ! -d ${CHROOT}/blissos ]; then
    echo "${CHROOT}/blissos Folder does not exists"
    mkdir -p ${CHROOT}/blissos
else
	echo "${CHROOT}/blissos folder found"
fi

echo -e "Extracting the zip now..."
7z e $ZIP_PATH -o${CHROOT}/blissos/

if [ ! -f ${CHROOT}/blissos/system.img ]; then
	echo "Somethings wrong, exiting"
	exit 1
fi

echo -e "Cleaning Up..."
rm -rf ${CHROOT}/updates/blissos/*.zip

echo "All set. Thanks for installing."
