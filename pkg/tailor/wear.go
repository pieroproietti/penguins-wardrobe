package tailor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pieroproietti/penguins-wardrobe/pkg/utils"
)

func Wear(costumeName string, noAcc bool, noFirm bool) error {
	if os.Geteuid() != 0 {
		utils.LogError("'wardrobe wear' needs to install packages and write to system paths; run it as root (e.g. 'su' first, or 'sudo wardrobe wear %s' if sudo is configured for your user).", costumeName)
		return fmt.Errorf("must be run as root")
	}

	utils.LogNormal("Starting costume application for: %s", costumeName)
	v2Dir, err := getWardrobeV2Dir()
	if err != nil {
		utils.LogError("Wardrobe root error: %v", err)
		return err
	}
	costumeDir := filepath.Join(v2Dir, "costumes", costumeName)
	if _, err := os.Stat(costumeDir); os.IsNotExist(err) {
		if strings.HasPrefix(costumeName, "accessories/") || strings.HasPrefix(costumeName, "costumes/") {
			costumeDir = filepath.Join(v2Dir, costumeName)
		} else {
			accDir := filepath.Join(v2Dir, "accessories", costumeName)
			if _, errAcc := os.Stat(accDir); errAcc == nil {
				costumeDir = accDir
			}
		}
	}
	if _, err := os.Stat(costumeDir); os.IsNotExist(err) {
		return fmt.Errorf("costume '%s' not found in %s", costumeName, costumeDir)
	}

	yamlFile := findYaml(costumeDir)
	suit, err := loadSuit(yamlFile)
	if err != nil {
		return err
	}

	// Enforce distribution compatibility BEFORE anything else, including
	// kernel header installation, so an incompatible system aborts cleanly
	// without touching the machine or producing unrelated apt output.
	if err := checkCostumeCompatibility(costumeDir, suit); err != nil {
		return err
	}

	// DKMS safety: make sure the headers for the RUNNING kernel are in
	// place before any package is unpacked, so DKMS postinsts that build
	// for the current kernel do not abort mid-transaction. Only reached on
	// compatible systems.
	ensureKernelHeaders()

	utils.LogNormal("--- Applying Costume: %s ---", suit.Name)

	installedPackages, failedPackages, err := applySuit(costumeDir, suit)
	if err != nil {
		return err
	}

	if !noAcc && len(suit.Accessories) > 0 {
		utils.LogNormal("--- Processing %d accessories ---", len(suit.Accessories))
		for _, accName := range suit.Accessories {
			if noFirm && (accName == "firmwares" || strings.Contains(accName, "firmware")) {
				utils.LogNormal("Skipping firmware accessory (%s) due to --no-firmwares flag", accName)
				continue
			}

			var accDir string
			if strings.HasPrefix(accName, "./") || strings.HasPrefix(accName, "../") {
				accDir = filepath.Join(costumeDir, accName)
			} else if strings.HasPrefix(accName, "accessories/") {
				accDir = filepath.Join(v2Dir, accName)
			} else {
				accDir = filepath.Join(v2Dir, "accessories", accName)
			}

			if accYaml := findYaml(accDir); accYaml != "" {
				if accSuit, err := loadSuit(accYaml); err == nil {
					utils.LogNormal("Accessory: %s (%s)", accName, filepath.Base(accYaml))
					accInstalled, accFailed, _ := applySuit(accDir, accSuit)
					installedPackages = append(installedPackages, accInstalled...)
					failedPackages = append(failedPackages, accFailed...)
				} else {
					utils.LogNormal(utils.ColorYellow+"WARNING: could not load accessory '%s': %v"+utils.ColorReset, accName, err)
				}
			} else {
				utils.LogNormal(utils.ColorYellow+"WARNING: accessory '%s' not found in %s"+utils.ColorReset, accName, accDir)
			}
		}
	}

	var purgedPackages []string
	var failedPurges []string

	installedBefore, _ := currentlyInstalledPackages()

	// Install everything in the manifest that's missing. The manifest is
	// resolved per-distribution: a "*_<codename>-packages.list" file wins
	// over the generic packages_manifest declared in index.yaml.
	if manifestPath := resolveDistroManifest(costumeDir, suit.PackagesManifest); manifestPath != "" {
		utils.LogNormal("--- Declarative manifest (authoritative install list): %s ---", manifestPath)
		if targetManifest, err := loadPackageManifest(manifestPath); err == nil {
			utils.LogNormal("[%s] Installing %d manifest packages...", suit.Name, len(targetManifest))
			manifestFailed := installWithRetries(targetManifest, 3)
			failedPackages = append(failedPackages, manifestFailed...)
			installedPackages = append(installedPackages, diffStr(targetManifest, manifestFailed)...)
		} else {
			utils.LogNormal(utils.ColorYellow+"WARNING: could not read packages_manifest %s: %v"+utils.ColorReset, manifestPath, err)
		}
	}

	// Load packages from external install file if specified
	if installPath := findManifestPath(costumeDir, suit.PackagesInstallFile); installPath != "" {
		utils.LogNormal("--- Loading packages from external install file: %s ---", installPath)
		if filePackages, err := loadPackageManifest(installPath); err == nil {
			utils.LogNormal("[%s] Installing %d packages from external file...", suit.Name, len(filePackages))
			fileFailed := installWithRetries(filePackages, 3)
			failedPackages = append(failedPackages, fileFailed...)
			installedPackages = append(installedPackages, diffStr(filePackages, fileFailed)...)
		} else {
			utils.LogNormal(utils.ColorYellow+"WARNING: could not read packages_install_file %s: %v"+utils.ColorReset, installPath, err)
		}
	}

	// Deterministic removal: purge exactly the vendor's remove list
	var removeList []string
	removeList = append(removeList, suit.PackagesRemove...)
	if removePath := findManifestPath(costumeDir, suit.PackagesRemoveFile); removePath != "" {
		utils.LogNormal("--- Declarative remove list: %s ---", removePath)
		if fileRemove, err := loadPackageManifest(removePath); err == nil {
			removeList = append(removeList, fileRemove...)
		} else {
			utils.LogNormal(utils.ColorYellow+"WARNING: could not read packages_remove_file %s: %v"+utils.ColorReset, removePath, err)
		}
	}
	if len(removeList) > 0 {
		purgeExplicit(removeList)
	}

	// DKMS healing: the manifest usually installs a NEWER kernel, and DKMS
	// postinsts run before that kernel's headers are on disk, aborting and
	// leaving dpkg half-configured (which then poisons every later apt-get
	// call, e.g. quirinux-firmware failing on dependencies). Repair the
	// state and retry before writing the final report.
	failedPackages = healAndRetryFailed(failedPackages)

	installedAfter, _ := currentlyInstalledPackages()
	if len(installedBefore) > 0 && len(installedAfter) > 0 {
		for p := range installedBefore {
			if _, ok := installedAfter[p]; !ok {
				purgedPackages = append(purgedPackages, p)
			}
		}
	}

	copySkelToUser()
	reportPath, reportErr := writeWearReport(wearReport{
		CostumeName:   suit.Name,
		Installed:     installedPackages,
		Purged:        purgedPackages,
		FailedInstall: failedPackages,
		FailedPurge:   failedPurges,
	})

	clearScreen()
	utils.LogNormal("Costume '%s' applied. Installed: %d | Removed: %d | Could not be installed: %d | Could not be removed: %d",
		suit.Name, len(installedPackages), len(purgedPackages), len(failedPackages), len(failedPurges))

	if reportErr != nil {
		utils.LogNormal(utils.ColorYellow+"WARNING: could not write detailed report: %v"+utils.ColorReset, reportErr)
	} else {
		utils.LogNormal("Detailed report: %s", reportPath)
	}
	if suit.Reboot {
		utils.LogNormal(utils.ColorYellow + "This costume recommends a reboot to finish applying all changes." + utils.ColorReset)
	}
	printKernelCleanupReminder()
	if suit.DisplayManagerNotice {
		printDisplayManagerNotice()
	}
	return nil
}

