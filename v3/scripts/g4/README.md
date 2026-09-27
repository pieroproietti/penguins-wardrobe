# g4tools: Interfaccia uomo-macchina per anziani del Terminale (e smemorati)

I file presenti in questa cartella **non sono indispensabili** per il funzionamento, la compilazione o la rimasterizzazione di `penguins-eggs`. Il core del progetto vive, si gestisce e si compila benissimo anche senza di essi.

Allora perché sono qui? 

Sono nati per un'esigenza puramente umana: fungere da ponte e semplificare la vita all'autore (me stesso) e a tutti quegli "smemorati di lungo corso" che, pur avendo speso una vita sui sistemi, preferiscono risparmiare la brillantezza cerebrale per l'architettura del codice piuttosto che sprecarla a ricordare a memoria sfilze di flag, variabili d'ambiente o le contorte sintassi dei comandi di Git, e Proxmox.

---

### 🛠️ La Suite di Orchestrazione p4 (per proxmox)

Per gestire l'ambiente di sviluppo nativo su Proxmox senza l'overhead di Vagrant, `penguins-eggs` include la suite **p4** (Proxmox). Questo set di comandi permette di orchestrare, sincronizzare e controllare le *forge* di test (Alpine, Arch, Debian, Devuan, Fedora, Manjaro, Opensuse ed Ubuntu) direttamente dal terminale locale.

* **`p4autocomplete`**
  Installa l'autocompletamento Bash per l'intera suite nello spazio utente locale (`~/.local/share/bash-completion`). Permette di usare il tasto `TAB` per completare dinamicamente i comandi e i nomi delle distribuzioni target.

* **`p4config`**
  Gestisce la configurazione dell'ambiente. Definisce le coordinate del nodo Proxmox host (`father`), gli ID delle macchine virtuali e i parametri di rete necessari per il collegamento.

* **`p4create`**
  Inizializza o clona una nuova *forge-* (Macchina Virtuale) sull'hypervisor, preparandola per il ciclo di sviluppo.

* **`p4push`**
  Il cuore operativo del sistema. Sincronizza istantaneamente il codice sorgente locale con la *fucina* bersaglio (tramite filesystem condiviso `9p`) e innesca la build nativa del pacchetto (RPM, DEB o PKG), bypassando dinamicamente i blocchi di sicurezza come SELinux.

* **`p4ssh`**
  Apre una connessione diretta alla console della *fucina* (sfruttando la porta seriale `ttyS0` o SSH). Fondamentale per il debug manuale a basso livello e l'ispezione del sistema in tempo reale.

* **`p4push`**
copia i sorgenti di penguins-eggs nella fucina.

* **`p4start`**
  Invia il comando di accensione all'host Proxmox per risvegliare la *fucina* specificata.

---

### 🛠️ Alias miei per git
Git è fantastico, ma alcune operazioni avanzate sui tag o sui rami remoti richiedono comandi lunghi e rischiosi. La serie `g4` standardizza le operazioni ripetitive sul monorepo:

Giusto wrapper per non dover reimpostare i miei default e la mia password auromatica, ma puo essere usato da chiunque basta modificat g4config.

---

## Un consiglio per il viandante
Se sei capitato in questa cartella, sei liberissimo di usare questi script così come sono se il tuo ambiente host è configurato in modo compatibile.

Tuttavia, il consiglio migliore che posso darti è un vecchio principio dell'hacking: **non copiare alla cieca, fatti i tuoi**. Prendi ispirazione da questi helper, adattali ai tuoi ritmi, alle tue dita sul terminale e alle tue abitudini.

La macchina è al nostro servizio, non il contrario. La pigrizia intelligente è la più alta forma di ottimizzazione sistemistica.

---

## 3. Installazione e Uso

Per fare in modo che il tuo terminale riconosca automaticamente tutti i comandi `v4*` e `g4*` da qualsiasi directory, e per assicurarti che il layout della tastiera italiana sia sempre attivo, appendi queste righe in coda al tuo file `~/.bashrc`:

```bash
# Forza il layout della tastiera italiana all'apertura del terminale
setxkbmap it

# Aggiunge la cartella bin del monorepo penguins-eggs al PATH dell'utente
export PATH="$HOME/forge/bin:$PATH"
```
