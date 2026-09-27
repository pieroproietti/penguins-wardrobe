#!/bin/bash

# --- Parametri ---
# Inserire qui l'indirizzo IP del telefono (Gateway) fornito dal provider.
ILIAD_GATEWAY="10.242.162.232"
# -----------------

echo "=== Avvio configurazione tethering USB ==="

# Ricarica la configurazione di rete pulita da /etc/network/interfaces
ifreload -a

# Trova dinamicamente l'interfaccia USB generata dallo smartphone (inizia per 'enx')
TETHER_IFACE=$(ip -br link | grep '^enx' | awk '{print $1}')

if [ -z "$TETHER_IFACE" ]; then
    echo "Errore: Nessuna interfaccia USB trovata. Verifica il cavo e le impostazioni del telefono."
    exit 1
fi

echo "Interfaccia trovata: $TETHER_IFACE"

# Pulisce un eventuale bridge precedente (silenzia gli errori se non esiste)
ip link del vmbr1 > /dev/null 2>&1

# Crea il bridge in RAM e attiva le interfacce fisiche
echo "Creazione bridge provvisorio vmbr1..."
ip link add name vmbr1 type bridge
ip link set "$TETHER_IFACE" up
ip link set vmbr1 up

# Collega l'interfaccia USB al nuovo bridge
ip link set "$TETHER_IFACE" master vmbr1

# Rilascia vecchi lease e richiede l'assegnazione del nuovo IP dinamico
echo "Negoziazione DHCP in corso..."
dhclient -r vmbr1 2>/dev/null
dhclient -v vmbr1

# Pulisce le rotte di default conflittuali e forza l'instradamento verso il cellulare
echo "Impostazione della rotta di default verso $ILIAD_GATEWAY..."
ip route del default > /dev/null 2>&1
ip route add default via "$ILIAD_GATEWAY" dev vmbr1

echo "========================================="
echo "Configurazione completata! Test di connessione in corso:"
ping -c 4 8.8.8.8
