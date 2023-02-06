# penguins-wardrobe

It is a repository mainly consisting of .yaml files and simple bash scripts used by eggs to create customizations of Linux systems starting from a minimal image - referred to as "naked" - to achieve a complete system.

I have used this approach both for creating some personal customizations and for creating three Linux systems with waydroid named: wagtail, warbler, and whipbird based on wayland and respectively on gnome, KDE Plasma, and weston.

![wagtail-warbler-whipbird](./DOCUMENTATION/images/wagtail-warbler-whipbird.png)

# The methafora

The methafora consist in a wardrobe, costumes and accessories. 

![wardorbe](./DOCUMENTATION/images/wardrobe/51616859915_5f8eaabfa4_w.jpg)

penguins-wardrobe is a repository for costumes and accessories, managed using Git and organized into directories for costumes, accessories, and documentation. The wardrobe allows for easy organization and consolidation of Linux customizations, making it simple to find and reuse previous work.

It is possible to dress the CLI system with a beautiful GUI interface, but it can also be used as a server without a GUI. This method has proven useful for development and organization

## Costumes and accessories

A costume consist in a directory named after the costume and an index.yml file. 

The type of content of a costume is actually defined in [i-materia.ts](https://github.com/pieroproietti/penguins-eggs/blob/master/src/interfaces/i-materia.ts) from [penguins-eggs](https://github.com/pieroproietti/penguins-eggs):

Here we have an example of costume definition from my colibri, I use it to generate my personal developer working station. You can find a more updated version here; [index.yaml](https://github.com/pieroproietti/penguins-wardrobe/blob/main/costumes/colibri/index.yml)


```
# wardrobe: .
# costume: /colibri
---
name: colibri
...
  hostname: true
reboot: true
```

### General informations
```
name: colibri
description: >-
  desktop xfce4 plus all that I need to develop eggs, firmwares and anydesk
  repos
author: artisan
release: 0.0.3
```

### sequence
sequence is the crucial part of costumes and accessories, this is executed in that sequence and the idea is to make it the more possible atomic. Can contain:

* repositories
  * sources_list
  * sources_list_d
* preinst
* dependencies
* packages
* packages_no_install_recommends
* try_packages
* try_packages_no_install_recommends
* debs
* packages_python
* accessories
* try_accessories

the idea back sequence is it to make it the more possible atomic.

#### repositories
That you need to add to your ```/etc/apt/surces.list``` and, mostly, on ```/etc/aot/surces.list.d``` directory.

repositories consist in two item:

* ```sources_list``` component names to be used: main, contrib, non-free 
* ```sources_list_d``` commands to add other repositories in /etc/apt/sources.list.d

#### preinst
Sometime we need to make same action before to install packages, we can add here one or more scripts to do that actions.

#### packages
A simple array of packages to be installed.

#### packages_no_install_recommenfs
A simple array of packages to be installed with option ```--no-install-recommends```.

#### debs
This is a boolean field, if it's true, the content of the directory ./debs with be installed via the command: ```dpkg -i ./debs/*.deb```.

#### packages_python
A simple array of python packages to be installed with pip.

#### accessories
List of accessories to be installed to complete the costume. eg:
```
accessories:
- base
- eggs-dev # defined in /accessory
- waydroid #
- ./firmware # here we will look to an accessory defined inside the costume
```

### customize
costomize contain the final actions to finalize the installation and customize the result. Can contain:
* dirs
* hostname
* scripts

##### dirs
dirs is boolean, if true, directory ./dirs inside the costume will be copied to / of the system,

example: 
You need to copy something in ```/etc/skel```, in ```/usr/share/applications``` and your background on ```/usr/share/background```. You can 
put all in ```dirs```:

```
- dirs   + etc   + skel  
         + usr   + share   + applications + install-to-waydroid.desktop
                           + backgrounds  + waydroid
```
##### hostname
is boolean too, and if true, /etc/hostname will take the name of the costume and /etc/hosts will be changed in accord

##### scripts
scripts contain an array of one or more scripts to be used to customize the result.

You can add others scripts and directories inside the costume or use general scripts under ```../../scripts/``` like ```../../scripts/config_desktop_link.sh```.

**Scripts examples**

Scripts are called from ```customize/scripts``` and executed on the specific order.

- install-image-from-local.sh (a script to copy sustem.img and vendor.img from local)
- no-hw-accelleration.sh (script to set waydroid with no-hw-accelleration)

# Accessories
accessories can be defined inside a costume or outside. External accessories live in ```./accessories```, internal accessories live inside the costume or in others accessories who declare them.

They have the same structure of costume and are called recursively from the costume it self. You can see it as a belt to dress with your pants or a bag associated to your chotches.

Accessories can be installed alone or called from costumes. 

For example: waydroid is an accessory and it's used by wagtail (gnome3) or warbler (KDE), the same is eggs_dev who is added in colibri, or firmware added to duck amd owl who need a good hardware compatibility.

wagtail and warbler are made for developers, so have just an internal firmwares accessory filled mostly with wifi cards.

__Note__ ```sudo eggs wardrobe wear``` accept a flag ```--no_firmwares``` to skip it, in the case we are building for virtual machines or tests.

(*) wagtail. warbler and wispbird before to end the installation run also a special script [add_wifi_firmwares.sh](https://github.com/pieroproietti/penguins-wardrobe/blob/main/scripts/add_wifi_firmwares.sh) to add in Debian bookworm firmwares from bullseye.


## wardrobe get

```
eggs wardrobe get
```

Clone the [penguins-wardrobe](https://github.com/pieroproietti/penguins-wardrobe) in ```~/.wardrobe```, the command accept argument [REPO] so, you can work with your personal wardrobe too. For example:

```
eggs wardrobe https://github.com/quirinux-so/penguins-wardrobe
```

will get in ```~/.wardrobe``` the quirinux version.

## wardrobe list
List costumes and accessoried in wardrobe.

```
eggs wardrobe list 
```

## wardrobe show COSTUME
Show index.yml of the costume.
```
eggs wardrobe show colibri --wardrobe ../my-own-wardrobe
```

## sudo wardrobe wear COSTUME
Start the process of wear a costume, at the end your system will be modified in accord of.

```
sudo eggs wardrobe wear colibri 
```

# Actual costumes:
* ```colibri``` is a light XFC4 for developers you can easily start to improve eggs.
* ```duck``` come with cinnamon - probably is the right desktop for peoples coming from windows - here complete plus office, gimp and vlc
* ```owl``` is a XFCE4 for graphics designers, this a simple/experimental bird, based on the work of Charlie Martinez [quirinux](https://blog.quirinux.org/)

* ```wagtail```, a wayland/Gnome/waydroid installation;
* ```warbler```, a wayland/KDE/waydroid installation;
* ```whipbird```: a wayland/weston/waydroid installation.

# Accessories
* base
* eggs-dev
* firmwares
* graphics
* liquorix
* multimedia
* office
* waydroid

# Themes

A theme for eggs, let you to customize the apparence of the live image. You can customize the boot of live images made with eggs using themes.

An eggs theme is simply an organized set of files and directories, YAML being the most widely used.

```educaandos``` was the first example of an external eggs theme available, others themes are: neon, telos, ufficiozero and waydroid imported from the previous themes included inside eggs before.

## analizing a theme

A theme consists of a simple folder under addons, called with the name of vendor (in the example: blissos), that includes:

```
educaandos/
    theme
        applications
        artwork
        calamares
            branding
            modules
        livecd
```

### applications

Just a desktop link, it will be copied to /usr/share/applications/ and on the Desktop.

### artwork
The icon for your desktop link, will be copied on /usr/share/icons/

### calamares
Contain the calamares customizations and it's by far the most important part of the theme.

Each sample configuration file contains documentation on the options it contains.

The configuration of calamares is specified in yaml files and contains within it documentation for the various options. The main calamares settings.conf file is automatically created by eggs, only partition, locale and users are used here.

Of course, the reference information is in the calamares repository and on the calamares.io website

#### branding
branding.desc file is generated by eggs, please refer to branding for more information

#### modules
* locale.yml
* partitions.yml
* users.yml
* locale.yml (Not used yet)
* partitions.yml (Not used yet)
* users.yml (*)

(*) In ```EducaAndOS``` to ensure administration capability for our user, we need a groups configuration specified in users.yml. Note: If it is not specified as thema educaandos, the configuration of groups in which the live user and the installed system user will be part, will not have the possibility of administration.

* livecd
Take cure of the apparence of your live.


Usage
```
sudo eggs produce --fast --theme ../path/to/theme
```
example: 

```
sudo eggs produce --fast  --theme .wardrobe/themes/educaandos-plus
```
You can also clone the wardrobe with Git, and take the theme from:
```
sudo eggs produce --fast  --theme /penguins-wardrobe/themes/educaandos-plus
```

# Further plans
A theme is a type of addons for eggs. There are other addons as well: adapt that provides a link to resize the window on a virtual machine, pve that shows on the desktop the link to the local Proxmox VE server, rsupport, etc. They live inside penguins-eggs at the moment.

I hope with time -including through collaborations - to add more possibilities to both addons and themes. You don't necessarily have to be a developer to create a theme, in fact graphic designers, translators, textbook writers or proofreaders are welcome.

You can request me to be added as a collaborator to this repository and thus participate in the development.


# Config

This directory it is used to let a minimal customization of unattended installation. 

```sudo eggs install --unattended``` is equivalent to ```sudo install --custom us``` so it's easy to have more customizations, just fork the repository and make a PR.

For example, you can copy us.yaml to bliss.yaml, change user and password, and get your custom configuration:

```sudo eggs install --custom bliss```

# Thats all folks!


