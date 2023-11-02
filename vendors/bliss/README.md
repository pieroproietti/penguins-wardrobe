# Bliss

This is a special theme, for bliss-go a Linux installer for blissDS.

## Prerequisites
bliss theme use 7zip and zenity or dialog, so they MUST to be present on the system.

I use to have a `bliss-go-root` folder with the follow directory inside `/updates/blissos/update.zip` and `blissos`, this `bliss-go-root` will be copied on the root of the machine before to create the image.

`sudo scp -r artisan@192.168.1.2:/home/artisan/bliss-go-root/* /`


## Usage

If you download theme with command: `eggs wardrobe get`, then you can write just `sudo eggs produce --theme bliss`

You can write the full path too: `sudo eggs produce --theme ./my-theme-bliss`


## calamares final settings

As special theme, include `calamares final settings`

* cfs-install 
* cfs-data-img
* cfs-bootloader

cfs work actually on Arch, Debian and Ubuntu. They are copied to `/usr/lib/x86_64-linux-gnu/calamares/modules/` by the code in [focal.ts](https://github.com/pieroproietti/penguins-eggs/blob/4f1b9c537a2e182b5a5b89c09f22821e0f6195d0/src/classes/incubation/distros/focal.ts#L98) on penguins-eggs or in others places depending of the distro/codename.

cfs can be used with calamares and krill (CLI installer) too, on krill they are evaluated from [krill-sequence.tsx](https://github.com/pieroproietti/penguins-eggs/blob/4f1b9c537a2e182b5a5b89c09f22821e0f6195d0/src/krill/krill-sequence.tsx#L630).

You can find this modules inside the folder `./theme/calamares/calamares-modules` and will be copied under `/usr/lib/x86_64-linux-gnu/calamares/modules/` building an ISO using the theme.

Note: prefix `cfs-` is mandatory, they are usually named named from the action they do, for example `cfs-data-img`

# no more changes necessary on settings.conf

**cfs-modules** will be automatically included in standard `/etc/calamares/settings.conf`, and executed immidiatly before umount. There is no more need to register them in settings.conf, they will be immidiatly added in `/etc/calamares/settings.conf` before `umount`.

