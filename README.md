# penguins-wardrobe

It is a repository mainly consisting of .yaml files and simple bash scripts used by eggs to create customizations of Linux systems starting from a minimal image - referred to as "naked" - to achieve a complete system.

I have used this approach both for creating some personal customizations and for creating three Linux systems with waydroid named: wagtail, warbler, and whipbird based on wayland and respectively on gnome, KDE Plasma, and weston.

![wagtail-warbler-whipbird](./DOCUMENTATION/images/wagtail-warbler-whipbird.png)

# The methafora

The methafora consist in a wardrobe, costumes and accessories. 

A wardrobe is a common repository managed with git organized with costumes, accessories and DOCUMENTATION dirs. Inside a wardrobe You can arrange your costumes collection to wear your penguins. 

I found this method usefull to develop, organize and consolidate my Linux customizations and a compact way to have all the operations in the same place, find and reuse them when I needs.

Of course you can dress your CLI system with wonderfull GUI interface, but it is also possible to wear it as server without a GUI at all.

## Costumes and accessories

A costume consist in a directory named after the costume and an index.yml file. 

The type of content of a costume is actually defined in i-materia.ts from eggs:
```
  name: string
  author: string
  description: string
  release: string
  distributions: string[]
  sequence: {
    repositories: {
      sources_list: string[]
      sources_list_d: string []
      update: boolean
      upgrade: boolean
    },
    preinst: string[]
    dependencies: string[]
    packages: string[]
    packages_no_install_recommends: string[]
    try_packages: string[]
    try_packages_no_install_recommends: string[]
    debs: boolean
    packages_python: string[]
    accessories: string[]
    try_accessories: string[]
  }
  customize: {
    dirs: boolean
    hostname: boolean
    scripts: string[]
  }
  reboot: boolean
```

This is an example of costume definition from my colibri, a configuration generating my personal developer working station:

