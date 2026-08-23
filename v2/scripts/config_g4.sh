#!/bin/bash
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
if [ -d "${DIR}/g4" ]; then
    cp -f "${DIR}/g4/"* /usr/local/bin
elif [ -n "${SUDO_USER}" ] && [ -d "/home/${SUDO_USER}/wardrobe/v2/scripts/g4" ]; then
    cp -f /home/${SUDO_USER}/wardrobe/v2/scripts/g4/* /usr/local/bin
fi
