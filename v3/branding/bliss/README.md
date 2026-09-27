# Bliss

This is a special theme, for bliss-go a Linux installer for blissDS.

## Prerequisites

bliss theme use 7zip and zenity or dialog, so they MUST to be present on the
system.

I use to have a `bliss-go-root` folder with the follow directory inside
`/updates/blissos/update.zip` and `blissos`, this `bliss-go-root` will be copied
on the root of the machine before to create the image.

`sudo scp -r artisan@192.168.1.2:/home/artisan/bliss-go-root/* /`

## Usage

If you download theme with command: `eggs wardrobe get`, then you can write just
`sudo eggs produce --theme bliss`

You can write the full path too: `sudo eggs produce --theme ./my-theme-bliss`
