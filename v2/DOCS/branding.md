# Branding: da Wardrobe a Tailor e Calamares

Il branding è un bundle di file selezionato dal costume. Wardrobe conserva gli asset, Tailor installa il bundle attivo e penguins-eggs lo usa durante la creazione della live e la preparazione dell'installer. Questa pagina descrive l'integrazione con penguins-eggs C/Go verificata nei sorgenti del 5 settembre 2026.

Per installazione, comandi e formato dei costumi, consultare la [guida comune](wardrobe-users-guide.md).

## Selezione e installazione

Nel file `index.yaml` del costume:

```yaml
name: my-desktop
branding: quirinux
```

Il nome deve identificare una directory direttamente sotto `v2/branding/`. Il bundle può essere condiviso da più costumi e non deve avere lo stesso nome del costume.

```text
v2/branding/quirinux/                          sorgente nell'atelier
           │ tailor wear my-desktop
           ▼
/etc/penguins-eggs.d/branding/                 bundle attivo sul sistema
```

Tailor copia il contenuto del bundle, senza conservare una sottocartella `quirinux`. Prepara prima una copia temporanea, rimuove il branding precedente e rinomina la copia nella destinazione. La selezione viene applicata dopo overlay e comandi conclusivi del costume.

| Operazione | Risultato |
| --- | --- |
| Costume con `branding: quirinux` | Sostituisce il bundle attivo con Quirinux |
| Costume senza `branding` o con valore vuoto | Rimuove `/etc/penguins-eggs.d/branding` |
| Accessorio applicato direttamente | Mantiene il branding attivo |
| `wear --dry-run` | Non modifica il branding |
| Nome di bundle inesistente | Interrompe l'applicazione del costume al controllo del branding |

Aggiornare l'atelier con `tailor get` non aggiorna automaticamente il bundle già installato. Il normale flusso per installare la nuova versione è applicare nuovamente il costume; questa operazione ripete anche la vestizione, non soltanto la copia del branding.

## Struttura di un bundle

```text
v2/branding/my-brand/
├── applications/
│   └── install-system.desktop.tmpl
├── artwork/
│   └── install-system.png
├── livecd/
│   ├── splash.png
│   ├── grub.theme.cfg
│   ├── isolinux.theme.cfg
│   ├── grub.main.cfg                  opzionale
│   └── isolinux.main.cfg              opzionale
└── calamares/
    ├── branding/
    │   ├── logo.png
    │   ├── welcome.png
    │   ├── show.qml
    │   ├── slide1.png
    │   └── branding.desc              solo se completo e necessario
    └── modules/                       opzionale
        └── users.yaml
```

Le directory e i file sono facoltativi secondo ciò che si desidera personalizzare. Un bundle contenente soltanto immagini per Calamares può usare il descrittore generato da eggs. I file predefiniti installati da eggs si trovano in `/etc/penguins-eggs.d/branding.default/`.

## Come viene generato branding.desc

All'avvio di `eggs sysinstall calamares`, eggs ricrea `/etc/penguins-eggs.d/installer.d/` e prepara l'ambiente. La preparazione è condivisa anche con Krill.

Per il branding di Calamares l'ordine è:

1. Estrazione della base incorporata nell'installer.
2. Copia degli asset da `/etc/penguins-eggs.d/branding.default/calamares/branding/`.
3. Generazione di `branding.desc` dal `branding.desc.tmpl` predefinito, usando `/etc/os-release` e la versione di eggs. Se la directory predefinita non esiste, viene usato il template incorporato.
4. Copia ricorsiva dei file da `/etc/penguins-eggs.d/branding/calamares/branding/` sopra il risultato.

Il file letto dall'installer è quindi:

```text
/etc/penguins-eggs.d/installer.d/branding/eggs/branding.desc
```

`settings.conf` seleziona `branding: eggs` e cerca sotto `installer.d/branding`. Anche un branding personalizzato deve usare:

```yaml
componentName: eggs
```

La sovrapposizione avviene **per file interi**, senza unione dei campi YAML. Un `branding.desc` del bundle sostituisce integralmente quello generato. Eggs non elabora un eventuale `branding.desc.tmpl` del bundle come template personalizzato: il rendering riguarda il template predefinito.

### Usare i testi generati dal sistema

Omettere `branding.desc` dal bundle e fornire gli asset desiderati:

- `logo.png` per icona e logo del prodotto;
- `welcome.png` per l'immagine di benvenuto;
- `show.qml` e le immagini che richiama per lo slideshow.

Il descrittore predefinito usa `NAME`, `PRETTY_NAME` e gli URL di `/etc/os-release`, con fallback verso penguins-eggs. La versione visualizzata proviene dalla versione di eggs. Il bundle Quirinux della raccolta utilizza questo approccio.

Un file chiamato soltanto `neon-logo.png` non sostituisce `logo.png`: occorre uniformare il nome dell'asset oppure fornire un descrittore completo che lo richiami.

### Fornire testi e stile personalizzati

Per controllare direttamente nome del prodotto, versione, URL e stile, includere un `branding.desc` completo. Come base usare il descrittore generato da una live della stessa versione di eggs, adattare i valori e conservare `componentName: eggs`.

I percorsi delle immagini e dello slideshow devono corrispondere ai file presenti nella directory finale. Conservare anche le relative licenze e i crediti.

