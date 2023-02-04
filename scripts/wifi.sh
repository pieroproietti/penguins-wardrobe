#!/bin/env bash
#
# main wifi hw for waydroid
#
mkdir /tmp/firmwares
cd /tmp/firmwares
GET="wget http://ftp.it.debian.org/debian/pool/non-free/f/firmware-nonfree/"
${GET}firmware-atheros_20210818-1~bpo11%2b1_all.deb # firmware-atheros
${GET}firmware-brcm80211_20210818-1~bpo11%2b1_all.deb # Binary firmware for Broadcom/Cypress 802.11 wireless card
${GET}firmware-ipw2x00_20210818-1~bpo11%2b1_all.deb #  Binary firmware for Intel Pro Wireless 2100, 2200 and 2915 [non-free]
${GET}firmware-iwlwifi_20210818-1~bpo11%2b1_all.deb # Binary firmware for Intel Wireless cards [non-free]
${GET}firmware-libertas_20210818-1~bpo11%2b1_all.deb # binary firmware for Marvell wireless cards [non-free]
${GET}firmware-linux-free 
${GET}firmware-realtek_20210818-1~bpo11%2b1_all.deb # Binary firmware for Realtek wired/wifi/BT adapters [non-free]
dpkg -i /tmp/firmwares/*.deb
