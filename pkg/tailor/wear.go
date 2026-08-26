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
		utils.LogError("'wardrobe wear' needs to install packages and write to system paths; run it as root (e.g. 'sudo wardrobe wear %s').", costumeName)
		return fmt.Errorf("must be run as root")
	}

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

	// Enforce distribution compatibility BEFORE anything else
	if err := checkCostumeCompatibility(costumeDir, suit); err != nil {
		utils.LogError("%s", incompatibleDistroMessage(suit.Name, suit.Distributions, currentDistroName()))
		return err
	}

	isDirectAccessory := strings.HasPrefix(costumeName, "accessories/") || (suit.Name != "" && !strings.Contains(costumeDir, "/costumes/"))

	icon := "👗"
	title := fmt.Sprintf("COSTUME: %s", suit.Name)
	if suit.Release != "" {
		title = fmt.Sprintf("COSTUME: %s (v%s)", suit.Name, suit.Release)
	}
	if isDirectAccessory {
		icon = "👝"
		title = fmt.Sprintf("ACCESSORY: %s", suit.Name)
	}

	ss := utils.StartSplitScreen(icon, title, suit.Description)
	if ss != nil {
		defer ss.Close()
	} else {
		utils.PrintBanner(icon, title, suit.Description)
	}

	// DKMS safety: ensure headers for running kernel are present
	if ss != nil {
		ss.SetAction("Checking kernel headers for DKMS...")
		if err := ensureKernelHeaders(); err != nil {
			ss.AddStep(fmt.Sprintf("%s[WARN]%s Kernel headers verification completed with warnings", utils.ColorYellow, utils.ColorReset))
		} else {
			ss.AddStep(fmt.Sprintf("%s[OK]%s Kernel headers verified", utils.ColorGreen, utils.ColorReset))
		}
	} else {
		spHeaders := utils.NewSpinner("Checking kernel headers for DKMS...")
		spHeaders.Start()
		if err := ensureKernelHeaders(); err != nil {
			spHeaders.Warn("Kernel headers verification completed with warnings")
		} else {
			spHeaders.Success("Kernel headers verified")
		}
	}

	SetLicensePromptPackages(suit.PackagesInteractive)

	installedPackages, failedPackages, err := applySuit(costumeDir, suit)
	if err != nil {
		return err
	}

	if !noAcc && len(suit.Accessories) > 0 {
		if ss == nil {
			utils.PrintSection("👝", fmt.Sprintf("ACCESSORIES (%d items)", len(suit.Accessories)))
		}
		for idx, accName := range suit.Accessories {
			if noFirm && (accName == "firmwares" || strings.Contains(accName, "firmware")) {
				if ss != nil {
					ss.AddStep(fmt.Sprintf("%s[INFO]%s Skipping firmware accessory '%s' (--no-firm)", utils.ColorYellow, utils.ColorReset, accName))
				} else {
					fmt.Printf("\n  %s[INFO] Skipping firmware accessory '%s' due to --no-firm flag%s\n", utils.ColorYellow, accName, utils.ColorReset)
				}
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
					if ss != nil {
						ss.AddStep(fmt.Sprintf("%s--> [%d/%d] Accessory: %s%s", utils.ColorCyan, idx+1, len(suit.Accessories), accName, utils.ColorReset))
					} else {
						utils.PrintSubSection("-->", fmt.Sprintf("[%d/%d] Accessory: %s", idx+1, len(suit.Accessories), accName))
					}
					accInstalled, accFailed, _ := applySuit(accDir, accSuit)
					installedPackages = append(installedPackages, accInstalled...)
					failedPackages = append(failedPackages, accFailed...)
				} else {
					if ss != nil {
						ss.AddStep(fmt.Sprintf("%s[WARN]%s Could not load accessory '%s'", utils.ColorYellow, utils.ColorReset, accName))
					} else {
						fmt.Printf("  %s[WARN] Could not load accessory '%s': %v%s\n", utils.ColorYellow, accName, err, utils.ColorReset)
					}
				}
			} else {
				if ss != nil {
					ss.AddStep(fmt.Sprintf("%s[WARN]%s Accessory '%s' not found", utils.ColorYellow, utils.ColorReset, accName))
				} else {
					fmt.Printf("  %s[WARN] Accessory '%s' not found in %s%s\n", utils.ColorYellow, accName, accDir, utils.ColorReset)
				}
			}
		}
	}

	var purgedPackages []string
	var failedPurges []string

	installedBefore, _ := currentlyInstalledPackages()

	// Install everything in the manifest that's missing
	if manifestPath := resolveDistroManifest(costumeDir, suit.PackagesManifest); manifestPath != "" {
		if targetManifest, err := loadPackageManifest(manifestPath); err == nil {
			if ss != nil {
				ss.AddStep(fmt.Sprintf("📋 Manifest reconciliation (%s, %d packages)", filepath.Base(manifestPath), len(targetManifest)))
				ss.SetAction("Reconciling %d manifest packages...", len(targetManifest))
			} else {
				utils.PrintSection("📋", "DECLARATIVE MANIFEST RECONCILIATION")
			}
			manifestFailed := installWithRetries(targetManifest, 3)
			failedPackages = append(failedPackages, manifestFailed...)
			installedPackages = append(installedPackages, diffStr(targetManifest, manifestFailed)...)
			if ss != nil {
				if len(manifestFailed) > 0 {
					ss.AddStep(fmt.Sprintf("%s[WARN]%s Manifest reconciled (%d failed)", utils.ColorYellow, utils.ColorReset, len(manifestFailed)))
				} else {
					ss.AddStep(fmt.Sprintf("%s[OK]%s Manifest reconciled (%d packages)", utils.ColorGreen, utils.ColorReset, len(targetManifest)))
				}
			}
		} else {
			if ss != nil {
				ss.AddStep(fmt.Sprintf("%s[WARN]%s Could not read manifest %s", utils.ColorYellow, utils.ColorReset, manifestPath))
			} else {
				fmt.Printf("  %s[WARN] Could not read packages_manifest %s: %v%s\n", utils.ColorYellow, manifestPath, err, utils.ColorReset)
			}
		}
	}

	// Deterministic removal: purge exactly the vendor's remove list
	var removeList []string
	removeList = append(removeList, suit.PackagesRemove...)
	if removePath := findManifestPath(costumeDir, suit.PackagesRemoveFile); removePath != "" {
		if fileRemove, err := loadPackageManifest(removePath); err == nil {
			removeList = append(removeList, fileRemove...)
		}
	}
	if len(removeList) > 0 {
		if ss != nil {
			ss.SetAction("Purging %d packages absent from manifest...", len(removeList))
		}
		purgeExplicit(removeList)
		if ss != nil {
			ss.AddStep(fmt.Sprintf("%s[OK]%s Declarative purge completed (%d packages)", utils.ColorGreen, utils.ColorReset, len(removeList)))
		}
	}

	// DKMS healing
	if ss != nil {
		ss.SetAction("Healing DKMS state...")
	}
	failedPackages = healAndRetryFailed(failedPackages)

	installedAfter, _ := currentlyInstalledPackages()
	if len(installedBefore) > 0 && len(installedAfter) > 0 {
		for p := range installedBefore {
			if _, ok := installedAfter[p]; !ok {
				purgedPackages = append(purgedPackages, p)
			}
		}
	}

	// User environment synchronization
	targetUser := getTargetUsername()
	if targetUser != "" && targetUser != "root" {
		if ss != nil {
			ss.SetAction("Synchronizing user environment (/etc/skel -> /home/%s)...", targetUser)
		}
		copySkelToUser()
		if ss != nil {
			ss.AddStep(fmt.Sprintf("%s[OK]%s User environment synchronized (%s)", utils.ColorGreen, targetUser, utils.ColorReset))
		}
	}

	// Close split screen before printing final summary box
	if ss != nil {
		ss.Close()
		ss = nil
	}

	reportPath, reportErr := writeWearReport(wearReport{
		CostumeName:   suit.Name,
		Installed:     installedPackages,
		Purged:        purgedPackages,
		FailedInstall: failedPackages,
		FailedPurge:   failedPurges,
	})

	summaryRows := [][2]string{
		{"Costume / Oggetto", suit.Name},
		{"Pacchetti installati", fmt.Sprintf("%d", len(installedPackages))},
		{"Pacchetti rimossi", fmt.Sprintf("%d", len(purgedPackages))},
		{"Non installati", fmt.Sprintf("%d", len(failedPackages))},
	}
	if reportErr == nil {
		summaryRows = append(summaryRows, [2]string{"Report dettagliato", reportPath})
	}
	summaryRows = append(summaryRows, [2]string{"Log di sistema", wardrobeLogFile})

	utils.PrintSummaryBox("✨ VESTIZIONE COMPLETATA!", summaryRows)

	if suit.Reboot {
		fmt.Printf("\n%s%s⚠ Questo costume consiglia di riavviare il sistema al termine.%s\n", utils.ColorYellow, utils.ColorBold, utils.ColorReset)
	}
	printKernelCleanupReminder()
	if suit.DisplayManagerNotice {
		printDisplayManagerNotice()
	}
	return nil
}