**Non inserire un file vuoto o un segnaposto come `# this is only a stub`.** Anche quel file viene copiato sopra il descrittore valido, eliminandone la configurazione. Se non serve personalizzare i testi, il file deve essere assente.

Le modifiche manuali a `installer.d` vengono perse al successivo avvio dell'installer, perché eggs rigenera la directory. Conservare le modifiche nel bundle dell'atelier e installarle sul sistema da rimasterizzare.

## Configurazioni dei moduli Calamares

Il percorso opzionale `calamares/modules/` contiene configurazioni applicate dopo quelle generate da eggs. Il codice copia i file `.yaml` e `.conf`, trasformando l'estensione di destinazione in `.conf`:

```text
branding/calamares/modules/users.yaml
    → installer.d/modules/users.conf
```

Anche qui la sostituzione riguarda il file completo, non singoli campi. Non fornire contemporaneamente `users.yaml` e `users.conf`, che hanno la stessa destinazione. I file `.yml` e le sottodirectory non vengono elaborati da questo overlay. Questi file controllano il comportamento dell'installazione e vanno verificati sulla live di destinazione.

## Boot della live

Durante la rimasterizzazione eggs utilizza gli asset in `livecd/` per splash e temi GRUB/ISOLINUX. Il passaggio `boot-assets` sceglie il bundle attivo se contiene `splash.png`; altrimenti usa `branding.default/livecd` anche per i temi. Non è una sovrapposizione per file come quella di Calamares: per un tema live personalizzato fornire anche lo splash. I menu di avvio sono distinti dal descrittore Calamares.

Il generatore dei menu usa la directory `branding/livecd` quando esiste, altrimenti `branding.default/livecd`. Se nella directory selezionata sono presenti **entrambi** `grub.main.cfg` e `isolinux.main.cfg`, li elabora come menu personalizzati. Altrimenti genera i menu standard, usando i temi predisposti nella ISO.

Nei menu personalizzati vengono sostituiti soltanto questi segnaposto:

| Segnaposto | Valore |
| --- | --- |
| `{{{fullname}}}` | `PRETTY_NAME` del sistema |
| `{{{kernel}}}` | Versione del kernel in esecuzione |
| `{{{vmlinuz}}}` | `/live/vmlinuz` |
| `{{{initrdImg}}}` | `/live/initrd.img` |
| `{{{kernel_parameters}}}` | Parametri di boot forniti dalla pipeline |
| `{{{rmModules}}}` | Stringa vuota |

Questa è una sostituzione limitata di valori, non un interprete Mustache completo. I parametri indispensabili al boot devono essere conservati nei menu personalizzati.

## Collegamento dell'installer

`applications/install-system.desktop.tmpl` personalizza il lanciatore preparato nella live. Un esempio essenziale:

```ini
[Desktop Entry]
Type=Application
Name=Install system
Exec=pkexec eggs sysinstall
Icon=install-system
Terminal=@@TERMINAL@@
Categories=System;
```

Eggs sostituisce `@@TERMINAL@@` in base alla disponibilità di Calamares e normalizza la riga `Exec` a `pkexec eggs sysinstall`. `artwork/install-system.png` fornisce l'icona personalizzata del collegamento. Questo template non usa la sintassi Go del `branding.desc.tmpl` predefinito.

## Diagnosi

Controllare i tre livelli: **bundle nella repository → copia installata → file generati nella live**. Un cambiamento nella repository non implica che sia già presente negli altri due livelli.

| Sintomo | Cosa controllare |
| --- | --- |
| Calamares perde la configurazione con un bundle specifico | Il suo `branding.desc` è completo oppure contiene soltanto commenti? |
| Compare il logo predefinito | Il bundle ha `logo.png`, oppure il descrittore richiama il nome corretto? |
| Il branding non viene selezionato correttamente | `componentName: eggs` e struttura `calamares/branding/` |
| Testi della distribuzione base con immagini personalizzate | Comportamento previsto se manca un descrittore personalizzato |
| Slideshow incompleto | File richiamati da `show.qml`, import QML e timer dello slideshow |
| Le modifiche scompaiono riavviando Calamares | Sono state fatte in `installer.d`, che viene rigenerata |
| Rimane una vecchia selezione | Confrontare `branding` nel costume e bundle realmente installato |
| Un modulo non cambia | Estensione `.yaml` o `.conf`, file completo e destinazione corretta |

Per l'avvio tramite `eggs sysinstall calamares`, consultare `/var/log/calamares-install.log`. Un errore prima del lancio di Calamares può invece comparire direttamente nell'output di eggs durante la preparazione dell'ambiente.

## Riferimenti al codice

- Tailor: `pkg/tailor/branding.go` e `pkg/tailor/wear.go` nella [repository penguins-tailor](https://github.com/pieroproietti/penguins-tailor).
- Eggs: `coa/pkg/sysinstall/setup/branding-desc.go`, `calamares-overlay.go` e `orchestrator.go` nella [repository penguins-eggs](https://github.com/pieroproietti/penguins-eggs).
- Live: `coa/brain.d/base.yaml.tmpl` e `coa/pkg/assets/configs/scripts/generate-menus.sh` nella stessa repository eggs.
