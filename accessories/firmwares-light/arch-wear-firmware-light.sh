#!/bin/bash
clear
echo "arch-wear-firmwares-light"
echo ""

if [[ $EUID -ne 0 ]]; then
   printf "%s need to run with root privileges. \nPlease, prefix it with sudo\n", "$0"
   exit 1
fi

pacman -Syu --noconfirm \
    firmware-atheros \ 
    firmware-brcm80211 \
    firmware-ipw2x00 \ 
    firmware-iwlwifi \
    firmware-libertas \
    firmware-linux-free \
    firmware-realtek \
    firmware-zd1211 \

pacman -Syu --noconfirm \
    firmware-amd-graphics \
    libdrm-amdgpu1 \
    libvulkan1 \
    mesa-opencl-icd \
    mesa-vulkan-drivers \
    vulkan-tools \
    vulkan-validationlayers \
    xserver-xorg-video-amdgpu


# lacks: 