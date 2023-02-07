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

I **type** dei contenuti del costume sono attualmente definiti typescript [i-materia.ts](https://github.com/pieroproietti/penguins-eggs/blob/master/src/interfaces/i-materia.ts) su [penguins-eggs](https://github.com/pieroproietti/penguins-eggs).

Qui abbiamo sono delle definizioni di costume in linguaggio YAML. ad esempio il mio colibri che utilizzo per il mio lavoro di sviluppatore, lo potete visualizzare tutto l'esempio direttamente sul wardrobe: [index.yaml](https://github.com/pieroproietti/penguins-wardrobe/blob/main/costumes/colibri/index.yml)


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

L'idea dietro la sequenza è stata di renderla minima ed indivisibile.

#### repositories
Definisce cosa abbiamo bisogno sia nella nostra ```/etc/apt/surces.list``` e principalmente, nella directory ```/etc/apt/sources.list.d```.

```repositories``` è formata da due item:

* ```sources_list``` componenti da utilizzare: **main**, **contrib**, **non-free**
* ```sources_list_d``` commandi per aggiungere altre repository all'interno di ```/etc/apt/sources.list.d```.

#### preinst
Abbiamo a volte la necessita di eseguire alcune azioni prima dell'installazione dei pacchetti, possimo aggiungere queste azioni in forma di script in questa sezione.

#### packages
Un semplice array di pacchetti da installare, è il cuore del sistema.

#### packages_no_install_recommenfs
Una semplice array di pacchetti da installare con l'opzione ```--no-install-recommends```.

#### debs

Questo è un campo booleano, se è true il contenuto della directory ./debs sarà installato con il comando ```dpkg -i ./debs/*.deb```.

#### packages_python
Un semplice array di pacchetti python che saranno installati con pip.

#### accessories
Una lista di accessori da installare per completare il costume. esempio:
```
accessories:
- base
- eggs-dev # defined in /accessory
- waydroid #
- ./firmware # here we will look to an accessory defined inside the costume
```

### customize
costomize contiene le azioni per finalizzare l'installazione e customizzare il risultto. Può contente:
* dirs
* hostname
* scripts

##### dirs
dirs è un campo booleano, se true la directory ./dirs interna al costume verrà copiata nella root del sistema.

Esempio: Abbiamo bisogno di copiare qualcosa in ```/etc/skel```, in ```/usr/share/applications``` ed il nostro background su ```/usr/share/background```. 

Possiamo mettere il tutto in ```dirs```:

```
- dirs   + etc   + skel  
         + usr   + share   + applications + install-to-waydroid.desktop
                           + backgrounds  + waydroid
```
##### hostname
Anche questo è un campo booleano e, se truei, il file ```/etc/hostname``` verrà posto con il nome del costume ed in accordo a ciò sarà modificato anche ```/etc/hosts```.

##### scripts
scripts contiene un array di uno o più script utilizzati per customizzare il risulto.

Potete aggiunge altri script e directory all'interno di del costume o utilizzare gli script generali sotto ```../../scripts/``` come ```../../scripts/config_desktop_link.sh```.

**Scripts examples**

Gli scripts sono chiamati da ```customize/scripts``` ed eseguiti nell'ordine specifico.

- install-image-from-local.sh (a script to copy sustem.img and vendor.img from local)
- no-hw-accelleration.sh (script to set waydroid with no-hw-accelleration)

# Accessories
Gli accessori possono essere definiti all'interno di un costume o risiededere fuori da esso in```./accessories```, gli accessori interni vivono all'interno di un costume o di altri accessori che li dichiarano.

Hanno la stessa struttura dei costumi e sono chiamata ricorsivamente da questi. Potete vederli come una cinta da mettere con deo pantaloni specifici o una borsa associata al proprio tailler.

Gli accessori possono sia essera installati da soli oppure chiamati da un costume. 

Ad esempio: waydroid è un accessoria ed è utilizzato da wagtail (gnome3) e  warbler (KDE), lo stesso per eggs-dev che viene aggiunto in colibri oppure  firmwares in duck ed owl che hanno bisogno di una maggiore compatibilità hardware.

colibri, wagtail, warbler e whispbird sono fatti per sviluppatori, così usano solo un accessorio firmwares interno con soprattutto driver per wifi.

__Note__ ```sudo eggs wardrobe wear``` accetta una flag ```--no_firmwares``` per saltare completamente il firmware nel caso stiamo lavorando per macchine virtuali o facendo dei test.

(*) wagtail. warbler e whispbird prima della fine dell'installazione chiamano pure uno speciale script [add_wifi_firmwares.sh](https://github.com/pieroproietti/penguins-wardrobe/blob/main/scripts/add_wifi_firmwares.sh) per aggiungere in Debian bookworm firmwares proveniente da bullseye.


## wardrobe get

```
eggs wardrobe get
```

Esegue il clone di [penguins-wardrobe](https://github.com/pieroproietti/penguins-wardrobe) in ```~/.wardrobe```, il comando accetta un argomento [REPO] so, così potete lavora con il vostro wardrobe privato. Ad esempio:

```
eggs wardrobe https://github.com/quirinux-so/penguins-wardrobe
```

scaricherà in  ```~/.wardrobe``` la versione di quirinux.

## wardrobe list
Mostra la lista dei costumes ed accessoried presenti nel wardrobe.

```
eggs wardrobe list 
```

## wardrobe show COSTUME
Mostra l'indice index.yml di un costume.
```
eggs wardrobe show colibri --wardrobe ../my-own-wardrobe
```

## sudo wardrobe wear COSTUME
Avvia il processo di vestizione di un costume, alla fine del processo il sistema sarà modificato secondo le indicazioni del costume.

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


