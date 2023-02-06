# penguinwardroe

Si tratta di un un repository costituito principalmente da file .yaml e da semplici  script bash usata da eggs per creare personalizzazioni di sistemi Linux a partire da un'immage - naked - per ottenere un sistema completo.

Ho utilizzato questo metodo sia per la creazione di alcune personalizzazioni
di sistemi Linux con waydroid denominatie wagtail, warbler e whipbird 
rispettivamente su Gnome, KDE Plasma e Weston.

![wagtail-warbler-whipbird](./images/wagtail-warbler-whipbird.png)

# La metafora del guardaroba

La metafora consiste in un guardaroba contenente costumi ed accessori per la vestizione.

![wardorbe](./images/wardrobe/51616859915_5f8eaabfa4_w.jpg)

penguins-wardrobe è una repository per costumes ed accessories, utilizzata usando Git ed organizzata per directory: costumes, accessories, config, themes e documentation. Il wardrobe permette una facile organizzazione e per la consolizzazione delle nostre esperienze di customizzazioni Linux, rendendo semplice cercare e riutilizzare il lavoro precedentemente svolto.

E' possibile vestire un sistema CLI con un splendida interfaccia GUI, ma è anche possibile utilizzare il wardrobe per collezionare configurazioni server senza necessariamente una interfaccia grafica. Questo metodo si è dimostrato utipe per lo sviluppo e l'organizzazione.

## Costumi ed accessori

Un costume consiste un una directory denominata con il nome del costume ed un file index.yaml. 

I tipi dei contenuti del costume sono attualmente definiti in [i-materia.ts](https://github.com/pieroproietti/penguins-eggs/blob/master/src/interfaces/i-materia.ts) su [penguins-eggs](https://github.com/pieroproietti/penguins-eggs):

Qui abbiamo sono delle definizioni di costume ad esempio il mio colibri che utilizzo per il mio lavoro di sviluppatore, potete visualizzare tutto l'esempio direttamente sul wardrobe: [index.yaml](https://github.com/pieroproietti/penguins-wardrobe/blob/main/costumes/colibri/index.yml)


```
# wardrobe: .
# costume: /colibri
---
name: colibri
...
  hostname: true
reboot: true
```

### Informazioni generali
```
name: colibri
description: >-
  desktop xfce4 plus all that I need to develop eggs, firmwares and anydesk
  repos
author: artisan
release: 0.0.3
```

### sequence 
La sequence è la parte cruciale sia dei costumi che degli accessori, viene eseguita in sequenza - da qua il nome - e l'idea è stata di crearla il più possibile atomica. Puà contenere:

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

L'idea diesto la sequenza è stata di renderla minima ed indivisibile.

#### repositories
Definisce cosa abbiamo bisogno sia nella nostra ```/etc/apt/surces.list``` e principalmente, nella directory ```/etc/aot/surces.list.d```.

repositories è formata da due item:

* ```sources_list``` componenti da utilizzare: main, contrib, non-free 
* ```sources_list_d``` commandi per aggiungere altre repository all'interno di ```/etc/apt/sources.list.d```.

#### preinst
Abbiamo a volte la necessita di eseguire alcune azioni prima dell'installazione dei pacchetti, possimo aggiungere queste azioni in forma di script in questa sezione.

#### packages
Un semplice array di pacchetti da installare, è il cuore del sistema.

#### packages_no_install_recommenfs
Una semplice array di pacchetti da installare con l'opzione ```--no-install-recommends```.

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

A theme consists of a simple folder under addons, called with the name of   
vendor (in the example: educanandos), that includes:

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

## More informations
There is a [Penguin's eggs official book](https://penguins-eggs.net/book/) and same other documentation - mostly for developers - on [penguins-eggs repo](https://github.com/pieroproietti/penguins-eggs) under **documents** and **i386**, in particular we have [hens, differents species](https://github.com/pieroproietti/penguins-eggs/blob/master/documents/hens-different-species.md) who descrive how to use eggs in manjaro.

* [blog](https://penguins-eggs.net)    
* [facebook penguin's eggs group](https://www.facebook.com/groups/128861437762355/)
* [telegram penguin's eggs channel](https://t.me/penguins_eggs) 
* [twitter](https://twitter.com/pieroproietti)
* [sources](https://github.com/pieroproietti/penguins-krill)

You can contact me at pieroproietti@gmail.com or [meet me](https://meet.jit.si/PenguinsEggsMeeting)

## Copyright and licenses
Copyright (c) 2017, 2023 [Piero Proietti](https://penguins-eggs.net/about-me.html), dual licensed under the MIT or GPL Version 2 licenses.


