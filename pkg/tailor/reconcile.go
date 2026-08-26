package tailor

import (
	"bufio"
	"github.com/pieroproietti/penguins-wardrobe/pkg/utils"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// neverPurgeBase is a safety net of packages that must never be removed
// during package purges. This is defense-in-depth against accidental bricking.
var neverPurgeBase = []string{
	// package management
	"dpkg",
	"apt",
	"apt-utils",
	"apt-transport-https",
	"ca-certificates",
	// base system
	"base-files",
	"base-passwd",
	"init",
	"sysvinit",
	"sysvinit-core",
	"systemd",
	"systemd-sysv",
	"libc6",
	"coreutils",
	"bash",
	"dash",
	"util-linux",
	"e2fsprogs",
	"mount",
	// remastering tools: the wardrobe must never uninstall itself
	"penguins-wardrobe",
	"wardrobe",
	// Init systems: never install/purge; each base system keeps its own.
	"systemd", "systemd-sysv", "systemd-timesyncd", "libsystemd0", "libpam-systemd", "libnss-systemd",
	"sysvinit", "sysvinit-core", "sysvinit-utils", "sysv-rc", "sysv-rc-conf", "insserv", "startpar",
	"elogind", "libelogind0", "libelogind-compat", "libpam-elogind",
	"eudev", "libeudev1", "libeudev-dev", "udev", "libudev1", "libudev-dev",
	"penguins-eggs",
	"coa",
	// boot
	"grub-pc",
	"grub-common",
	"grub2-common",
	"grub-efi-amd64",
	"linux-base",
	"initramfs-tools",
	// networking / remote access
	"openssh-server",
	"openssh-client",
	"network-manager",
	"network-manager-gnome",
}

// currentKernelPackage returns the package that owns the currently
// running kernel (e.g. "linux-image-6.1.0-10-amd64"), so it can be
// protected from removal.
func currentKernelPackage() string {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return ""
	}
	release := strings.TrimSpace(string(out))
	if release == "" {
		return ""
	}
	pkg := "linux-image-" + release
	if !isPackageInstalled(pkg) {
		return ""
	}
	return pkg
}

// loadPackageList reads package names from a file. It accepts
// multiple formats, auto-detected line by line:
// - plain: one package name per line ("thunderbird")
// - YAML-style: "    - thunderbird" or "- thunderbird"
// - dpkg -l / dpkg-query -W style: "ii thunderbird 1:128.0-1 amd64 ..."
// Blank lines, headers and lines starting with '#' are ignored.
func loadPackageList(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	seen := make(map[string]struct{})
	var result []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") || strings.HasSuffix(line, ":") || line == "---" {
			continue
		}

		var pkg string

		// YAML-style: "    - package" or "- package"
		if strings.HasPrefix(line, "- ") {
			pkg = strings.TrimPrefix(line, "- ")
		} else {
			fields := strings.Fields(line)
			if len(fields) >= 2 && len(fields[0]) <= 3 && isDpkgStatusCode(fields[0]) {
				// dpkg -l style: "ii package-name version arch description..."
				pkg = fields[1]
			} else {
				pkg = fields[0]
			}
		}

		pkg = normalizePkgName(pkg)
		if pkg == "" {
			continue
		}
		if _, ok := seen[pkg]; ok {
			continue
		}
		seen[pkg] = struct{}{}
		result = append(result, pkg)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func isDpkgStatusCode(s string) bool {
	switch s {
	case "ii", "rc", "un", "iF", "iU", "hi", "pn":
		return true
	}
	return false
}

// currentlyInstalledPackages reads the dpkg status database DIRECTLY from
// /var/lib/dpkg/status instead of shelling out to dpkg-query.
func currentlyInstalledPackages() (map[string]struct{}, error) {
	f, err := os.Open("/var/lib/dpkg/status")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	installed := make(map[string]struct{})
	var curPkg string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "Package: "):
			curPkg = normalizePkgName(strings.TrimPrefix(line, "Package: "))
		case strings.HasPrefix(line, "Status: "):
			if strings.TrimPrefix(line, "Status: ") == "install ok installed" && curPkg != "" {
				installed[curPkg] = struct{}{}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return installed, nil
}

// purgeExplicit purges exactly the given packages in a SINGLE apt
// transaction, then sweeps orphaned dependencies. Packages that are not
// installed, or that belong to the safety net (neverPurgeBase / running kernel),
// are silently skipped.
func purgeExplicit(toRemove []string) {
	installedSet, err := currentlyInstalledPackages()
	if err != nil {
		utils.LogNormal("WARNING: could not read installed packages; skipping explicit purge.")
		return
	}

	protect := make(map[string]struct{})
	for _, p := range neverPurgeBase {
		protect[normalizePkgName(p)] = struct{}{}
	}
	if k := currentKernelPackage(); k != "" {
		protect[normalizePkgName(k)] = struct{}{}
	}

	seen := make(map[string]struct{})
	var list []string
	for _, p := range toRemove {
		p = normalizePkgName(p)
		if p == "" {
			continue
		}
		if _, ok := seen[p]; ok {
			continue
		}
		seen[p] = struct{}{}
		if _, ok := protect[p]; ok {
			continue
		}
		if _, ok := installedSet[p]; !ok {
			continue
		}
		list = append(list, p)
	}

	if len(list) == 0 {
		utils.LogNormal("Explicit purge: nothing to remove.")
		return
	}

	logToFile(fmt.Sprintf("Explicit purge: removing %d packages...", len(list)))
	cmd := fmt.Sprintf("DEBIAN_FRONTEND=readline apt-get purge -o Dpkg::Options::='--force-confold' -y %s", strings.Join(list, " "))
	if err := utils.ExecTee(cmd, wardrobeLogFile); err != nil {
		logToFile("WARNING: bulk explicit purge reported an error; healing and retrying once...")
		_ = utils.ExecTee("dpkg --configure -a", wardrobeLogFile)
		_ = utils.ExecTee("DEBIAN_FRONTEND=readline apt-get install -f -y", wardrobeLogFile)
		_ = utils.ExecTee(cmd, wardrobeLogFile)
	}

	logToFile("Sweeping orphaned dependencies of removed packages...")
	_ = utils.ExecTee("DEBIAN_FRONTEND=readline apt-get autoremove -o Dpkg::Options::='--force-confold' --purge -y", wardrobeLogFile)
}

func findPackageListFile(costumeDir, filename string) string {
	if filename == "" {
		return ""
	}
	path := filepath.Join(costumeDir, filename)
	if _, err := os.Stat(path); err != nil {
		return ""
	}
	return path
}

