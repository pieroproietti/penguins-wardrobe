<div align="center">
  <img src="v2/penguins-wardrobe.png" alt="Penguins' Wardrobe" width="220">

  # Penguins' Wardrobe

  **A curated collection of Linux costumes, accessories, and finishing touches.**

  Dress a minimal installation as a polished desktop or server with declarative,
  reusable recipes for [Penguins' Tailor](https://github.com/pieroproietti/penguins-tailor).

  [Explore the wardrobe](v2/) · [Read the guide](v2/DOCS/wardrobe-users-guide.md) · [Visit Penguins' Eggs](https://penguins-eggs.net)
</div>

---

## The atelier

Penguins' Wardrobe keeps system customization separate from the tool that applies
it. **Wardrobe** supplies the YAML recipes and filesystem overlays; **Tailor**
selects them, resolves their components, and dresses the target system.

| Collection | What you will find |
| --- | --- |
| **Costumes** | Complete desktop and system profiles, from lightweight workstations to headless servers |
| **Accessories** | Optional capabilities such as firmware, Flatpak, multimedia, office tools, KVM, and Waydroid |
| **Vendors** | Distribution branding for live media, bootloaders, and Calamares |
| **Scripts** | Reusable finishing steps for display managers, hostnames, and desktop integration |

All current definitions live in [`v2/`](v2/).

## A quick fitting

Install [Penguins' Tailor](https://github.com/pieroproietti/penguins-tailor), then:

```bash
# Fetch or update the official wardrobe
tailor get

# Browse and inspect the available costumes
tailor list
tailor show colibri

# Preview the fitting without changing the system
tailor wear colibri --dry-run

# Wear the selected costume
sudo tailor wear colibri
```

> [!TIP]
> Start with `--dry-run` to review the complete fitting before applying it.

## Make it yours

Wardrobes are ordinary Git repositories, so the same structure can host a private
atelier, a company image, or a community distribution. Begin with the
[Wardrobe users' guide](v2/DOCS/wardrobe-users-guide.md) for the recipe format,
overlays, preseeding, and the complete Tailor workflow.

## License

Copyright © 2026 Piero Proietti. Dual-licensed under the
[MIT](LICENSE) or [GNU GPL v2](LICENSE) license.
