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
curDir=$(pwd)
tmp=`mktemp -d`
cd $tmp
GET="wget http://ftp.it.debian.org/debian/pool/non-free/f/firmware-nonfree/"
${GET}firmware-amd-graphics_20210818-1~bpo11+1_all.deb
${GET}firmware-atheros_20210818-1~bpo11%2b1_all.deb 
${GET}firmware-brcm80211_20210818-1~bpo11%2b1_all.deb
${GET}firmware-ipw2x00_20210818-1~bpo11%2b1_all.deb
${GET}firmware-iwlwifi_20210818-1~bpo11%2b1_all.deb
${GET}firmware-libertas_20210818-1~bpo11%2b1_all.deb
${GET}firmware-realtek_20210818-1~bpo11%2b1_all.deb
dpkg -i *.deb
cd $curDir
