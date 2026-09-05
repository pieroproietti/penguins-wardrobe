---
title: Guida a Penguins' Wardrobe e Tailor
authors: pieroproietti
lang: it
sidebar_position: 3
enableComments: true
---

# Guida a Penguins' Wardrobe e Tailor

Con **penguins-wardrobe** e **penguins-tailor**, creati da Piero Proietti, puoi partire da un'installazione Linux minimale e vestirla con un desktop, i programmi e le configurazioni che desideri. Il sistema così preparato può essere usato direttamente oppure rimasterizzato con penguins-eggs.

| Progetto | Ruolo |
| --- | --- |
| Penguins' Wardrobe | Contiene ricette YAML, costumi, accessori, branding e script |
| Penguins' Tailor | Fornisce `tailor`: scarica l'atelier e applica le ricette al sistema |
| Penguins' Eggs | Fornisce `eggs`: rimasterizza il sistema e avvia gli installer |

## Il sarto, il guardaroba e i costumi

**Tailor è il sarto**: legge la ricetta e veste il sistema. **Wardrobe è il guardaroba**, o atelier: raccoglie gli abiti e il materiale necessario a realizzarli.

Un **costume** definisce un allestimento completo, per esempio un desktop XFCE con i suoi programmi, il suo sfondo e la disposizione dei pannelli. Un **accessorio** raccoglie una funzione che può servire a più costumi, come gli strumenti per la grafica o la suite da ufficio. Puoi includerlo in un costume oppure applicarlo da solo.

```text
~/.wardrobe/v2/
├── costumes/       costumi completi
├── accessories/    accessori riutilizzabili
├── branding/       aspetto della live e dell'installer
├── scripts/        script condivisi
└── DOCS/           guide
```

Dentro ogni costume o accessorio trovi una ricetta YAML e, quando serve, una cartella **`sysroot/`**. La ricetta indica quali pacchetti installare e quali comandi eseguire; `sysroot/` contiene i file da portare nel sistema: **grafiche, sfondi, icone e configurazioni**.

