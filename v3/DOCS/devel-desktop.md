# Vestizione dichiarativa — branch devel

La collezione `v3/` accompagna `penguins-tailor` sul branch `devel`.
I costumi LightDM attivati sono colibri, chicks, eagle, gypaetus, quirinux
(Xfce) e duck (Cinnamon). Albatros/SDDM e Seagull/GDM conservano il percorso
precedente: i rispettivi backend dichiarativi non sono ancora implementati.

```yaml
desktop: xfce
display_manager: lightdm
session_type: x11
init: auto
autologin: true
login_background: /usr/share/backgrounds/colibri/3794764350_2839ca0b26_b.jpg
```

Sysroot rimane la directory `sysroot/` del costume, con fallback legacy a
`dirs/`: non occorre aggiungere un campo YAML per indicarne il percorso.
Le cinque scelte sono dunque desktop, login, tipo di sessione, sysroot e init.

## Significato di init

- `auto` (predefinito): rileva il gestore dei servizi in esecuzione.
- `systemd`: richiede systemd; abilita `lightdm.service`.
- `sysvinit`: richiede SysVinit con `update-rc.d`, come in Devuan.
- `openrc`: richiede OpenRC e `/etc/init.d/lightdm`; aggiunge LightDM al
  runlevel `default`.

Un valore esplicito non installa o sostituisce l'init: se non corrisponde a
quello rilevato, Tailor si ferma. La presenza di `systemctl` da sola non
identifica systemd. Chroot e sistemi offline non sono supportati.
Con SysVinit vengono disabilitati gli script installati degli altri login
manager noti; con OpenRC vengono rimosse le loro voci dal runlevel `default`.
Runlevel OpenRC personalizzati e servizi generici senza uno script LightDM
richiedono ancora configurazione specifica.

## Pacchetti e ordine

I sei costumi mantengono le liste di pacchetti e gli accessori esistenti.
Tailor non aggiunge metapacchetti alle ricette che dichiarano già pacchetti
o accessori. Solo una ricetta minimale senza liste riceve i pacchetti dedotti
da `desktop` e `display_manager`.

Su Debian la verifica della sessione avviene dopo l'installazione degli
accessori: Quirinux fornisce così i componenti Xfce. Gli overlay degli
accessori seguono ancora il ciclo preesistente; la verifica precede l'overlay
del costume. Sulle famiglie senza backend pacchetti, desktop, sessione,
greeter e servizio LightDM devono essere già installati e la verifica
precede tutti gli overlay.

La configurazione dichiarativa del login avviene dopo sysroot e, nel percorso
completo, dopo gli script finali. La nuova abilitazione del servizio non
riavvia la sessione corrente e imposta graphical.target per il prossimo avvio systemd.
Gli script legacy e gli script dei pacchetti possono invece gestire servizi
e target: rimangono parte delle ricette esistenti.

La configurazione di LightDM è ora gestita da Tailor: le sei ricette non
richiamano più lo script shell. Anche i comandi duplicati di attivazione del
display manager nella ricetta Quirinux sono stati rimossi.

`autologin: true` seleziona l'utente non-root che invoca Tailor tramite sudo o
doas (SUDO_USER, poi DOAS_USER, poi l'account corrente). Un'esecuzione diretta
come root non sceglie automaticamente un utente. `false` disabilita autologin
utente e guest nel seat predefinito; omesso mantiene la configurazione esistente.
La sessione autologin segue quella dichiarata e il timeout è zero. Le regole
PAM e le impostazioni specifiche per seat restano applicabili.

`login_background` indica il percorso assoluto dell'immagine installata da
sysroot. Richiede lightdm-gtk-greeter, seleziona questo greeter e ne aggiorna
lo sfondo dopo gli overlay. Uno sfondo personale può prevalere sul default.
Chicks usa la directory `chick/`; Quirinux non dichiara un'immagine e conserva
lo sfondo installato. Un'immagine dichiarata ma assente produce un errore prima
della scrittura della configurazione login.

Su systemd Tailor disabilita i display manager concorrenti abilitati, senza
fermarli o mascherarli, abilita LightDM e seleziona graphical.target.
La configurazione si applica anche nel percorso senza installazione pacchetti.
Le liste curate di pacchetti/accessori e gli altri script delle ricette restano;
Duck ed Eagle dichiarano esplicitamente lightdm-gtk-greeter.

## Sessione e prove

I sei costumi dichiarano X11, coerentemente con la loro configurazione
attuale. Wayland va scelto e provato separatamente con una sessione installata
riconosciuta. `auto` accetta solo una corrispondenza univoca; le preferenze
utente e le sezioni LightDM specifiche per seat possono prevalere sul default.

Da Tailor compilato dal branch `devel`, per usare questa copia locale:

```bash
cd /percorso/penguins-wardrobe
/percorso/tailor-devel wear colibri --dry-run --linear
```

Il dry-run mostra il piano; rinvia alla vestizione reale la verifica di init,
servizi e sessioni installate. Può produrre log e report. Per una prova reale
usare una VM sacrificabile con la distribuzione prevista dal costume.
Non serve pubblicare il branch né usare `wear --branch devel` per la copia
locale: `--branch` gestisce invece il checkout del repository del wardrobe.

La copia locale viene usata solo in assenza di `~/.wardrobe`. Con un wardrobe
installato, Tailor usa prima quella copia e preferisce la sua cartella `v3/`.