```
# wardrobe: .
# costume: /colibri
---
name: colibri
description: >-
  desktop xfce4 plus all that I need to develop eggs, firmwares and anydesk
  repos
author: artisan
release: 0.9.1
distributions:
  - bookworm
  - bullseye
  - buster
  - chimaera
  - focal
  - jammy
sequence:
  repositories:
    sources_list:
      - main
      - contrib
      - non-free
    sources_list_d:
      - rm -f /usr/share/keyrings/anydesk-stable.gpg
      - >-
        curl -fsSL "https://keys.anydesk.com/repos/DEB-GPG-KEY" | gpg --dearmor
        -o /usr/share/keyrings/anydesk-stable.gpg
      - >-
        echo "deb [signed-by=/usr/share/keyrings/anydesk-stable.gpg]
        http://deb.anydesk.com/ all main" | tee
        /etc/apt/sources.list.d/anydesk-stable.list > /dev/null    
    update: true
    upgrade: true
  packages:
    - adwaita-qt
    # anydesk
    - libxfce4ui-utils
    - lightdm
    - lightdm-autologin-greeter # ubuntu seem to need it, Debian install it automatically
    - lightdm-gtk-greeter # mandatory for ubuntu, without it will install gdm3
    - qt5ct
    - spice-vdagent
    - tango-icon-theme
    - thunar
    - xarchiver
    - xfce4-appfinder
    - xfce4-notifyd # va installo altrimenti network-manager-gnome richiama gnome-shell
    - xfce4-panel
    - xfce4-pulseaudio-plugin
    - xfce4-session
    - xfce4-settings
    - xfce4-terminal
    - xfce4-whiskermenu-plugin
    - xfconf
    - xfdesktop4
    - xfwm4
  packages_no_install_recommends:
    - network-manager-gnome
    - network-manager-openvpn
    - network-manager-openvpn-gnome
  try_packages:
    - firefox
    - firefox-esr
  accessories:
    - base
    - eggs-dev
    - flatpak
    - python3-dev
  try_accessories:
    - firmwares
customize:
  dirs: true
  scripts:
    - ../../scripts/config_desktop_link.sh
    - ../../scripts/config_lightdm.sh
    #
    # insert command g4artisan
    #
    - rm -f /usr/local/bin/g4*
    - curl -fsSL "https://raw.githubusercontent.com/pieroproietti/penguins-eggs/master/g4/g4artisan" -o /usr/local/bin/g4artisan
    - chmod +x /usr/local/bin/g4artisan
    - curl -fsSL "https://raw.githubusercontent.com/pieroproietti/penguins-eggs/master/g4/g4clone" -o /usr/local/bin/g4clone
    - chmod +x /usr/local/bin/g4clone
    - curl -fsSL "https://raw.githubusercontent.com/pieroproietti/penguins-eggs/master/g4/g4passwd" -o /usr/local/bin/g4passwd
    - chmod +x /usr/local/bin/g4passwd

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
sequence is the crucial part of costumes, this is executed in that sequence and the idea is to make it the more possible atomic. Can contain:

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
That you need to add / change your .list file in /etc/aot/surces.list and /etc/aot/surces.list.d

repositories consist in two item:

* ```sources_list``` component to be used: main, contrub, non-free 
* ```sources_list_d``` commands to add other repositories in /etc/apt/sources.list.d

#### preinst
Sometime we need to make same action before to install packages, we can add here one or more scripts to do that actions.

#### packages
A simple array of packages to be installed

#### packages_no_install_recommenfs
A simple array of packages to be installed with option no-install-recommenfs

#### debs
This is a boolean field, if it's true, the content of ./debs with be installed via a ```dpkg -i ./debs/*.deb``` command

#### packages_python
A simple array of python packages to be installed with pip

#### accessories
accessories are accessories and can live inside the costume or external. External accessories live in ./aceessories, internal accessories live inside the costume or in others accessories who declare them.

They have the same structure of costume and are called recursively from the costume.


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
scripts contain an array of one or more scripts to be used to customize the result

You can add others scripts and directories inside the costume or use general scripts under ```../../scripts/``` like ```../../scripts/config_desktop_link.sh```

**Scripts examples**

Scripts are called from customize/scripts and executed on the specific order.

- install-image-from-local.sh (a script to copy sustem.img and vendor.img from local)
- no-hw-accelleration.sh (script to set waydroid with no-hw-accelleration)

**Accessories**
An accessory it'a costume who live under accessories directory or inside the costume itself. You can see it as a belt to dress with your pants or a bag associated to your chotches.
Accessories can be installed alone or called from costumes. For example: waydroid is an accessory and it's used by wagtail (gnome3) or warbler (KDE), the same for firmwares who is added in colibri and other species who need a good hardware compatibility.

__Note__ eggs wardrobe wear accept a flag --no_firmwares to skip it, in the case you are building for virtual machines.

# wardrobe get

```
eggs wardrobe get
```

Clone the [penguins-wardrobe](https://github.com/pieroproietti/penguins-wardrobe) in ~/.wardrobe, the command accept argument [REPO] so, you can work with your personal wardrobe too. For example:

```
eggs wardrobe https://github.com/quirinux-so/penguins-wardrobe
```

will get in ~/.wardrobe the quirinux version.

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

## wardrobe wear COSTUME
Start the process of wear a costume, at the end your system will be modified in accord of.

```
eggs wardrobe wear colibri 
```
## wardrobe iroring COSTUME
ironing put in alphabetic order most of the sections of costumes. You can use it to get a screenshot of an ordered index.yml of the costume.

```
eggs wardrobe ironing colibri --wardrobe ./my-own-wardrobe 
```

# Costumes

* **colibri** is a light XFC4 for developers you can easily start to improve eggs.
* **duck** come with cinnamon - probably is the right desktop for peoples coming from windows - here complete plus office, gimp and vlc
* **owl** is a XFCE4 for graphics designers, this a simple/experimental bird, based on the work of Charlie Martinez [quirinux](https://blog.quirinux.org/)
* **warbler**, a wayland/KDE/waydroid installation 

# Accessories
* base
* eggs-dev
* firmwares
* graphics
* liquorix
* multimedia
* office
* waydroid

## Config

This directory it's used from eggs install --custom to let an unattended installation customized for you.

```sudo eggs install --unattended``` is equivalent to ```sudo install --custom un``` so it's easy to have more customizations.

For example, you can copy us.yaml to bliss.yaml and change user and password, if you want, and call your custom configuration with:

```sudo eggs install --custom bliss```

# Themes

A theme for eggs, let you to customize it and save your time.

An eggs theme is simply an organized set of files and directories, YAML being the most widely used.

educaandos is the first example of an external eggs theme available, others themes are: neon, telos, ufficiozero and waydroid imported from the previous themes included in eggs before.

## Themes
You can customize the boot of live images made with eggs using themes.

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

* users.yml

In EducaAndOS to ensure administration capability for our user, we need a groups configuration specified in users.yml. Note: If it is not specified as thema educaandos, the configuration of groups in which the live user and the installed system user will be part, will not have the possibility of administration.

livecd
Take cure of the apparence of your live.

Usage
sudo eggs produce --fast --theme ../path/to/theme
example
clone this repository:

git clone https://github.com/pieroproietti/penguins-addons
And produce your customized live:

sudo eggs produce --fast  --theme ../penguins-addons/educaandos-plus
Further plans
A theme is a type of addons for eggs. There are other addons as well: adapt that provides a link to resize the window on a virtual machine, pve that shows on the desktop the link to the local Proxmox VE server, rsupport, etc.

I hope with time -including through collaborations - to add more possibilities to both addons and themes. You don't necessarily have to be a developer to create a theme, in fact graphic designers, translators, textbook writers or proofreaders are welcome.

You can request me to be added as a collaborator to this repository and thus participate in the development.

