---
title: Wardrobe users' guide
authors: pieroproietti
lang: it
sidebar_position: 3
enableComments: true
---
import Translactions from '@site/src/components/Translactions';

<Translactions />

`Penguins' wardrobe`, non fa parte di `eggs` e non è necessaria la sua conoscenza ed il suo utilizzo per rimasterizzare il proprio sistema o per creare delle proprie customizzazioni originali. E' piuttosto una metodologia che consente - nel mentre si sviluppa un progetto di rimasterizzazione - di tenere traccia dei passi effettuati e riutilizzarli all'occorrenza. 

Si tratta di un repository [penguins-wardrobe](https://github.com/pieroproietti/penguins-wardrobe) costituito principalmente da file yaml e da semplici  script bash ed utilizzato da `eggs` attraverso il comando `wardrobe` per creare personalizzazioni di sistemi Linux a partire da una sistema minimale CLI già installato, nella nostra terminologia `naked`.

Ho utilizzato questo metodo sia per la creazione di alcune personalizzazioni generiche: `colibri`, `duck`, `owl` e `eagle`, sia per testare configurazioni su architetture diverse (arm64).

Potete SEMPRE utilizzare `eggs` a partire da vostre customizzazioni originali senza utilizzare il comando `wardrobe` ma seguendo le vostre metodologie.

## La metafora del guardaroba

La metafora consiste in un guardaroba contenente `costumes` ed `accessories` per la vestizione.

![wardrobe](/img/wardrobe/51616859915_5f8eaabfa4_w.jpg)

`penguins-wardrobe` è una repository, gestita con `git` ed organizzata per directory: `costumes`, `accessories`, `config`, `vendors`, `servers` e `documentation`. Il suo utilizzo prevede di effettuare il fork della stessa al fine di creare customizzazioni proprie, permette così sia una facile organizzazione del nostro lavoro, sia il consolidamento delle nostre esperienze, rendendo più semplice cercare e riutilizzare del lavoro lavoro fatto in precedenza o importare lavori fatti da terzi.

E' possibile vestire un sistema CLI con un splendida interfaccia GUI, ma è anche possibile utilizzare il wardrobe per collezionare configurazioni server (`servers`) senza necessariamente una interfaccia grafica. 

Questo metodo si è dimostrato utile per lo sviluppo e l'organizzazione del lavoro.

## `costumes`

Un `costume` consiste essenzialmente in una directory denominata con il nome del `costume` e contenente dei file yaml specifici per la distribuzione, ad esempio `debian.yaml`, `arch.yaml` o `alpine.yaml`.

Questi file specificano la composizione del `costume`, e vengono analizzati dalla classe [`tailor.ts`](https://github.com/pieroproietti/penguins-eggs/blob/master/src/classes/tailor.ts) di `eggs`. In sostanza, noi forniamo le indicazioni ed il sarto ci cuce il vestito.

In `wardrobe`, quindi abbiamo soltanto delle informazioni che specificano soprattutto repository e pacchetti in linguaggio yaml. Questo esempio è tratto dal file `debian.yaml` del mio `colibri`.

```yaml
# wardrobe: .
# costume: /colibri
---
name: colibri
description: >-
  desktop xfce4 plus all that I need to develop eggs, firmwares
author: artisan
release: 0.9.1
distributions:
  - bullseye
  - bookworm
  - trixie
  - chimaera
  - daedalus
sequence:
  repositories:
    sources_list:
      - main
      - contrib
      - non-free
      - non-free-firmware
    update: true
    upgrade: true
  packages:
    - firefox-esr
    - libxfce4ui-utils
    - lightdm
...
  accessories:
    - base
    - eggs-dev
finalize:
  customize: true
  cmds:
    - ../../scripts/config_desktop_link.sh
    - ../../scripts/config_lightdm.sh
    - ../../scripts/config_g4.sh
    - ../../scripts/hostname.sh
    - plymouth-set-default-theme bgrt
    - update-initramfs -u
reboot: true
```

La sintassi utilizzata è yaml, piuttosto semplice da leggere, mentre per la scrittura potete contare su numerosi addon per praticamente ogni editor.

Andiamo a vedere come è composto il file yaml di definizione (es. `debian.yaml`).

Possiamo suddividerlo in tre parti:
* [`header`](#header)
* [`sequence`](#sequence)
* [`finalize`](#finalize)

### `header`
Definisce il nome, l'autore, descrizione e la release del costume. Una parte importante è `distributions`: se la distribuzione corrente non è inclusa (o compatible), il costume non verrà applicato.

```yaml
name: colibri
description: >-
  desktop xfce4 plus all that I need to develop eggs, firmwares
author: artisan
release: 0.9.1
distributions:
  - bullseye
  - bookworm
```

### `sequence`
La `sequence` è la parte cruciale sia dei `costumes` che degli `accessories`, viene eseguita in sequenza - da qua il nome - e l'idea è stata di renderla minima ed indivisibile. Può contenere:

* [`repositories`](#repositories)
  * [`sources_list`](#sources_list)
  * [`sources_list_d`](#sources_list_d)
* [`preinst`](#preinst)
* [`packages`](#packages)
* [`packages_no_install_recommends`](#packages_no_install_recommends)
* [`try_packages`](#try_packages)
* [`try_packages_no_install_recommends`](#try_packages_no_install_recommends)
* [`debs`](#debs)
* [`packages_python`](#packages_python)
* [`accessories`](#accessories)
* [`try_accessories`](#try_accessories)

L'idea dietro la `sequence` è stata quella di renderla il più possibile atomica.

Vediamo come è composta in dettaglio la sequenza.

#### `repositories`
Definisce cosa abbiamo bisogno sia nella nostra `/etc/apt/sources.list` e - principalmente - nella directory `/etc/apt/sources.list.d`.

`repositories` è formata da due item:

##### `sources_list`
Specifica dei componenti da utilizzare: **`main`**, **`contrib`**, **`non-free`**, **`non-free-firmware`**

##### `sources_list_d`
Specifica dei commandi per aggiungere altre repository all'interno di `/etc/apt/sources.list.d`.

#### `preinst`
Abbiamo a volte la necessita di eseguire alcune azioni prima dell'installazione dei pacchetti, possiamo aggiungere queste azioni in forma di script in questa sezione.

#### `packages`
Un semplice array di pacchetti da installare, è il cuore del sistema.

#### `packages_no_install_recommends`
Una semplice array di pacchetti da installare con l'opzione `--no-install-recommends`.

#### `try_packages`
Si comporta come [packages](#packages) ma non fallisce se non trova il pacchetto.

#### `try_packages_no_install_recommends`
Funziona come [packages_no_install_recommends](#packages_no_install_recommends) ma non fallisce se non trova il pacchetto.

#### `debs`

Questo è un campo booleano, se è true il contenuto della directory `./debs` sarà installato con il comando `dpkg -i ./debs/*.deb`.

#### `packages_python`
Un semplice array di pacchetti `python` che saranno installati con `pip`.

#### `accessories`
Una lista di accessori da installare per completare il `costume`. esempio:

```yaml
accessories:
- base
- eggs-dev # defined in /accessory
- ./firmwares # here we will use an accessory defined inside the costume, note ./
```

#### `try_accessories`
Come  [`accessories`](#accessories) ma non fallisce.

### `finalize`
`finalize` contiene le azioni per finalizzare l'installazione e customizzare il risultato. Può contenere:
* [`customize`](#customize)
* [`cmds`](#cmds)

##### `customize`
`customize` è un campo booleano. Se impostato a `true`, il contenuto della directory `sysroot` (o `dirs`) interna al `costume` verrà copiato nella `root` del sistema.

Esempio: Abbiamo bisogno di copiare la nostra customizzazione del desktop `/etc/skel` ed il nostro background su `/usr/share/background`. 

Possiamo organizzare la cartella `sysroot` del nostro costume riproducendo il percorso:
```
sysroot/
  etc/
    skel/
  usr/
    share/
      backgrounds/
```
Tutto il contenuto verrà copiato nel sistema target.

##### `cmds`
`cmds` contiene un `array` di comandi o script utilizzati per customizzare il risultato finale.

Potete aggiungere script e directory all'interno del `costume` o utilizzare gli script comuni contenuti nella cartella `scripts` come `../../scripts/config_desktop_link.sh`.

**Esempi**

I comandi vengono eseguiti nell'ordine specifico.
```yaml
finalize:
  customize: true
  cmds:
    - ../../scripts/config_desktop_link.sh
    - ../../scripts/config_lightdm.sh
    - update-initramfs -u
```

### `reboot`
Se vero in sistema verrà riavviato dopo la vestizione.

## `Accessories`
Gli accessori possono essere definiti all'interno di un `costume` o fuori dallo stesso - principalmente in `accessories` ma anche interni ad un altro `costume`. Gli accessori interni vivono all'interno di un `costume` o di altri accessori che li dichiarano.

Hanno la stessa struttura dei `costumes` e vengono richiamati da questi. Potete vederli come una cintura da mettere con dei pantaloni alla moda o una borsa associata al tailler.

Gli `accessories`, però sono installabili anche singolarmente, ad esempio: `sudo eggs wardrobe wear accessories/eggs-dev` ci installerà gli strumenti di sviluppo utilizzati per creare `eggs`. 

Tutto quello che risiede sotto `accessories` è un accessorio.

## I comandi
Abbiamo solo quattro comandi `wardrobe`: `get`, `list`, `show` e `wear`.

### `wardrobe get [REPO]`

```bash
eggs wardrobe get
```

Esegue il clone di [penguins-wardrobe](https://github.com/pieroproietti/penguins-wardrobe) in `~/.wardrobe`, il comando accetta un argomento [REPO], così potete lavorare con il vostro `wardrobe` invece di quello standard. Ad esempio:

```bash
eggs wardrobe https://github.com/quirinux-so/penguins-wardrobe
```

scaricherà in  `~/.wardrobe` la versione del `wardrobe` di `quirinux`.

Il consiglio per utilizzare [`penguins-wardrobe`](https://github.com/pieroproietti/penguins-wardrobe) per i propri progetti è che - dopo i primi passi - conviene crearsi un proprio `fork` del progetto ed utilizzare quello.

### `wardrobe list`
Mostra la lista dei `costumes` ed `accessories` presenti in `~/.wardrobe`.

```bash
eggs wardrobe list 
```

### `wardrobe show COSTUME`
Mostra l'indice di un `costume` presente in `~/.wardrobe`.

```bash
eggs wardrobe show colibri
```

### `sudo wardrobe wear COSTUME`
Avvia il processo di "vestizione" di un `costume`. I pacchetti del `costume` verranno installati, così come degli `accessories` interni o esterni, infine verrà eseguito `finalize` ed alla fine del processo il sistema sarà modificato secondo le istruzioni del `costume` stesso.

```bash
sudo eggs wardrobe wear colibri 
```

## Costumi esistenti
Alcuni dei costumi presenti nel wardrobe:
* `albatros`: costumi per varie architetture.
* `chicks`: una configurazione leggera e giocosa, adatta all'education.
* `colibri`: è un Desktop XFCE4 leggero che utilizzo per sviluppare `eggs` e che fornisce un ambiente completo.
* `duck`: utilizza `cinnamon` e probabilmente è il giusto Desktop per utilizzatori abituati all'interfaccia di Windows - inoltre fornisce `libreoffice`, `gimp`, `vlc` e vari.
* `eagle`: varie configurazioni, incluse versioni per arm64.
* `gypaetus`: configurazioni base.
* `owl`: basato su XFCE4 per grafici e designer.
* `seagull`: altre varianti.

## `accessories`
Al momento su `penguins-wardrobe` sono presenti vari `accessories`, tra cui:
* `base`: aggiunge spice-vdagent e configurazioni base.
* `educational`: software didattico.
* `eggs-dev`: repositories e tool di sviluppo (nodejs, vscode, etc).
* `firmwares`: installazione firmware.
* `firmwares-light`: versione ridotta dei firmware.
* `flatpak`: supporto flatpak.
* `graphics`: pacchetti grafici.
* `kvm`: il necessario per far girare VM.
* `liquorix`: installa il kernel liquorix.
* `live-installer`: installer grafico.
* `multimedia`: selezione di pacchetti multimediali.
* `nextcloud`: client nextcloud.
* `office`: libreoffice, gimp, vlc.
* `waydroid`: installer per waydroid.

## `vendors`
I `vendors` (precedentemente chiamati `themes`) rendono possibile la customizzazione dell'immagine `live`. E' possibile sia customizzare il boot delle immagini ISO che l'aspetto dell'installer GUI `calamares`.

Come per i `costumes` e gli `accessories`, i `vendors` sono costituiti da una directory, dei file yaml, delle icone, sfondi e quant'altro necessario.

Un `vendor` contiene tipicamente una directory `theme`. Ad esempio `vendors/educaandos-plus/theme` contiene:
* `applications`: link .desktop
* `artwork`: icone
* `calamares`: configurazioni per l'installer Calamares (branding, moduli locale, partition, users).
* `livecd`: aspetto del boot live (GRUB, Isolinux theme).

Esempio di utilizzo:
```bash
sudo eggs produce --theme vendors/educaandos-plus
```

## `servers`
La directory `servers` contiene configurazioni per sistemi server, ad esempio `lamp` (Linux, Apache, MySQL, PHP). Funzionano in modo analogo ai `costumes` ma sono orientati a servizi backend piuttosto che a interfacce desktop.

## `Config`
Questa directory è utilizzata per permettere una customizzazione delle opzioni per l'installazione `--unattended`. 

```bash
sudo eggs install --unattended
```
è equivalente a `sudo install --custom us` in questo modo è relativamente facile avere customizzazioni diverse.

## Nota su Arch Linux e Manjaro
Il wardrobe supporta diverse distribuzioni grazie ai file specifici (es. `arch.yaml`, `alpine.yaml`). Sebbene nato su Debian, la struttura permette di definire pacchetti e comandi adatti anche ad altri gestori di pacchetti (pacman, apk) all'interno dello stesso costume o in costumi dedicati.

## Utilizzo dell'AI per creare Costumes
La struttura dichiarativa in YAML dei `costumes` e degli `accessories` si presta magnificamente all'utilizzo dell'Intelligenza Artificiale. Potete chiedere ad assistenti come **Antigravity** di generarvi un nuovo costume specificando semplicemente:

"Crea un costume wardrobe per Debian bookworm con desktop Mate, firefox, vlc e i tools di sistema standard".

L'AI sarà in grado di produrre un file `debian.yaml` valido e pronto per essere testato, o di analizzare un costume esistente per aggiungere funzionalità.
