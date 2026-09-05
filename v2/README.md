# Wardrobe v2

This directory is the collection consumed by `tailor`. For installation and commands, read the shared [Wardrobe and Tailor user guide](DOCS/wardrobe-users-guide.md). For live media and installer artwork, read the [branding guide](DOCS/branding.md).

```text
v2/
├── costumes/<name>/       complete recipes
├── accessories/<name>/    reusable recipes
├── branding/<name>/       visual identity bundles
├── scripts/               shared commands
└── DOCS/                  Wardrobe and Tailor documentation
```

## Recipes

A costume or accessory typically contains:

```text
index.yaml                recipe and metadata
packages.yaml             optional additional packages
packages.preseed          optional Debian Debconf answers
sysroot/                  optional files copied onto /
scripts/                  optional recipe scripts
```

Tailor prefers `index.yaml`, then `index.yml`. When neither exists, it searches distribution-specific filenames. It selects one recipe; it does not merge `index.yaml` with `debian.yaml` or `arch.yaml`. See the [recipe reference](DOCS/wardrobe-users-guide.md#scrivere-una-ricetta).

```bash
tailor get
tailor list
tailor show colibri
tailor show accessories/multimedia
sudo tailor wear colibri
sudo tailor wear accessories/multimedia
```

The current `wear` package backend supports Debian and derivatives through APT. Check the recipe's `distributions` and package names for the target system.

## Branding

Select a bundle in a costume's `index.yaml`:

```yaml
name: my-desktop
branding: quirinux
```

Tailor copies the contents of `branding/quirinux/` directly into `/etc/penguins-eggs.d/branding/`, replacing the previous selection. A costume without `branding` removes the active bundle; directly wearing an accessory preserves it.

For Calamares, omit `branding.desc` when the generated distro identity is suitable. If you supply that file, it must be complete and use `componentName: eggs`. Empty files and comment-only stubs overwrite the generated descriptor and break the configuration. The [branding guide](DOCS/branding.md) explains the full layout and precedence.

## Working on an atelier

Tailor normally reads the real user's `~/.wardrobe/v2`, or `~/.wardrobe` for a collection without a `v2` wrapper. Local `./v2` is only a fallback when that installed wardrobe is absent. Running Tailor from a development checkout does not automatically select that checkout.

Use a separate Git branch for recipe work. Review `tailor show`, the selected YAML, and the files under `sysroot/` before applying changes. The [user guide](DOCS/wardrobe-users-guide.md) covers previews and log locations.
