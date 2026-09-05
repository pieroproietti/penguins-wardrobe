# Penguins' Wardrobe

![Penguins' Wardrobe](v2/penguins-wardrobe.png)

Penguins' Wardrobe is Piero Proietti's atelier of Linux system recipes. It contains costumes, accessories, branding bundles, and shared scripts. [Penguins' Tailor](https://github.com/pieroproietti/penguins-tailor) provides the `tailor` command that applies them; [Penguins' Eggs](https://github.com/pieroproietti/penguins-eggs) remasters the configured system into a bootable live ISO.

This repository hosts the shared Wardrobe and Tailor user documentation.

## Documentation

- [Wardrobe and Tailor user guide — Italiano](v2/DOCS/wardrobe-users-guide.md): installation, commands, recipes, package handling, overlays, and troubleshooting.
- [Branding guide — Italiano](v2/DOCS/branding.md): selecting a bundle, live boot artwork, installer launchers, and Calamares configuration.
- [Collection layout](v2/README.md): where recipes and assets belong.

## Start with Tailor

Install Tailor using the [user guide](v2/DOCS/wardrobe-users-guide.md#installare-tailor), then fetch and inspect a costume:

```bash
tailor get
tailor list
tailor show colibri
```

Apply a costume on a compatible Debian-family system:

```bash
sudo tailor wear colibri
```

`wear` installs packages, runs recipe commands, copies system configuration, and synchronizes `/etc/skel` into the target user's home. Compatibility and package choices depend on the selected recipe. The current package backend is APT; Arch recipe files do not yet imply a working Arch `wear` backend.

For a preview, use `sudo tailor wear colibri --dry-run --linear`. This skips recipe application, but the current implementation still runs the initial APT index refresh and may fetch the atelier and write logs or reports.

After checking the configured desktop, use the C/Go Penguins' Eggs CLI to remaster it:

```bash
sudo eggs remaster
```

## What the atelier contains

| Directory | Purpose |
| --- | --- |
| [Costumes](v2/costumes/) | Complete system and desktop recipes |
| [Accessories](v2/accessories/) | Package groups and configuration that costumes can share |
| [Branding](v2/branding/) | Live boot artwork and installer identity |
| [Scripts](v2/scripts/) | Shared recipe commands |

A costume's `branding` property selects a bundle from `v2/branding/`. Tailor replaces `/etc/penguins-eggs.d/branding` with that bundle's contents. Wearing a costume without this property removes the previous active branding. See the [branding guide](v2/DOCS/branding.md) before supplying a custom `branding.desc`.

## Related projects

- [Penguins' Tailor](https://github.com/pieroproietti/penguins-tailor) — the Go CLI that applies recipes.
- [Penguins' Eggs](https://github.com/pieroproietti/penguins-eggs) — the C/Go remastering engine and installers.
- [Quirinux atelier](https://github.com/charliemartinez/penguins-wardrobe) — Charlie Martínez's costume collection.
- [Project website](https://penguins-eggs.net).

## Credits and licenses

Created and maintained by Piero Proietti. See [LICENSE](LICENSE) for the repository license. Bundled artwork, themes, and scripts may carry their own copyright and license notices; retain those when reusing them.
