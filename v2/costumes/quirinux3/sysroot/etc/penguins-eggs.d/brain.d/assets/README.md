# Quirinux branding for eggs (coa)

This directory is copied as-is to `/etc/penguins-eggs.d/brain.d/assets/`
when running `coa wardrobe wear quirinux`. From there, it is
automatically used by both `eggs produce` and `eggs sysinstall calamares`.

* `splash.png`: GRUB/ISOLINUX background for the live system. Already
  supported by coa, with no code changes required (see
  `coa/brain.d/base.yaml.tmpl`).

* `calamares/`: contains `branding.desc`, the logo, slideshow
  (`show.qml`, `slide1.png`, and `welcome.png`), and Calamares
  translations (`lang/*.qm`), exactly as provided in
  `eggs-quirinux-config_1.4.11+q2`. Requires the
  `feature/vendor-branding-hook` patch in penguins-eggs.

* `calamares/finish.sh`: a cleaned-up version of the original
  `fin-instalacion` script from Quirinux's legacy wardrobe (Penguins
  Legacy). It renames references to "quirinux" under the real user's
  `.config` and `.local` directories, and enables
  `GRUB_DISABLE_OS_PROBER` for dual-boot installations. The parts that
  coa already handles on its own have been removed (`$CHROOT`
  detection, `update-initramfs`, `update-grub`, and `grub-install`),
  since those are already executed generically by
  `oa-chroot-runner.sh` through `debian.bash.tmpl`. Requires the
  `feature/vendor-finish-step` patch in penguins-eggs.


