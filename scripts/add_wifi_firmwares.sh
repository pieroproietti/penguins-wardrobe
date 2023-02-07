#!/bin/env bash
#
# main wifi hw for waydroid
#
# In boowkowm we don't have:
# - firmware-amd-graphics # Binary firmware for AMD/ATI graphics chips
# - firmware-atheros    # Binary firmware for Atheros wireless cards 
# - firmware-brcm80211  # Binary firmware for Broadcom/Cypress 802.11 wireless cards
# - firmware-ipw2x00 # Binary firmware for Intel Pro Wireless 2100, 2200 and 2915
# - firmware-iwlwifi # Binary firmware for Intel Wireless cards
# - firmware-libertas # Binary firmware for Marvell wireless cards
# - firmware-realtek # Binary firmware for Realtek wired/wifi/BT adapters 
# let's go
source="192.168.1.2:/opt/wardrobe/bookworm/firmware/"
mountpoint=$(mktemp -d)
dest=$(mktemp -d)
sshfs ${source} ${mountpoint}
cp ${mountpoint}/firmware-amd-graphics*.deb ${dest}
cp ${mountpoint}/firmware-atheros_*.deb ${dest}
cp ${mountpoint}/firmware-brcm80211_*.deb ${dest}
cp ${mountpoint}/firmware-ipw2x00_*.deb ${dest}
cp ${mountpoint}/firmware-iwlwifi_*.deb ${dest}
cp ${mountpoint}/firmware-libertas_*.deb ${dest}
cp ${mountpoint}/firmware-realtek_*.deb ${dest}
dpkg -i ${mountpoint}/*.deb 
