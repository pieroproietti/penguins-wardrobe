# penguins-wardrobe

A more complete costumes wardrobe for penguin's eggs

## Costumes

A costume consist in a directory named after the costume and an index.yml file. 

It is defined in i-materis.ts from eggs:
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

this is an example from colibri:
```
# wardrobe: .
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
    - anydesk
    - firefox-esr
  accessories:
    - base
    - eggs-dev
    - firmwares
    - xfce4
customize:
  dirs: true
  scripts:
    # desktop_background_set.sh
    - desktop_link_set.sh
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


### customize
costomize contain the final actions to finalize the installation and customize the result. Can contain:
* dirs
* scripts








You can add others scripts and directories inside:

- dirs (contain all file and directories you want to be copied to your /)

## Directories example: 
You need to copy something in ```/etc/skel```, in ```/usr/share/applications``` and your background on ```/usr/share/background```. You can 
put all in ```dirs```:

```
- dirs   + etc   + skel  
         + usr   + share   + applications + install-to-waydroid.desktop
                           + backgrounds  + waydroid
```

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

