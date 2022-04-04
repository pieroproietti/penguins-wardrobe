# penguins-wardrobe

A more complete costumes wardrobe for penguin's eggs

# Costumes

A costume consist in a directory named after the costume and an index.yml file. 

You can add others scripts and directories inside:

- dirs (contain all file and directories you want to be copied to your /)

## Directories example: 
You need to copy something in ```/etc/skel```, in ```/usr/share/applications``` and your background on ```/usr/share/background```. You can 
put all in ```dirs```:

```
- dirs   + etc   + skel  
         + usr   + share   + applications + install-to-waydroid.desktop
                           + backgrounds  + waydroid
```

## Scripts examples:

Scripts are called from customizations.scripts and executed on the specific order.

- install-image-from-local.sh (a script to copy sustem.img and vendor.img from local)
- no-hw-accelleration.sh (script to set waydroid with no-hw-accelleration)

# Accessories
An accessory it'a costume who live under accessories directory. You can see it as a belt to dress with your pants or a bag associated to your chotches.
Accessories are used alone or from costumes, for example: waydroid is an accessoru and it's used gwaydroid, kwaydrois, the same for firmware who is added in hen, gwaydroid, kwaydroid and all future costumes who need a good hardware compatibility.


# clone wardrobe

Cloning wardrobe let you to change the size and materials of the costumes to better fit your need.

```
git clone https://github.com/pieroproietti/penguins-wardrobe
```

## list costumes
```
eggs wardrobe list --wardrobe ./penguins-wardrobe
```

## show costume
```
eggs wardrobe show --wardrobe ./penguins-wardrobe --costume gwaydroid
```

## wear costume
```
eggs wardrobe wear --wardrobe ./penguins-wardrobe --costume gwaydroid
```

## ironing costume
ironing put in alphabetic order most of the sections of costumes. You can use it to get a screenshot of ordered costume and use it to put order in your costume.

```
eggs wardrobe ironing --wardrobe ./penguins-wardrobe --costume gwaydroid
```



# Costumes

## Hen
XFCE4 desktop, complete to build eggs.

## kwaydroid
Light KDE configuration and waydroid

## kwaydroid
Light GNOME configuration and waydroid

# Accessories

## base
## eggs-dev
## firmwares
## waydroid

