# livecd customization

LiveCD customization is composed of:

- grub.theme.cfg
- isolinux.theme.cfg
- splash.png

## How it work

eggs take cure to copy your theme configurations on /boot/grub and /isolinux on
the ISO image.

## grub

You can adapt the grub theme file according to your needs. During the production
of eggs a grub.cfh will be geerated and copied under `/boot/grub` directory of
the image.

- **grub.theme.cfg** -> will became -> **/boot/grub/theme.cfg** under the iso

## isolinux

You can adapt the isolinux theme. During the production of eggs an
isolilinux.cfg will be generated and copied under `/isolinux` directory of the
image.

- **isolinux.theme.cfg** -> will became -> **/isolinux/theme.cfg** under the iso

## splash

Both isolinux and grub themes use the same splash file.