// checkCostumeCompatibility enforces the distribution compatibility declared
// by the costume. It compares the "distributions" key from index.yaml
// (already parsed into suit.Distributions) against the running distribution
// codename. If the costume declares supported distributions and the running
// one is not among them, the wear is aborted BEFORE any package is installed.
// The wardrobe-check script, when present, is run only to print a localized
// message; the enforcement does not depend on it.
func checkCostumeCompatibility(costumeDir string, suit *Suit) error {
	if len(suit.Distributions) == 0 {
		// Costume does not declare supported distributions; assume compatible
		// with any distribution (backward compatibility with older wardrobes)
		return nil
	}

	current := currentDistroCodename()
	if current == "" {
		utils.LogNormal("WARNING: could not detect the running distribution; skipping compatibility check.")
		return nil
	}

	for _, d := range suit.Distributions {
		if strings.EqualFold(strings.TrimSpace(d), current) {
			return nil
		}
	}

	// Incompatible. Print ONE detailed explanation: the localized message
	// from the wardrobe-check script when present, or a generic one otherwise.
	script := filepath.Join(costumeDir, "wardrobe-check")
	if _, err := os.Stat(script); err == nil {
		cmd := exec.Command("bash", script)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		_ = cmd.Run()
	} else {
		fmt.Fprintf(os.Stderr,
			"This costume (%s) is only compatible with: %s. Detected distribution: %s.\n",
			suit.Name, strings.Join(suit.Distributions, ", "), current)
	}

	// Short error: the detailed explanation was already printed above, so
	// keep the returned error terse to avoid repeating the long text.
	return fmt.Errorf("aborted: distribution %q not supported", current)
}

