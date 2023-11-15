# eagle-arm64

This started with the same code from eagle, a workstation based on proxmox VE
plus xfce4 desktop.

Here the difference, due the different architecture, are:

- removed repository `pve-no-subscription` from proxmox in favour of `pveport`
  of jiangcuo@bingsin.com;
- inserted codium to replace vscde not available;
- removed accessories: eggs-dev, firmwares/network-cable,
  firmwares/network-wifi, firmwares/video-amd not available;

# Usage

From a naked system:

- `$eggs wardrobe get`
- `$sudo eggs wardrobe wear eagle-arm64`

# To get or create a naked system

You can download a prepared naked system from sourceforge page of penguins-eggs
or you can prepare alone.

## installing a prepared naked system

Boot from ISO, then:

`sudo eggs install -un`

Note: this will destroy your hard disk.

## creating a naked system with original Debian net-install

Create a CLI only installation of Debian, add packages: `sudo, git`.

- `$git clone https://github.com/pieroproietti/addaura`
- `$cd addaura`
- `$sudo ./addaura.sh`