func getTargetUsername() string {
	if sudoUser := os.Getenv("SUDO_USER"); sudoUser != "" {
		return sudoUser
	}
	if u := firstHumanUser(); u != nil {
		return u.Username
	}
	return ""
}

// checkCostumeCompatibility enforces the distribution compatibility declared by the costume.
func checkCostumeCompatibility(costumeDir string, suit *Suit) error {
	if len(suit.Distributions) == 0 {
		return nil
	}

	current := currentDistroCodename()
	if current == "" {
		logToFile("WARNING: could not detect the running distribution; skipping compatibility check.")
		return nil
	}

	for _, d := range suit.Distributions {
		if strings.EqualFold(strings.TrimSpace(d), current) {
			return nil
		}
	}

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

	return fmt.Errorf("aborted: distribution %q not supported", current)
}

// currentDistroCodename reads /etc/os-release and returns the VERSION_CODENAME
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

// ensureKernelHeaders installs the kernel headers matching the currently running kernel.
func ensureKernelHeaders() error {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		logToFile(fmt.Sprintf("WARNING: could not determine running kernel version: %v", err))
		return err
	}
	release := strings.TrimSpace(string(out))
	if release == "" {
		return nil
	}
	archOut, _ := exec.Command("dpkg", "--print-architecture").Output()
	arch := strings.TrimSpace(string(archOut))
	if arch == "" {
		arch = "amd64"
	}
	pkgs := fmt.Sprintf("linux-headers-%s linux-headers-%s", release, arch)
	logToFile(fmt.Sprintf("Ensuring kernel headers are present before DKMS installs: %s", pkgs))
	return utils.ExecTee("DEBIAN_FRONTEND=readline apt-get install -o Dpkg::Options::='--force-confold' -y "+pkgs, wardrobeLogFile)
}

