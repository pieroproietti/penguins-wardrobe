# Warbler plasma/wayland/waydroid

This is the second version of this experiment, I want create a slim
ISO to be used for develpment.

There is a motivation on the lacks of sddm in this system, unfortunately
sddm bring dependencies from X11 almost in this version I'm using 
from Debian/bookworm .

# That you get

* waydroid
* lineage-18.1-20230121-VANILLA-waydroid_x86_64 system.img (766 MB)
* lineage-18.1-20230121-MAINLINE-waydroid_x86_64 vendor.img (165 MB)
* vscode, node, git and necessary to develop
* firmwares wifi cards (you can ask for more additions)

# how to install
The live is autologin, just look the instructions in console:

```sudo eggs install --unattended```

# how to use
```startplasma-wayland``` to get the GUI, click on waydroid to run it.

# how to remaster
This system can be easily remastered with eggs:

```sudo eggs produce --fast```

You will get an iso of your system, without you user data.

```sudo eggs produce --clone```

You will get an iso of your system, with you user data.

enjoy!

piero.proietti@gmail.com
