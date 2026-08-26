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
- **Export**: Transfer native packages (`wardrobe export pkg`) or execution logs and reports (`wardrobe export log`) to remote storage via SSH.
- **Build**: Integrated packaging tool to compile binaries and produce native distribution packages (`wardrobe tools build`).
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

## 👔 Command Reference

### Basic Commands

- **`wardrobe get`**
  Clones or updates the wardrobe costumes repository into the hidden home directory `~/.wardrobe`.
  ```bash
  wardrobe get
  ```

- **`wardrobe list`**
  Lists all available costumes found in the wardrobe repository along with a brief description.
  ```bash
  wardrobe list
  ```

- **`wardrobe show <costume>`**
  Shows detailed metadata for a specific costume (e.g. description, supported distributions, packages, accessories, and commands).
  ```bash
  wardrobe show colibri
  ```

- **`wardrobe wear <costume>`**
  Applies the specified costume to the system. Requires root privileges (`sudo`).
  ```bash
  sudo wardrobe wear colibri
  ```
  **Flags:**
  - `-i, --interactive`: Run in full interactive mode with live terminal output and keyboard input for prompts.
  - `--no-acc`: Skip installing accessory packages.
  - `--no-firm`: Skip installing firmware accessories.

---

### Export Commands

The **`export`** command suite automates the transfer of generated artifacts and logs to configured remote destinations:

#### 1. Export Native Packages (`wardrobe export pkg`)
Transfers compiled native packages (`.deb`, `.rpm`, `.pkg.tar.zst`, `.apk`) corresponding to the current distribution family to the remote storage server (`root@192.168.1.2:/eggs/`). It establishes an SSH multiplexed connection for efficient multi-file transfer.

```bash
# Export the built package
wardrobe export pkg

# Clean old versions on the remote server before exporting
wardrobe export pkg --clean
```

**Flags:**
- `--clean`: Removes previous versions of the package matching the distribution pattern on the remote server before uploading the new one.

#### 2. Export Logs and Reports (`wardrobe export log`)
Collects and uploads the main wardrobe log file (`/var/log/wardrobe.log`) and the latest detailed wear report (`/var/log/wardrobe/wardrobe-report-*.txt`) to the target server in a single SSH session without requiring manual file copying.

```bash
# Export logs to default remote destination
wardrobe export log

# Export logs with custom SSH user, IP, and destination directory
wardrobe export log -u artisan -i 192.168.1.50 -d /home/artisan/logs
```

**Flags:**
- `-u, --user <username>`: Remote SSH username (default: `artisan`).
- `-i, --ip <address>`: Remote IP address or hostname (default: `192.168.1.2`).
- `-d, --dir <path>`: Destination directory on the remote machine (default: `/home/artisan`).

---

### Packaging & Auxiliary Tools

- **`wardrobe tools build`**
  Compiles binaries and generates distribution-specific packages (`.deb` for Debian/Ubuntu, `PKGBUILD`/`.pkg.tar.zst` for Arch Linux, `.rpm` for Fedora/OpenSUSE, `.apk` for Alpine). Must be run as a regular user (not root).
  ```bash
  wardrobe tools build
  ```

---

## 📖 Documentation

A comprehensive guide is available in the documentation folder:
- [Wardrobe Users' Guide (Italian)](./v2/DOCS/wardrobe-users-guide.md)

---

## 📜 License

MIT License. Copyright (c) 2026 Piero Proietti.