Durante la vestizione, **il contenuto di `sysroot/` viene copiato sulla radice `/` del sistema**, mantenendo la struttura delle directory. Per esempio, `sysroot/usr/share/backgrounds/mio-sfondo.png` diventa `/usr/share/backgrounds/mio-sfondo.png`. La sezione [Sysroot: dare al sistema il proprio aspetto](#sysroot-dare-al-sistema-il-proprio-aspetto) mostra come prepararla.

Il **branding** raccoglie invece gli asset usati da eggs per il menu di avvio della live e per l'installer. Il costume lo sceglie con la proprietà `branding`. La [guida al branding](branding.md) spiega come prepararlo e come viene utilizzato da Calamares.

## Installare Tailor

### Installazione dal repository di eggs

Per usare Tailor, installa il pacchetto **`penguins-tailor`**, già disponibile nel repository di eggs. Su Debian e derivate, se hai già configurato quel repository, basta:

```bash
sudo apt update
sudo apt install penguins-tailor
tailor version
```

Il comando da usare dopo l'installazione è `tailor`. Puoi proseguire direttamente con [Dal costume alla live](#dal-costume-alla-live).

Se hai già il comando `eggs` ma non hai ancora configurato il suo repository, aggiungilo prima dell'installazione:

```bash
sudo eggs tools repo add
sudo apt update
sudo apt install penguins-tailor
```

### Compilazione dai sorgenti

Se vuoi sviluppare Tailor o compilare una tua versione, i sorgenti si trovano nella [repository penguins-tailor](https://github.com/pieroproietti/penguins-tailor).

Per compilare dai sorgenti servono Git, Make e Go 1.25 o successivo, secondo il `go.mod` attuale:

```bash
git clone https://github.com/pieroproietti/penguins-tailor.git
cd penguins-tailor
make
sudo make install
tailor version
```

Il Makefile usa normalmente `/tmp/tailor-build-dir` per gli artefatti e installa il programma in `/usr/local/bin/tailor`. Eseguire la compilazione come utente normale.

L'applicazione delle ricette richiede anche gli strumenti di sistema che Tailor invoca: APT/dpkg sulle distribuzioni attualmente gestite, Git per l'atelier, rsync per gli overlay e gli strumenti richiesti dagli script del costume.

### Compatibilità attuale

Il backend pacchetti di `tailor wear` è implementato per **Debian e derivate**, come Devuan e Ubuntu. Il controllo iniziale riconosce anche la famiglia Arch, ma la creazione del gestore pacchetti si interrompe perché il relativo backend non è ancora implementato. La presenza di ricette `arch.yaml` non garantisce quindi l'esecuzione di `wear` su Arch.

Gli strumenti di packaging e configurazione dei repository hanno un ambito distinto: supportare un formato di pacchetto non equivale a supportare l'applicazione dei costumi su quella distribuzione.

## Dal costume alla live

```bash
# Scarica l'atelier o aggiorna quello già configurato
tailor get

# Esamina la raccolta e un costume
tailor list
tailor show colibri

# Applica il costume su una distribuzione compatibile
sudo tailor wear colibri
```

Verificare il riepilogo, il desktop e i servizi configurati. Se il costume installa un nuovo kernel, riavviare prima di rimasterizzare per utilizzare quello desiderato.

Con penguins-eggs C/Go installato:

```bash
sudo eggs remaster
```

Per installare la live ottenuta su disco, dalla live stessa:

```bash
sudo eggs sysinstall calamares
# Oppure l'installer testuale
sudo eggs sysinstall krill
```

## Comandi Tailor

### Scaricare o aggiornare l'atelier

```bash
tailor get
tailor get https://github.com/pieroproietti/penguins-wardrobe
tailor get https://github.com/charliemartinez/penguins-wardrobe
tailor get --url https://github.com/pieroproietti/penguins-wardrobe --branch main
```

`get` accetta un URL posizionale oppure `-u, --url`, e un ramo con `-b, --branch`. Accetta anche un URL con suffisso `#ramo`.

Senza URL esplicito, mantiene l'origine Git dell'atelier esistente; al primo utilizzo sceglie quello di Piero Proietti. Se l'origine coincide, aggiorna con Git. Se si specifica un'origine diversa, **sostituisce la directory dell'atelier esistente**: salvare prima eventuali modifiche locali.

La directory è normalmente `~/.wardrobe` dell'utente reale. Con privilegi elevati Tailor cerca l'utente attraverso `SUDO_USER`, il login e gli utenti del sistema. Cerca poi la raccolta in `~/.wardrobe/v2`, oppure nella radice di `~/.wardrobe`. La directory locale `./v2` è un fallback soltanto se la raccolta installata non esiste.

### Esaminare costumi e accessori

```bash
tailor list
tailor show colibri
tailor show accessories/multimedia
```

`list` elenca i costumi. `show` presenta metadati, distribuzioni dichiarate, accessori e un'anteprima dei pacchetti. Per la lista completa e il branding selezionato leggere anche il file YAML della ricetta.

### Applicare una ricetta

```bash
sudo tailor wear colibri
sudo tailor wear accessories/multimedia
sudo tailor wear colibri --linear
sudo tailor wear colibri --branch main
```

| Opzione | Effetto |
| --- | --- |
| `--linear` | Output lineare al posto dell'interfaccia a schermo diviso |
| `-b, --branch` | Recupera/seleziona il ramo dell'atelier prima dell'applicazione |
| `-n, --dry-run` | Simula l'applicazione della ricetta |

Gli alias `--no-split` e `--simulate` sono ancora accettati, ma nascosti nell'help. Usare i nomi della tabella nei nuovi esempi.

Per una simulazione:

```bash
sudo tailor wear colibri --dry-run --linear
```

La simulazione salta installazione dei pacchetti della ricetta, preseed, overlay, script, attivazione del branding e sincronizzazione della home. **Il codice attuale esegue comunque `apt-get update` all'inizio**, prima di elaborare la ricetta; può inoltre recuperare l'atelier e scrivere log/report. Per questo l'esempio usa `sudo`: senza privilegi la simulazione può fermarsi sull'aggiornamento degli indici. `tailor show` permette invece di ispezionare la ricetta senza avviare `wear`.

### Compilazione, repository ed esportazione

| Comando | Uso |
| --- | --- |
| `tailor version` | Mostra la versione |
| `tailor tools build` | Compila e prepara il pacchetto nativo dalla repository di Tailor; eseguire senza sudo |
| `sudo tailor tools repo add` | Configura il repository dei pacchetti |
| `sudo tailor tools repo rm` | Rimuove la configurazione del repository |
| `tailor export pkg` | Esporta pacchetti verso lo storage remoto configurato |
| `tailor export log` | Esporta i log |

`tools repo` presuppone che Tailor sia già disponibile. Per i parametri di destinazione consultare `tailor export log --help`; espone `--user`, `--ip` e `--dir`. `export --clean` abilita la pulizia degli artefatti precedenti sul server. Usare `tailor <comando> --help` per l'help della versione installata.

## Scrivere una ricetta

Costruiamo un piccolo costume chiamato `my-desktop`: installerà XFCE e porterà nel sistema i nostri sfondi, le icone e la configurazione del desktop. Gli accessori seguono la stessa struttura, sotto `v2/accessories/`.

### Passo 1: preparare le cartelle

Nell'atelier locale crea la directory del costume:

```bash
mkdir -p ~/.wardrobe/v2/costumes/my-desktop
cd ~/.wardrobe/v2/costumes/my-desktop
mkdir -p sysroot/etc/skel/.config
mkdir -p sysroot/usr/share/backgrounds/my-desktop
mkdir -p sysroot/usr/share/icons/my-desktop
```

Il costume potrà contenere:

```text
v2/costumes/my-desktop/
├── index.yaml
├── packages.yaml          facoltativo
├── packages.preseed       facoltativo
├── sysroot/               facoltativo
│   └── etc/skel/.config/
└── scripts/               facoltativo
```

### Passo 2: scegliere programmi e accessori

Salva questo esempio in `index.yaml`. È una ricetta per Debian Trixie:

```yaml
name: my-desktop
description: "Desktop XFCE per Debian Trixie"
author: "Il tuo nome"
release: "1.0.0"
distributions:
  - trixie

sequence:
  repositories:
    update: true
  packages:
    - xfce4
    - lightdm
    - lightdm-gtk-greeter
    - firefox-esr
  accessories:
    - multimedia

finalize:
  cmds:
    - "systemctl enable lightdm"
```

L'esempio omette `branding`: applicarlo rimuove l'eventuale branding attivo. Per selezionarne uno, aggiungere per esempio `branding: quirinux` al livello principale e assicurarsi che esista `v2/branding/quirinux/`.

### Passo 3: aggiungere l'aspetto e le configurazioni

Metti lo sfondo in `sysroot/usr/share/backgrounds/my-desktop/`, le icone in `sysroot/usr/share/icons/my-desktop/` e le configurazioni del desktop in `sysroot/etc/skel/.config/`. Prepara soltanto i file necessari al tuo costume, seguendo gli esempi della sezione successiva.

### Passo 4: provare il costume

```bash
tailor show my-desktop
sudo tailor wear my-desktop --dry-run --linear
sudo tailor wear my-desktop
```

La simulazione conserva il comportamento descritto nella sezione dei comandi: aggiorna comunque gli indici APT iniziali. Dopo la vestizione, entra nella sessione desktop e verifica programmi, sfondo, icone e pannelli. Quando il risultato è quello desiderato, puoi creare la live con `sudo eggs remaster`.

## Sysroot: dare al sistema il proprio aspetto

La cartella `sysroot/` è una piccola riproduzione della struttura del sistema Linux. Durante `tailor wear`, Tailor ne copia **il contenuto su `/`**. La cartella `sysroot` stessa non viene creata nella destinazione: i percorsi al suo interno diventano percorsi reali del sistema.

Per esempio:

```text
my-desktop/sysroot/
├── etc/
│   ├── skel/
│   │   └── .config/
│   │       └── xfce4/             configurazione di pannelli e desktop
│   └── lightdm/
│       └── lightdm-gtk-greeter.conf
└── usr/
    └── share/
        ├── backgrounds/
        │   └── my-desktop/
        │       └── sfondo.png
        └── icons/
            └── my-desktop/
                └── logo.png
```

La corrispondenza è diretta:

| File o directory nel costume | Destinazione durante la vestizione |
| --- | --- |
| `sysroot/usr/share/backgrounds/my-desktop/sfondo.png` | `/usr/share/backgrounds/my-desktop/sfondo.png` |
| `sysroot/usr/share/icons/my-desktop/logo.png` | `/usr/share/icons/my-desktop/logo.png` |
| `sysroot/etc/lightdm/lightdm-gtk-greeter.conf` | `/etc/lightdm/lightdm-gtk-greeter.conf` |
| `sysroot/etc/skel/.config/xfce4/` | `/etc/skel/.config/xfce4/` |

In questo modo il costume può portare con sé l'aspetto che hai preparato: sfondi, loghi, temi, icone, impostazioni del gestore di accesso e configurazione del desktop. Copiare uno sfondo lo rende disponibile; per farlo apparire sul desktop devi includere anche la configurazione che lo seleziona. Lo stesso vale per un tema di icone.

### Il ruolo di /etc/skel

`/etc/skel` contiene le configurazioni iniziali destinate ai nuovi utenti. Nel costume puoi prepararle sotto `sysroot/etc/skel/`, includendo anche i file nascosti come `.config` e `.bashrc`.

Al termine della vestizione Tailor sincronizza `/etc/skel` anche nella home dell'utente individuato, normalmente quello che ha invocato `sudo`. Per esempio, con l'utente `piero`:

```text
sysroot/etc/skel/.config/xfce4/
    → /etc/skel/.config/xfce4/
    → /home/piero/.config/xfce4/
```

Le impostazioni del costume arrivano così anche all'utente esistente. I file omonimi possono essere sovrascritti: scegli le configurazioni che vuoi effettivamente distribuire, senza copiare indiscriminatamente tutta la tua home.

### Può contenere di tutto, ma senza esagerare

`sysroot/` non è limitata alla grafica: può contenere file destinati a qualsiasi parte del filesystem. Proprio per questo conviene mantenerla piccola e comprensibile, con ciò che caratterizza il costume.

Per i programmi usa le liste dei pacchetti; per le operazioni di configurazione usa gli script quando opportuno; in `sysroot/` metti i file che vuoi distribuire. Evita di trascinarci cache, file temporanei, credenziali personali o intere directory di sistema copiate dalla macchina di lavoro. Ogni file incluso deve avere uno scopo chiaro.

Anche gli accessori possono avere una propria `sysroot/`. Tailor applica prima quelle degli accessori e poi quella del costume: quando due file hanno la stessa destinazione, quello del costume prevale. La copia aggiunge e sovrascrive i file forniti; non elimina dalla destinazione tutti gli altri file assenti dal costume.

Nei costumi precedenti puoi trovare il nome `dirs/`: Tailor lo usa se `sysroot/` non esiste. Nei nuovi costumi usa `sysroot/`.

## Riferimento della ricetta

### Campi della ricetta

| Campo | Significato |
| --- | --- |
| `name`, `description`, `author`, `release` | Identità e metadati della ricetta |
| `distributions` | Elenco dei codename ammessi, confrontati con `VERSION_CODENAME` |
| `branding` | Nome della directory del bundle, per i costumi |
| `sequence.repositories` | Configurazione repository e aggiornamenti |
| `sequence.packages` | Pacchetti ordinari |
| `sequence.packages_no_install_recommends` | Pacchetti da installare senza raccomandati APT |
| `sequence.packages_interactive` | Pacchetti che richiedono risposte interattive |
| `sequence.cmds` | Comandi intermedi, prima degli accessori del costume |
| `sequence.accessories` | Accessori da applicare nell'ordine dichiarato |
| `finalize.cmds` | Comandi conclusivi della ricetta |
| `reboot` | Richiede il riavvio al termine dell'applicazione reale |
| `display_manager_notice` | Mostra l'avviso specifico relativo a LightDM di Quirinux |

Il parser accetta anche `packages`, `accessories` e `cmds` al livello principale. Nei nuovi costumi usare la struttura `sequence`/`finalize`. `finalize.customize` viene letto, ma il percorso attuale di `wear` applica gli overlay e sincronizza `/etc/skel` indipendentemente da quel valore: non usarlo come interruttore per disabilitare queste operazioni.

Se `distributions` è assente, non viene imposto un elenco di codename. Se il sistema non espone `VERSION_CODENAME`, il controllo viene saltato con un avviso. Un costume incompatibile viene rifiutato; può fornire `tailor-check` (o il precedente `wardrobe-check`) per mostrare istruzioni, senza aggirare il rifiuto.

### Quale YAML viene letto

Tailor cerca prima `index.yaml`, poi `index.yml`. In loro assenza cerca file basati su identità della distribuzione, derivazione e famiglia, quindi applica ulteriori fallback come `debian.yaml` e `arch.yaml`. Seleziona il primo file trovato: **non fonde più ricette per distribuzione**. La presenza di un fallback non certifica la compatibilità dei suoi pacchetti con l'host.

`packages.yaml` o `packages.yml` aggiunge invece pacchetti alla ricetta selezionata, evitando di aggiungere nomi già presenti. Può contenere una lista semplice, `packages:` oppure `sequence.packages:`:

```yaml
packages:
  - git
  - rsync
```

### Repository e preseed

`sequence.repositories` espone `sources_list` per i componenti APT, `sources_list_d` per comandi di configurazione di repository aggiuntivi, e i booleani `update` e `upgrade`. I comandi devono essere appropriati alla distribuzione dichiarata dalla ricetta.

Per rispondere alle domande Debconf, collocare `packages.preseed` accanto al file YAML del costume o dell'accessorio:

```text
lightdm shared/default-x-display-manager select lightdm
```

Tailor carica le risposte con `debconf-set-selections` prima dei pacchetti della ricetta. Questo meccanismo riguarda Debconf; i pacchetti inseriti in `packages_interactive` mantengono le proprie richieste interattive.

### Script e ordine di applicazione

Il flusso principale di un costume è:

1. Aggiornamento iniziale degli indici, caricamento della ricetta e controlli di compatibilità e branding.
2. Preseed, configurazione repository, verifica degli header del kernel e installazione dei pacchetti del costume.
3. `sequence.cmds`, poi accessori con relativi pacchetti, overlay e comandi.
4. Tentativo di recupero dei pacchetti falliti.
5. Overlay del costume e `finalize.cmds` del costume.
6. Sostituzione o rimozione del branding attivo.
7. Sincronizzazione di `/etc/skel` nella home dell'utente individuato, riepilogo e report.

Applicare un altro costume non disinstalla automaticamente tutti i pacchetti del precedente.

Gli accessori si indicano con il nome, con `accessories/<nome>` oppure con un percorso relativo alla ricetta (`./` o `../`). Per un accessorio applicato direttamente, usare la forma esplicita `tailor wear accessories/<nome>`.

Per gli script, Tailor cerca il primo elemento del comando come file relativo alla directory della ricetta; se lo trova, lo esegue da quel percorso. Se non sono specificati argomenti, aggiunge il nome della ricetta come argomento. Non assumere che la directory corrente dello script sia quella della ricetta.

## Log e diagnosi

Il log tecnico ordinario è `/var/log/tailor/tailor-<identità-sistema>.log`. Il report è normalmente `/var/log/tailor/tailor-report-<identità-sistema>-<data-ora>.txt`; se la directory non è scrivibile il report può finire nella directory temporanea. Il riepilogo indica il percorso effettivo.

Il report distingue pacchetti installati, non disponibili e falliti. Alcuni errori di pacchetti e script vengono registrati senza fermare tutta la vestizione: verificare il report e il sistema risultante anche quando il comando raggiunge il riepilogo.

| Problema | Controllo |
| --- | --- |
| Costume non trovato | Atelier effettivo, directory `costumes/`, nome e file YAML |
| Ricetta locale ignorata | `~/.wardrobe` ha precedenza sul fallback `./v2` |
| Distribuzione incompatibile | `distributions` e `VERSION_CODENAME` del sistema |
| Errore del gestore pacchetti su Arch | Backend `wear` non ancora implementato |
| Simulazione bloccata da APT | L'aggiornamento iniziale richiede privilegi anche con `--dry-run` |
| Pacchetto non disponibile | Nome per la distribuzione e repository configurati |
| Branding precedente o incompleto | Seguire la [diagnosi del branding](branding.md#diagnosi) |

## Contribuire

Creare costumi in `v2/costumes`, accessori riutilizzabili in `v2/accessories` e identità visive in `v2/branding`. Documentare nella ricetta le distribuzioni provate. Verificare YAML, file referenziati, risultato della vestizione e, per il branding, avvio della live e dell'installer.

Per modificare il programma, lavorare nella [repository Tailor](https://github.com/pieroproietti/penguins-tailor); questa repository ospita le ricette e la guida comune. L'[atelier Quirinux](https://github.com/charliemartinez/penguins-wardrobe), mantenuto da Charlie Martínez, segue lo stesso modello di separazione.

Autore: Piero Proietti. Conservare le licenze e i crediti dei singoli asset oltre alla [licenza della repository](../../LICENSE).