// healAndRetryFailed repairs the half-configured dpkg state and retries failed packages.
func healAndRetryFailed(failed []string) []string {
	if len(failed) == 0 {
		return nil
	}

	logToFile("Healing dpkg state before retrying failed packages...")
	_ = utils.ExecTee("dpkg --configure -a", wardrobeLogFile)
	_ = utils.ExecTee("DEBIAN_FRONTEND=readline apt-get install -f -o Dpkg::Options::='--force-confold' -y", wardrobeLogFile)

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

	logToFile(fmt.Sprintf("Retrying %d packages now that kernel headers are in place...", len(retry)))
	installWithRetries(retry, 1)

	var still []string
	for _, p := range failed {
		if !isPackageInstalled(p) {
			still = append(still, p)
		}
	}
	return still
}

// applySuit applies a costume or accessory definition with clean spinners
func applySuit(dir string, suit *Suit) ([]string, []string, error) {
	var installedPackages []string
	var failedPackages []string
	ss := utils.GetSplitScreen()

	// Preseed (debconf-set-selections)
	if preseedFile := findPreseed(dir); preseedFile != "" {
		if ss != nil {
			ss.SetAction("Applying preseed selections (%s)...", filepath.Base(preseedFile))
			if err := applyPreseed(preseedFile, suit.Name); err != nil {
				ss.AddStep(fmt.Sprintf("%s[WARN]%s Preseed selections applied with warnings", utils.ColorYellow, utils.ColorReset))
			} else {
				ss.AddStep(fmt.Sprintf("%s[OK]%s Preseed selections applied (%s)", utils.ColorGreen, filepath.Base(preseedFile), utils.ColorReset))
			}
		} else {
			spPreseed := utils.NewSpinner(fmt.Sprintf("Applying preseed selections (%s)...", filepath.Base(preseedFile)))
			spPreseed.Start()
			if err := applyPreseed(preseedFile, suit.Name); err != nil {
				spPreseed.Warn("Preseed selections applied with warnings")
			} else {
				spPreseed.Success("Preseed selections applied (%s)", filepath.Base(preseedFile))
			}
		}
	}

	// Repositories
	if suit.Sequence != nil && suit.Sequence.Repositories != nil {
		if ss != nil {
			ss.SetAction("Configuring package repositories & updating cache...")
		}
		setupRepositories(suit.Sequence.Repositories, suit.Name)
		if ss != nil {
			ss.AddStep(fmt.Sprintf("%s[OK]%s Repositories configured & updated", utils.ColorGreen, utils.ColorReset))
		}
	}

	// Packages
	if len(suit.Packages) > 0 {
		if ss != nil {
			ss.SetAction("Installing packages (%d packages)...", len(suit.Packages))
		}
		failed := installWithRetries(suit.Packages, 3)
		failedPackages = append(failedPackages, failed...)
		installed := diffStr(suit.Packages, failed)
		installedPackages = append(installedPackages, installed...)
		if ss != nil {
			if len(failed) > 0 {
				ss.AddStep(fmt.Sprintf("%s[WARN]%s Installed %d packages (%d failed)", utils.ColorYellow, utils.ColorReset, len(installed), len(failed)))
			} else {
				ss.AddStep(fmt.Sprintf("%s[OK]%s Installed %d packages", utils.ColorGreen, utils.ColorReset, len(installed)))
			}
		}
	}

	// External install file
	if installPath := findManifestPath(dir, suit.PackagesInstallFile); installPath != "" {
		if filePackages, err := loadPackageManifest(installPath); err == nil {
			if ss != nil {
				ss.SetAction("Installing %d packages from external file (%s)...", len(filePackages), filepath.Base(installPath))
			}
			failed := installWithRetries(filePackages, 3)
			failedPackages = append(failedPackages, failed...)
			installed := diffStr(filePackages, failed)
			installedPackages = append(installedPackages, installed...)
			if ss != nil {
				if len(failed) > 0 {
					ss.AddStep(fmt.Sprintf("%s[WARN]%s Installed %d external packages (%d failed)", utils.ColorYellow, utils.ColorReset, len(installed), len(failed)))
				} else {
					ss.AddStep(fmt.Sprintf("%s[OK]%s External file packages installed (%d packages)", utils.ColorGreen, utils.ColorReset, len(installed)))
				}
			}
		}
	}

	// Packages No Recommends
	if len(suit.PackagesNoRecommends) > 0 {
		if ss != nil {
			ss.SetAction("Installing packages without recommends (%d packages)...", len(suit.PackagesNoRecommends))
		}
		failed := installNoRecommends(suit.PackagesNoRecommends)
		failedPackages = append(failedPackages, failed...)
		installed := diffStr(suit.PackagesNoRecommends, failed)
		installedPackages = append(installedPackages, installed...)
		if ss != nil {
			if len(failed) > 0 {
				ss.AddStep(fmt.Sprintf("%s[WARN]%s Installed %d packages without recommends (%d failed)", utils.ColorYellow, utils.ColorReset, len(installed), len(failed)))
			} else {
				ss.AddStep(fmt.Sprintf("%s[OK]%s Installed %d packages without recommends", utils.ColorGreen, utils.ColorReset, len(installed)))
			}
		}
	}

	// Packages Interactive
	if len(suit.PackagesInteractive) > 0 {
		if ss != nil {
			ss.SetAction("Configuring %d interactive packages...", len(suit.PackagesInteractive))
		}
		failed := installInteractive(suit.PackagesInteractive)
		failedPackages = append(failedPackages, failed...)
		installed := diffStr(suit.PackagesInteractive, failed)
		installedPackages = append(installedPackages, installed...)
		if ss != nil {
			if len(failed) > 0 {
				ss.AddStep(fmt.Sprintf("%s[WARN]%s Some interactive packages could not be installed", utils.ColorYellow, utils.ColorReset))
			} else {
				ss.AddStep(fmt.Sprintf("%s[OK]%s Interactive packages configured", utils.ColorGreen, utils.ColorReset))
			}
		}
	}

	// Packages Remove
	if len(suit.PackagesRemove) > 0 {
		if ss != nil {
			ss.SetAction("Removing %d unwanted packages...", len(suit.PackagesRemove))
		}
		removePackages(suit.PackagesRemove)
		if ss != nil {
			ss.AddStep(fmt.Sprintf("%s[OK]%s Unwanted packages removed (%d packages)", utils.ColorGreen, utils.ColorReset, len(suit.PackagesRemove)))
		}
	}

	// Sysroot Overlay
	sysrootPath := filepath.Join(dir, "sysroot")
	if _, err := os.Stat(sysrootPath); os.IsNotExist(err) {
		sysrootPath = filepath.Join(dir, "dirs")
	}
	if _, err := os.Stat(sysrootPath); err == nil {
		if ss != nil {
			ss.SetAction("Applying filesystem overlay (sysroot)...")
		}
		cmd := fmt.Sprintf("rsync -aAX %s/ /", sysrootPath)
		err := utils.ExecTee(cmd, wardrobeLogFile)
		if ss != nil {
			if err != nil {
				ss.AddStep(fmt.Sprintf("%s[WARN]%s Filesystem overlay applied with warnings", utils.ColorYellow, utils.ColorReset))
			} else {
				ss.AddStep(fmt.Sprintf("%s[OK]%s Filesystem overlay applied", utils.ColorGreen, utils.ColorReset))
			}
		}
	}

	// Finalization commands
	if len(suit.Cmds) > 0 {
		if ss != nil {
			ss.SetAction("Running finalization scripts (%d commands)...", len(suit.Cmds))
		}
		for idx, command := range suit.Cmds {
			fields := strings.Fields(command)
			cmdName := command
			if len(fields) > 0 {
				cmdName = filepath.Base(fields[0])
			}
			if ss != nil {
				ss.SetAction("[%d/%d] Finalization: %s", idx+1, len(suit.Cmds), cmdName)
			}
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
					_ = utils.ExecTee(fullCmd, wardrobeLogFile)
					continue
				}
			}
			_ = utils.ExecTee(command, wardrobeLogFile)
		}
		if ss != nil {
			ss.AddStep(fmt.Sprintf("%s[OK]%s Finalization completed", utils.ColorGreen, utils.ColorReset))
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
		logToFile("WARNING: unable to determine a non-root target user, skipping /etc/skel sync")
		return
	}

	logToFile(fmt.Sprintf("Syncing /etc/skel -> %s", userHome))
	cmd := fmt.Sprintf("rsync -a --no-o --no-g --chown=%s:%s /etc/skel/ %s/", targetUser, targetUser, userHome)
	_ = utils.ExecTee(cmd, wardrobeLogFile)
}
