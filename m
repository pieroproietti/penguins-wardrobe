#!/bin/sh
set -e
make clean package "$@"

. /etc/os-release

case "$ID" in
    alpine)
        sudo apk add --allow-untrusted penguins-wardrobe-*.apk
        ;;

    arch)
        sudo pacman -U --noconfirm penguins-wardrobe-*.pkg.tar.zst
        ;;

    debian)
        sudo dpkg -i penguins-wardrobe_*.deb
        ;;
    fedora)
        sudo dnf reinstall -y penguins-wardrobe-*.rpm
        ;;
    opensuse*)
        sudo zypper --no-gpg-checks install -y penguins-wardrobe-*.rpm
        ;;
    *)
        # fallback su LIKE_ID
        case "$ID_LIKE" in
            *arch*)   sudo pacman -U --noconfirm penguins-wardrobe-*.pkg.tar.zst ;;
            *debian*|*devuan*|*ubuntu*) sudo dpkg -i penguins-wardrobe_*.deb ;;
            *fedora*|*rhel*) sudo dnf install -y penguins-wardrobe-*.rpm ;;
            *) echo "Distro non supportata: $ID"; exit 1 ;;
        esac
        ;;
esac


