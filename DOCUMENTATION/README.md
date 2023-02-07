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


__Nota__ ```sudo eggs wardrobe wear``` accetta una flag ```--no_firmwares``` per saltare completamente il firmware nel caso stiamo lavorando per macchine virtuali o facendo dei test.

(*) wagtail. warbler e whispbird prima della fine dell'installazione chiamano pure uno speciale script [add_wifi_firmwares.sh](https://github.com/pieroproietti/penguins-wardrobe/blob/main/scripts/add_wifi_firmwares.sh) per aggiungere in Debian bookworm firmwares proveniente da bullseye.

Gli stessi drivers possono essere scaricati da [non-free]
(http://cdimage.debian.org/cdimage/unofficial/non-free/firmware/bookworm/current/).


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
Mentre costumes ed accessories si applicano ad un sistema installato, i themes lavorano per la customizzazione per l'immagine live. E' possibile customizzare il boot delle immagini iso create con eggs creando ed utilizzando i themes.

Un tema è semplicemente un sistema organizzato di file e directory, anche per i temi .yaml è il linguaggio maggiormente utilizzato.

```educaandos``` è stato il primo esempio di un tema esterno disponibile, altri temi che potete trovare nel wardrobe sono: neon, telos, ufficiozero and waydroid precedentemente presenti direttamente all'interno di eggs.

## analizziamo un tema

Un tema consiste un una semplice directory sotto themes, denominata con il nome del vendor (in questo esempio: educanandos), che include:

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

Solo un link .desktop, verrà copiato in /usr/share/applications/ e sulla cartella Desktop.

### artwork
L'icona per il tuo link -desktop, verrà copiata in /usr/share/icons/.

### calamares
Contiene la configurazione per calamares ed è la parte più importante del tema.

I file di configurazione di capamares sono scritti sempre in yaml e contengono la documantazione per le varie optioni. Il principale file di configurazione settings.conf viene automaticamente generato da eggs, solo partition, locale ed users sono usati da wardrobe.

Per le informazioni di riferimento sulla configurazione di questi file si rimanda al sito di [calamares](https://calamares.is).

#### branding
branding.desc questo file viene generato da eggs, dai riferimento a [branding](https://github.com/calamares/calamares/blob/calamares/src/branding/default/branding.desc) per maggiori informazioni.

#### modules
* locale.yml
* partitions.yml
* users.yml (*)
* locale.yml (Not used yet)
* partitions.yml (Not used yet)


(*) In ```EducaAndOS``` per avere i diritti di amministrazione per l'utente, abbiamo la necessità di configurare lo stesso in un gruppo specifico.

* livecd
Si prende cura dell'aspetto del boot da live.


Usage
```
sudo eggs produce --fast --theme ../path/to/theme
```
esempio: 

```
sudo eggs produce --fast  --theme .wardrobe/themes/educaandos-plus
```

Potete anche clonare il wardrobe con Git e prendere il thema da esso:
```
sudo eggs produce --fast  --theme /penguins-wardrobe/themes/educaandos-plus
```

# Further plans
A theme is a type of addons for eggs. There are other addons as well: adapt that provides a link to resize the window on a virtual machine, pve that shows on the desktop the link to the local Proxmox VE server, rsupport, etc. They live inside penguins-eggs at the moment.

I hope with time -including through collaborations - to add more possibilities to both addons and themes. You don't necessarily have to be a developer to create a theme, in fact graphic designers, translators, textbook writers or proofreaders are welcome.

You can request me to be added as a collaborator to this repository and thus participate in the development.


# Config
Questa directory è utilizzata con una customizzazione minimale per alcune opzioni dell'installazione --unattended. 

```sudo eggs install --unattended``` è equivalente a ```sudo install --custom us``` in questo modo è relativamente facile avere customizzazioni diverse, semplicemente creando un fork di questa repository ed una PR al sottoscritto.

Ad esempio potete copiare us.yaml in bliss.yaml, e cambiare il nome dell'utente e la password ed avere la vostra personalizzazione con:

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