// currentDistroCodename reads /etc/os-release and returns the VERSION_CODENAME
// value (e.g. "daedalus", "bookworm", "excalibur"), or an empty string if it
// cannot be determined.
func currentDistroCodename() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "VERSION_CODENAME=") {
			return strings.Trim(strings.TrimPrefix(line, "VERSION_CODENAME="), `"`)
		}
	}
	return ""
}

// ensureKernelHeaders installs the kernel headers matching the currently
// running kernel (plus the architecture meta-package) before any DKMS
// package is unpacked. A DKMS postinst aborts the whole transaction when
// the headers for a target kernel are missing, leaving dpkg in a
// half-configured state.
func ensureKernelHeaders() {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		utils.LogNormal("WARNING: could not determine running kernel version: %v", err)
		return
	}
	release := strings.TrimSpace(string(out))
	if release == "" {
		return
	}
	archOut, _ := exec.Command("dpkg", "--print-architecture").Output()
	arch := strings.TrimSpace(string(archOut))
	if arch == "" {
		arch = "amd64"
	}
	pkgs := fmt.Sprintf("linux-headers-%s linux-headers-%s", release, arch)
	utils.LogNormal("Ensuring kernel headers are present before DKMS installs: %s", pkgs)
	utils.Exec("DEBIAN_FRONTEND=noninteractive apt-get install -o Dpkg::Use-Pty=0 -y " + pkgs)
}

// healAndRetryFailed repairs the half-configured dpkg state that DKMS
// packages leave behind when kernel headers were not yet in place, then
// retries every failed package that actually exists in the apt cache.
// Packages that are simply absent from the repositories stay in the
// returned list so they keep being reported as failed.
func healAndRetryFailed(failed []string) []string {
	if len(failed) == 0 {
		return nil
	}

	utils.LogNormal("Healing dpkg state before retrying failed packages...")
	utils.Exec("dpkg --configure -a")
	utils.Exec("DEBIAN_FRONTEND=noninteractive apt-get install -f -o Dpkg::Use-Pty=0 -y")

	available := getAvailablePackages()
	var retry []string
	for _, p := range failed {
		if available == nil {
			retry = append(retry, p)
			continue
		}
		if _, ok := available[normalizePkgName(p)]; ok {
			retry = append(retry, p)
		}
	}
	if len(retry) == 0 {
		return failed
	}

	utils.LogNormal("Retrying %d packages now that kernel headers are in place...", len(retry))
	installWithRetries(retry, 1)

	var still []string
	for _, p := range failed {
		if !isPackageInstalled(p) {
			still = append(still, p)
		}
	}
	return still
}

