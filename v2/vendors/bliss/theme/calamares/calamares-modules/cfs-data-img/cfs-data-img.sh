#!/bin/bash

CHROOT=$(mount | grep proc | grep calamares | awk '{print $3}' | sed -e "s#/proc##g")

# check to see if ${CHROOT}/blissos/data.img exists already, if not create it
if [ ! -f ${CHROOT}/blissos/data.img ]; then
	# Ask user what size they would like their data.img to be (4G,6G,8G,10G,12G,14G,16G, or other)
	
	echo "" | zenity --progress --text "Display test" --auto-close 2> /dev/null
	if [ $? -eq 0 ]
	then
		size=$(zenity --entry --title="Create Bliss OS data.img" --text="Enter the size would you like your data.img to be? (4G, 8G, 12G, 16G, or other):" --entry-text "8G")
	else
		size=$(dialog --title "Create Bliss OS data.img" --inputbox "Enter the size would you like your data.img to be? (4G, 8G, 12G, 16G, or other):" 8 40)
	fi
	echo "Selected: $size"
	if [ "$size" == "" ]; then
		echo "blank value found. Exiting"
		exit 1
	fi

	# Create the data.img using the $size specified in GB
	truncate -s $size ${CHROOT}/blissos/data.img
	mkfs.ext4 -F -b 4096 -L "${CHROOT}/data" "${CHROOT}/blissos/data.img"
	chmod 777 "${CHROOT}/blissos/data.img"
fi
