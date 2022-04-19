# wardrobe (guardaroba)

Ognuno di noi ha l'esperienza di un guardaroba, il classico armadio in camera da letto da cui prendere i vestiti e gli accessori per vestirsi.

![wardorbe](./images/wardrobe/51616859915_5f8eaabfa4_w.jpg)


Ho voluto utilizzare questa metafora perchè semplice da ricordare, un ```costume``` da indossare, uno o più ```accessories``` da aggiungere. 

Gli accessori possono sia essere definiti nel costume e nel "tiretto" ```accessories```

Il nostro armadio è a due ante, da una parte i costumi dall'altra le definizioni di accessori. 

Volendo essere larghi ed esagerati lo potremmo definire a tre ante includendo anche la DOCUMENTATION che state leggendo.

La nostra struttura del wardrobe è questa:

```
~/wardrobe ─┬─ accessories
            ├─ costumes 
            └─ DOCUMENTATION
```

# wardrobe

Naturalmente occorre avere un wardrobe per cercare i costumi e per utilizzarlo. 

## wardrobe get [REPO]
Scarica il wardrobe specificato in REPO che può essere una comune repository git. 

Esempio:

```eggs wardrobe get```

Se non viene specificata alcuna repository verrà utilizzata la repository communidy [penguins-wardrobe](https://github.com/pieroproietti/penguins-wardrobe)


## wardrobe list [--w wardrobe]

## wardrobe show COSTUME [-w wardrobe]

## wardrobe wear COSTUME [-w wardrobe]