// applySuit applies a costume/accessory and returns the list of packages
// that could not be installed (across packages, packages_no_install_recommends
// and packages_interactive), so the caller can report them to the user.
func applySuit(dir string, suit *Suit) ([]string, []string, error) {
	var installedPackages []string
	var failedPackages []string

	if suit.Sequence != nil && suit.Sequence.Repositories != nil {
		setupRepositories(suit.Sequence.Repositories, suit.Name)
		utils.LogNormal("[%s] Refreshing package index after repository changes...", suit.Name)
		if err := utils.Exec("apt-get update"); err != nil {
			utils.LogNormal("[%s] WARNING: apt-get update failed, newly added repositories may be unusable: %v", suit.Name, err)
		}
	}

	if len(suit.Packages) > 0 {
		utils.LogNormal("[%s] Attempting package installation: %v", suit.Name, suit.Packages)
		failed := installWithRetries(suit.Packages, 3)
		failedPackages = append(failedPackages, failed...)
		installedPackages = append(installedPackages, diffStr(suit.Packages, failed)...)
	} else {
		utils.LogNormal("[%s] No packages to install.", suit.Name)
	}

	if len(suit.PackagesNoRecommends) > 0 {
		utils.LogNormal("[%s] Installing packages without recommends: %v", suit.Name, suit.PackagesNoRecommends)
		failed := installNoRecommends(suit.PackagesNoRecommends)
		failedPackages = append(failedPackages, failed...)
		installedPackages = append(installedPackages, diffStr(suit.PackagesNoRecommends, failed)...)
	}

	if len(suit.PackagesInteractive) > 0 {
		utils.LogNormal("[%s] Installing interactive packages (license prompts may appear): %v", suit.Name, suit.PackagesInteractive)
		failed := installInteractive(suit.PackagesInteractive)
		failedPackages = append(failedPackages, failed...)
		installedPackages = append(installedPackages, diffStr(suit.PackagesInteractive, failed)...)
	}

	if len(suit.PackagesRemove) > 0 {
		utils.LogNormal("[%s] Removing packages not needed by this vendor: %v", suit.Name, suit.PackagesRemove)
		removePackages(suit.PackagesRemove)
	}

	sysrootPath := filepath.Join(dir, "sysroot")
	if _, err := os.Stat(sysrootPath); os.IsNotExist(err) {
		sysrootPath = filepath.Join(dir, "dirs")
	}
	if _, err := os.Stat(sysrootPath); err == nil {
		utils.LogNormal("[%s] Overlay folder found: %s", suit.Name, sysrootPath)
		utils.LogNormal("[%s] Running rsync to root /...", suit.Name)
		cmd := fmt.Sprintf("rsync -aAXv %s/ /", sysrootPath)
		if err := utils.Exec(cmd); err != nil {
			utils.LogNormal("[%s] Error during overlay: %v", suit.Name, err)
		} else {
			utils.LogNormal("[%s] Overlay completed successfully.", suit.Name)
		}
	} else {
		utils.LogNormal("[%s] No sysroot/dirs folder found, skipping overlay.", suit.Name)
	}

	if len(suit.Cmds) > 0 {
		utils.LogNormal("[%s] Running %d post-installation commands...", suit.Name, len(suit.Cmds))
		for _, command := range suit.Cmds {
			utils.LogNormal("[%s] Executing: %s", suit.Name, command)
			fields := strings.Fields(command)
			if len(fields) > 0 {
				relScript := filepath.Join(dir, fields[0])
				if stat, err := os.Stat(relScript); err == nil && !stat.IsDir() {
					rest := strings.TrimSpace(command[len(fields[0]):])
					var fullCmd string
					if rest != "" {
						fullCmd = fmt.Sprintf("%s %s", relScript, rest)
					} else {
						fullCmd = fmt.Sprintf("%s %s", relScript, suit.Name)
					}
					utils.Exec(fullCmd)
					continue
				}
			}
			utils.Exec(command)
		}
	}

	return installedPackages, failedPackages, nil
}

func copySkelToUser() {
	targetUser := os.Getenv("SUDO_USER")
	var userHome string
	if targetUser != "" {
		userHome = filepath.Join("/home", targetUser)
	} else if u := firstHumanUser(); u != nil {
		targetUser = u.Username
		userHome = u.HomeDir
	}

	if targetUser == "" || targetUser == "root" {
		utils.LogNormal("WARNING: unable to determine a non-root target user, skipping /etc/skel sync to avoid leaving files owned by root")
		return
	}

	utils.LogNormal("Syncing /etc/skel -> %s", userHome)
	cmd := fmt.Sprintf("rsync -a --no-o --no-g --chown=%s:%s /etc/skel/ %s/", targetUser, targetUser, userHome)
	utils.Exec(cmd)
}
