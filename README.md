# nota per Charlie

ho creato questa nuova nuova repo penguins-wardrobe.

Attenzione: adesso la repo contiene sia il programma che il wardarobe sotto v2.

Ho estratto da penguins-eggs il wardrobe ed il necessario per farlo compilare. poi lo ho aggiornato con la tua versione https://github.com/charliemartinez/penguins-eggs al 22 agosto 2026.

Si compila con ./m e niente, dovrebbe essere semplicemente allineato al tuo.


# penguins-wardrobe/v2

**penguins-wardrobe** is a standalone, lightweight tool written in Go to manage and apply system configurations, desktop environments, and themes ("costumes") to Linux distributions.

---

## 🚀 Features

- **Get**: Download or update the costume repository (`wardrobe get`).
- **List**: Enumerate available costumes and their descriptions (`wardrobe list`).
- **Show**: Inspect detailed information and packages required by a costume (`wardrobe show <costume>`).
- **Wear**: Seamlessly apply a costume to the system (`sudo wardrobe wear <costume>`), configuring repositories, packages, sysroot overlays, and user skel settings.
- **Distro-Aware**: Automatically identifies target distributions (Debian, Ubuntu, Arch, Alpine, Fedora, OpenSUSE, etc.) and generates assistance prompts if non-Debian package managers are present.

---

## 📦 Installation

```bash
git clone https://github.com/pieroproietti/penguins-wardrobe.git
cd penguins-wardrobe
make
sudo make install
```

---

## 👔 Usage

```bash
# Download / clone wardrobe costumes
wardrobe get

# List available costumes
wardrobe list

# Show details of a costume
wardrobe show colibri

# Wear a costume (requires root)
sudo wardrobe wear colibri

# Flags:
# --no-acc   Do not install accessories
# --no-firm  Do not install firmware
```

---

## 📜 License

MIT License. Copyright (c) 2026 Piero Proietti.
