<div align="center">
  <img src="penguins-wardrobe.png" alt="Penguins' Wardrobe" width="180">

  # The Wardrobe · v2

  **The working collection of costumes, accessories, branding, and scripts.**

  [User guide](DOCS/wardrobe-users-guide.md) · [Tailor](https://github.com/pieroproietti/penguins-tailor) · [Project home](https://penguins-eggs.net)
</div>

---

## Find your way around

```text
v2/
├── costumes/       complete desktop and system recipes
├── accessories/    reusable, optional components
├── branding/       visual identity and distribution customizations
├── scripts/        shared configuration helpers
└── DOCS/           guides and reference material
```

| Collection | Examples | Purpose |
| --- | --- | --- |
| [`costumes/`](costumes/) | `colibri`, `duck`, `eagle`, `seagull` | Ready-to-wear system profiles |
| [`accessories/`](accessories/) | `firmwares`, `flatpak`, `kvm`, `office` | Modular features used alone or by a costume |
| [`branding/`](branding/) | `nexa`, `spiral`, `ufficiozero` | Live media, bootloader, and installer identity |
| [`scripts/`](scripts/) | display manager and hostname helpers | Reusable finishing steps |

Recipes may also include a `packages.preseed` file for unattended Debconf answers
and a `sysroot/` or `dirs/` tree for files that must be overlaid onto the target
system.

A costume can select an optional visual identity with the top-level `branding`
property. Its value is the name of a directory in [`branding/`](branding/):

```yaml
name: quirinux
branding: quirinux
```

When the property is absent, no branding is applied.

## Try on a costume

With [Penguins' Tailor](https://github.com/pieroproietti/penguins-tailor) installed:

```bash
tailor get                       # clone or refresh the wardrobe
tailor list                      # browse available costumes
tailor show colibri              # inspect one recipe
tailor wear colibri --dry-run    # preview without making changes
sudo tailor wear colibri         # apply it
```

Accessories can be fitted directly too:

```bash
sudo tailor wear accessories/multimedia
```

For custom repositories, recipe anatomy, command flags, exports, and debugging,
continue with the **[Wardrobe users' guide](DOCS/wardrobe-users-guide.md)**.

## Related ateliers

- [Penguins' Tailor](https://github.com/pieroproietti/penguins-tailor) — the CLI that applies these recipes
- [Atelier Quirinux](https://github.com/charliemartinez/atelier-quirinux) — an advanced custom wardrobe
- [Penguins' Eggs](https://penguins-eggs.net) — the wider remastering ecosystem

## License

Copyright © 2026 Piero Proietti. Dual-licensed under the
[MIT](LICENSE) or [GNU GPL v2](LICENSE) license.
