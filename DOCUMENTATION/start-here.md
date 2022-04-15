# start here

It's long time I enjoy using Linux, both - professional and hobbist  - Started about 2000 with Red Hat and Debian, switched to Debian with Knoppis, using Proxmox Ve since version 0.7.

I was really fashinated from knoppix, some so nice to use and install who you can browse the we during the installations, at that time was common not.

With remastersys and - later with systemback - I arranged my customization too, not only for work, not only for hobbies.

Same year later, with the end of systemback and experience buildin a pxe server in our office, I started to think about some to replace them, from this idea come eggs.

eggs is capable of remaster, take a system and get a live iso. With the time, and collaborations, I thouth to something to master not only re-master. 

wardrobe is an answer of the needs everyone trying to create on Debian, Devuan, Ubuntu respin need to know, use and extend.

# wardrobe
A wardrobe is a repository managed with git containing various definitions of costumes and accessories.

# costume

A costume consist in a directory named after the costume and an index.yml file. 

The king of content of costumes is actually defined in i-materia.ts from eggs:

```
export interface IMateria {
   name: string
   author: string
   description: string
   release: string
   distributions: string []
   sequence: {
       repositories: {
           sources_list: string []
           sources_list_d: string []
           update: boolean
           upgrade: boolean
       },
       preinst: string[]
       packages: string []
       packages_no_install_recommends: string []
       debs: boolean
       packages_python: string []
       accessories: string[]
  }
  customize: {
    dirs: boolean
    hostname: boolean
    scripts: string []
  }
  reboot: boolean
}
```

this is an example from my colibri, the configuration of my system as  developer:

```
# wardrobe: https//github.com/pieroproietti/penguins-wardrobe
# costume: /colibri
---
name: colibri
description: >-
  desktop xfce4 plus all that I need to develop eggs, firmwares and anydesk
  repos
author: artisan
release: 0.0.3
distributions:
  - bookworm
  - bullseye
  - buster
  - chimaera
  - focal
  - impish
  - jammy
sequence:
  repositories:
    sources_list:
      - main
      - contrib
      - non-free
    sources_list_d:
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
    - anydesk
    - firefox-esr
    - libxfce4ui-utils
    - lightdm
    - network-manager-gnome
    - network-manager-openvpn
    - network-manager-openvpn-gnome
    - qt5ct
    - tango-icon-theme
    - thunar
    - xarchiver
    - xfce4-appfinder
    - xfce4-panel
    - xfce4-pulseaudio-plugin
    - xfce4-session
    - xfce4-settings
    - xfce4-terminal
    - xfce4-whiskermenu-plugin
    - xfconf
    - xfdesktop4
    - xfwm4
  accessories:
    - base
    - eggs-dev
    - firmwares
customize:
  dirs: true
  scripts:
    - desktop_link_set.sh
  hostname: true
reboot: true

```


# let's analyze the content of index.yml
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
* packages
* packages_no_install_recommends
* debs
* packages_python
* accessories

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
accessories are accesspries, can be internal or external. External accessories live in ./aceessories, internal accessories live inside the costume o accossory who declare it.

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
is boolean too, and if present hostname will take the name of the costume and hosts will be changed in accord

##### scripts
cripts contain an array of one or more scripts to be used to customize the result







You can add others scripts and directories inside:


## Scripts examples:

Scripts are called from customizations.scripts and executed on the specific order.

- install-image-from-local.sh (a script to copy sustem.img and vendor.img from local)
- no-hw-accelleration.sh (script to set waydroid with no-hw-accelleration)

# Accessories
An accessory it'a costume who live under accessories directory. You can see it as a belt to dress with your pants or a bag associated to your chotches.
Accessories are used alone or from costumes, for example: waydroid is an accessoru and it's used gwaydroid, kwaydrois, the same for firmware who is added in hen, gwaydroid, kwaydroid and all future costumes who need a good hardware compatibility.


# clone wardrobe

Cloning wardrobe let you to change the size and materials of the costumes to better fit your need.

```
git clone https://github.com/pieroproietti/penguins-wardrobe
```

## list costumes
```
eggs wardrobe list --wardrobe ./penguins-wardrobe
```

## show costume
```
eggs wardrobe show --wardrobe ./penguins-wardrobe --costume colibri
```

## wear costume
```
eggs wardrobe wear --wardrobe ./penguins-wardrobe --costume colibri
```

## ironing costume
ironing put in alphabetic order most of the sections of costumes. You can use it to get a screenshot of ordered costume and use it to put order in your costume.

```
eggs wardrobe ironing --wardrobe ./penguins-wardrobe --costume colibri
```

# Costumes

## colibri
colibri is a light XFC4 for developers you can easily start to improve eggs.

## duck
duck come with cinnamon - probably is the right desktop for peoples coming from windows - here complete plus office, gimp and vlc

## owl
owl is a XFCE4 for graphics designers, this a simple/experimental bird, based on the work of Clarlie Martinez quirinux

## wagtail
GNOME waydroid installation.

## warbier
KDE waydroid installation

# Accessories

## base
## cinnamon
## eggs-dev
## firmwares
## office
## waydroid
## XFCE4

