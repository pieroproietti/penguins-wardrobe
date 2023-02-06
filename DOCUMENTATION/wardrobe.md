# wardrobe (guardaroba)

Ognuno di noi ha l'esperienza di un guardaroba, il classico armadio in camera da letto da cui prendere i vestiti e gli accessori per vestirsi.

![wardorbe](./images/wardrobe/51616859915_5f8eaabfa4_w.jpg)


Ho voluto utilizzare questa metafora perchè semplice da ricordare, dei ```costumes``` da indossare, uno o più ```accessories``` da aggiungere. 

Gli accessori possono sia essere definiti all'interno del costume e nel "tiretto" ```accessories```

Il nostro armadio è a due ante, da una parte la definizione dei costumi dall'altra quella degli accessori. 

Volendo essere larghi lo potremmo definire anche a tre o quattro ante includendo anche la DOCUMENTATION  che state leggendo e la directory CONFIG per le configurazioni per installazione unattended.

La nostra struttura del wardrobe è questa:

```
~/wardrobe ─┬─ accessories
            ├─ costumes 
            ├─ config
            ├─ DOCUMENTATION
            ├─ scripts
            ├─ servers
            └─ themes
```

## accessories
Contiene gli accessori dei costumi
## costumes
Contiene la definizione dei costumi
## config
Contiene file di configurazione .YAML utilizzati da ```eggs install --custom file```
## DOCUMENTATION
Contiene questa documentatione

## scripts
Contiene degli script per configurare i link sul desktop ed il desltop manager. Vengono chiamati nella sezione finale per configurare lo sfondo, i link e l'autologin.

## servers
Contiene dei costumi server.

## themes
Contiene i temi di eggs, precedentemente in penguins-addons.


# I comandi wardrobe

## wardrobe get [REPO]

Scarica il wardrobe specificato in REPO che può essere una comune repository git. 

Come abbiamo detto un wardrobe è una semplice repository git, per prima cosa dobbiamo scaricarla in locale. 

Lo possiamo fare con il comando ```git clone https://github.com/pieroproietti/penguins-wardrobe``` e lo faremo per lo sviluppo, però dato che in questo modo avremo sempre la noia di utilizzare un path per wardrobe diverso, per ovviare è stato aggiunto il comando ```eggs wardrobe get``` che scarica la repository [penguins-wardrobe](https://github.com/pieroproietti/penguins-wardrobe) senza trascinarsi indietro tutta la storia ```--depth 1```.

Esempio:

```eggs wardrobe get```

Può essere utilizzato questo comando anche passando come REPO la repository da utilizzare. Esempio:

```eggs wardrobe get https://github.com/quirinux-ga/penguins-wardrobe```

In questo caso verrà impostato come wardrobe di default quello di quirinux, una distribuzione per la grafica, cartoon.


## wardrobe list [--w wardrobe]
Se non viene passato un particolare wardrobe, list andrà a mostrare i ```costume``` e gli ```accessory``` presenti nel wardrobe di default.

## wardrobe show COSTUME [-w wardrobe]

Viene visualizzato in formato yaml o json il contentuto di index.yml del costume in oggetto.


## wardrobe wear COSTUME [-w wardrobe]
Finalmente il comando per la vestizione: legge la definizione del costume e procede all'installazione di quanto necessario per inddossarlo. Alla fine il nostro sistema verrà configurato secondo quanto prescritto nel costume.

Dato che questo comando andrà ad installare pacchetti, effettuare delle copie in aree riservate, etc, necessita dei diritti di root e dovrà essere chiamato con sudo:

```sudo eggs wardrobe wear colibri```

