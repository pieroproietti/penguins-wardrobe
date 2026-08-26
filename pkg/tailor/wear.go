package tailor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pieroproietti/penguins-wardrobe/pkg/utils"
)

var isInteractiveMode bool

// SetInteractiveMode enables or disables full interactive terminal mode
func SetInteractiveMode(enabled bool) {
	isInteractiveMode = enabled
}

// IsInteractiveMode returns whether interactive mode is currently active
func IsInteractiveMode() bool {
	return isInteractiveMode
}

func Wear(costumeName string, noAcc bool, noFirm bool, interactive bool) error {
	SetInteractiveMode(interactive)

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

	if !isDirectAccessory {
		versionStr := ""
		if suit.Release != "" {
			versionStr = fmt.Sprintf(" (v%s)", suit.Release)
		}
		utils.PrintBanner("👗", fmt.Sprintf("COSTUME: %s%s", suit.Name, versionStr), suit.Description)
	} else {
		utils.PrintBanner("👝", fmt.Sprintf("ACCESSORY: %s", suit.Name), suit.Description)
	}

	// DKMS safety: ensure headers for running kernel are present
	if isInteractiveMode {
		fmt.Printf("  %s--> Checking kernel headers for DKMS...%s\n", utils.ColorCyan, utils.ColorReset)
		if err := ensureKernelHeaders(); err != nil {
			fmt.Printf("  %s[WARN] Kernel headers verification completed with warnings%s\n", utils.ColorYellow, utils.ColorReset)
		} else {
			fmt.Printf("  %s[OK] Kernel headers verified%s\n", utils.ColorGreen, utils.ColorReset)
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
		utils.PrintSection("👝", fmt.Sprintf("ACCESSORIES (%d items)", len(suit.Accessories)))
		for idx, accName := range suit.Accessories {
			if noFirm && (accName == "firmwares" || strings.Contains(accName, "firmware")) {
				fmt.Printf("\n  %s[INFO] Skipping firmware accessory '%s' due to --no-firm flag%s\n", utils.ColorYellow, accName, utils.ColorReset)
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
					utils.PrintSubSection("-->", fmt.Sprintf("[%d/%d] Accessory: %s", idx+1, len(suit.Accessories), accName))
					accInstalled, accFailed, _ := applySuit(accDir, accSuit)
					installedPackages = append(installedPackages, accInstalled...)
					failedPackages = append(failedPackages, accFailed...)
				} else {
					fmt.Printf("  %s[WARN] Could not load accessory '%s': %v%s\n", utils.ColorYellow, accName, err, utils.ColorReset)
				}
			} else {
				fmt.Printf("  %s[WARN] Accessory '%s' not found in %s%s\n", utils.ColorYellow, accName, accDir, utils.ColorReset)
			}
		}
	}

	var purgedPackages []string
	var failedPurges []string

	installedBefore, _ := currentlyInstalledPackages()

	// Install everything in the manifest that's missing
	if manifestPath := resolveDistroManifest(costumeDir, suit.PackagesManifest); manifestPath != "" {
		utils.PrintSection("📋", "DECLARATIVE MANIFEST RECONCILIATION")
		if targetManifest, err := loadPackageManifest(manifestPath); err == nil {
			var spMan *utils.Spinner
			if !isInteractiveMode {
				spMan = utils.NewSpinner(fmt.Sprintf("Reconciling %d manifest packages...", len(targetManifest)))
				spMan.Start()
			} else {
				fmt.Printf("  %s--> Reconciling %d manifest packages (%s)...%s\n", utils.ColorCyan, len(targetManifest), filepath.Base(manifestPath), utils.ColorReset)
			}
			manifestFailed := installWithRetries(targetManifest, 3, spMan)
			failedPackages = append(failedPackages, manifestFailed...)
			installedPackages = append(installedPackages, diffStr(targetManifest, manifestFailed)...)
			if spMan != nil {
				if len(manifestFailed) > 0 {
					spMan.Warn("Manifest reconciled with %d missing packages", len(manifestFailed))
				} else {
					spMan.Success("Manifest packages reconciled (%d packages)", len(targetManifest))
				}
			} else {
				if len(manifestFailed) > 0 {
					fmt.Printf("  %s[WARN] Manifest reconciled with %d missing packages%s\n", utils.ColorYellow, len(manifestFailed), utils.ColorReset)
				} else {
					fmt.Printf("  %s[OK] Manifest packages reconciled (%d packages)%s\n", utils.ColorGreen, len(targetManifest), utils.ColorReset)
				}
			}
		} else {
			fmt.Printf("  %s[WARN] Could not read packages_manifest %s: %v%s\n", utils.ColorYellow, manifestPath, err, utils.ColorReset)
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
		spPurge := utils.NewSpinner(fmt.Sprintf("Purging %d packages absent from manifest...", len(removeList)))
		spPurge.Start()
		purgeExplicit(removeList)
		spPurge.Success("Declarative purge completed")
	}

	// DKMS healing
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
		spUser := utils.NewSpinner(fmt.Sprintf("Synchronizing user environment (/etc/skel -> /home/%s)...", targetUser))
		spUser.Start()
		copySkelToUser()
		spUser.Success("User environment synchronized (%s)", targetUser)
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
	if isInteractiveMode {
		return utils.ExecTee("DEBIAN_FRONTEND=readline apt-get install -o Dpkg::Options::='--force-confold' -y "+pkgs, wardrobeLogFile)
	}
	return utils.ExecLog("DEBIAN_FRONTEND=noninteractive apt-get install -o Dpkg::Use-Pty=0 -y "+pkgs, wardrobeLogFile)
}

// healAndRetryFailed repairs the half-configured dpkg state and retries failed packages.
func healAndRetryFailed(failed []string) []string {
	if len(failed) == 0 {
		return nil
	}

	logToFile("Healing dpkg state before retrying failed packages...")
	utils.ExecLog("dpkg --configure -a", wardrobeLogFile)
	utils.ExecLog("DEBIAN_FRONTEND=noninteractive apt-get install -f -o Dpkg::Use-Pty=0 -y", wardrobeLogFile)

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
	installWithRetries(retry, 1, nil)

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

	// Preseed (debconf-set-selections)
	if preseedFile := findPreseed(dir); preseedFile != "" {
		if isInteractiveMode {
			fmt.Printf("  %s--> Applying preseed selections (%s)...%s\n", utils.ColorCyan, filepath.Base(preseedFile), utils.ColorReset)
			if err := applyPreseed(preseedFile, suit.Name); err != nil {
				fmt.Printf("  %s[WARN] Preseed selections applied with warnings%s\n", utils.ColorYellow, utils.ColorReset)
			} else {
				fmt.Printf("  %s[OK] Preseed selections applied (%s)%s\n", utils.ColorGreen, filepath.Base(preseedFile), utils.ColorReset)
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
		var spRepo *utils.Spinner
		if !isInteractiveMode {
			spRepo = utils.NewSpinner("Configuring package repositories & updating cache...")
			spRepo.Start()
		} else {
			fmt.Printf("  %s--> Configuring package repositories & updating cache...%s\n", utils.ColorCyan, utils.ColorReset)
		}
		setupRepositories(suit.Sequence.Repositories, suit.Name, spRepo)
		if spRepo != nil {
			spRepo.Success("Repositories configured & updated")
		} else {
			fmt.Printf("  %s[OK] Repositories configured & updated%s\n", utils.ColorGreen, utils.ColorReset)
		}
	}

	// Packages
	if len(suit.Packages) > 0 {
		var spPkg *utils.Spinner
		if !isInteractiveMode {
			spPkg = utils.NewSpinner(fmt.Sprintf("Installing packages (%d packages)...", len(suit.Packages)))
			spPkg.Start()
		} else {
			fmt.Printf("  %s--> Installing %d packages...%s\n", utils.ColorCyan, len(suit.Packages), utils.ColorReset)
		}
		failed := installWithRetries(suit.Packages, 3, spPkg)
		failedPackages = append(failedPackages, failed...)
		installed := diffStr(suit.Packages, failed)
		installedPackages = append(installedPackages, installed...)
		if spPkg != nil {
			if len(failed) > 0 {
				spPkg.Warn("Installed %d packages (%d could not be installed)", len(installed), len(failed))
			} else {
				spPkg.Success("Installed %d packages", len(installed))
			}
		} else {
			if len(failed) > 0 {
				fmt.Printf("  %s[WARN] Installed %d packages (%d failed)%s\n", utils.ColorYellow, len(installed), len(failed), utils.ColorReset)
			} else {
				fmt.Printf("  %s[OK] Installed %d packages%s\n", utils.ColorGreen, len(installed), utils.ColorReset)
			}
		}
	}

	// External install file
	if installPath := findManifestPath(dir, suit.PackagesInstallFile); installPath != "" {
		if filePackages, err := loadPackageManifest(installPath); err == nil {
			var spExt *utils.Spinner
			if !isInteractiveMode {
				spExt = utils.NewSpinner(fmt.Sprintf("Installing %d packages from external file (%s)...", len(filePackages), filepath.Base(installPath)))
				spExt.Start()
			} else {
				fmt.Printf("  %s--> Installing %d packages from external file (%s)...%s\n", utils.ColorCyan, len(filePackages), filepath.Base(installPath), utils.ColorReset)
			}
			failed := installWithRetries(filePackages, 3, spExt)
			failedPackages = append(failedPackages, failed...)
			installed := diffStr(filePackages, failed)
			installedPackages = append(installedPackages, installed...)
			if spExt != nil {
				if len(failed) > 0 {
					spExt.Warn("Installed %d packages from external file (%d failed)", len(installed), len(failed))
				} else {
					spExt.Success("External file packages installed (%d packages)", len(filePackages))
				}
			} else {
				if len(failed) > 0 {
					fmt.Printf("  %s[WARN] Installed %d packages from external file (%d failed)%s\n", utils.ColorYellow, len(installed), len(failed), utils.ColorReset)
				} else {
					fmt.Printf("  %s[OK] External file packages installed (%d packages)%s\n", utils.ColorGreen, len(installed), utils.ColorReset)
				}
			}
		}
	}

	// Packages No Recommends
	if len(suit.PackagesNoRecommends) > 0 {
		var spNoRec *utils.Spinner
		if !isInteractiveMode {
			spNoRec = utils.NewSpinner(fmt.Sprintf("Installing packages without recommends (%d packages)...", len(suit.PackagesNoRecommends)))
			spNoRec.Start()
		} else {
			fmt.Printf("  %s--> Installing %d packages without recommends...%s\n", utils.ColorCyan, len(suit.PackagesNoRecommends), utils.ColorReset)
		}
		failed := installNoRecommends(suit.PackagesNoRecommends, spNoRec)
		failedPackages = append(failedPackages, failed...)
		installed := diffStr(suit.PackagesNoRecommends, failed)
		installedPackages = append(installedPackages, installed...)
		if spNoRec != nil {
			if len(failed) > 0 {
				spNoRec.Warn("Installed %d packages without recommends (%d failed)", len(installed), len(failed))
			} else {
				spNoRec.Success("Installed %d packages without recommends", len(installed))
			}
		} else {
			if len(failed) > 0 {
				fmt.Printf("  %s[WARN] Installed %d packages without recommends (%d failed)%s\n", utils.ColorYellow, len(installed), len(failed), utils.ColorReset)
			} else {
				fmt.Printf("  %s[OK] Installed %d packages without recommends%s\n", utils.ColorGreen, len(installed), utils.ColorReset)
			}
		}
	}

	// Packages Interactive
	if len(suit.PackagesInteractive) > 0 {
		fmt.Printf("  %s[INFO] Installing interactive packages (prompts may appear):%s\n", utils.ColorCyan, utils.ColorReset)
		failed := installInteractive(suit.PackagesInteractive)
		failedPackages = append(failedPackages, failed...)
		installed := diffStr(suit.PackagesInteractive, failed)
		installedPackages = append(installedPackages, installed...)
		if len(failed) > 0 {
			fmt.Printf("  %s[WARN] Some interactive packages could not be installed%s\n", utils.ColorYellow, utils.ColorReset)
		} else {
			fmt.Printf("  %s[OK] Interactive packages configured%s\n", utils.ColorGreen, utils.ColorReset)
		}
	}

	// Packages Remove
	if len(suit.PackagesRemove) > 0 {
		var spRem *utils.Spinner
		if !isInteractiveMode {
			spRem = utils.NewSpinner(fmt.Sprintf("Removing unwanted packages (%d packages)...", len(suit.PackagesRemove)))
			spRem.Start()
		} else {
			fmt.Printf("  %s--> Removing unwanted packages (%d packages)...%s\n", utils.ColorCyan, len(suit.PackagesRemove), utils.ColorReset)
		}
		removePackages(suit.PackagesRemove)
		if spRem != nil {
			spRem.Success("Unwanted packages removed (%d packages)", len(suit.PackagesRemove))
		} else {
			fmt.Printf("  %s[OK] Unwanted packages removed (%d packages)%s\n", utils.ColorGreen, len(suit.PackagesRemove), utils.ColorReset)
		}
	}

	// Sysroot Overlay
	sysrootPath := filepath.Join(dir, "sysroot")
	if _, err := os.Stat(sysrootPath); os.IsNotExist(err) {
		sysrootPath = filepath.Join(dir, "dirs")
	}
	if _, err := os.Stat(sysrootPath); err == nil {
		var spSys *utils.Spinner
		if !isInteractiveMode {
			spSys = utils.NewSpinner("Applying filesystem overlay (sysroot)...")
			spSys.Start()
		} else {
			fmt.Printf("  %s--> Applying filesystem overlay (sysroot)...%s\n", utils.ColorCyan, utils.ColorReset)
		}
		cmd := fmt.Sprintf("rsync -aAX %s/ /", sysrootPath)
		var err error
		if isInteractiveMode {
			err = utils.ExecTee(cmd, wardrobeLogFile)
		} else {
			err = utils.ExecLog(cmd, wardrobeLogFile)
		}
		if err != nil {
			if spSys != nil {
				spSys.Warn("Filesystem overlay applied with warnings")
			} else {
				fmt.Printf("  %s[WARN] Filesystem overlay applied with warnings%s\n", utils.ColorYellow, utils.ColorReset)
			}
		} else {
			if spSys != nil {
				spSys.Success("Filesystem overlay applied")
			} else {
				fmt.Printf("  %s[OK] Filesystem overlay applied%s\n", utils.ColorGreen, utils.ColorReset)
			}
		}
	}

	// Finalization commands
	if len(suit.Cmds) > 0 {
		var spCmds *utils.Spinner
		if !isInteractiveMode {
			spCmds = utils.NewSpinner(fmt.Sprintf("Running finalization scripts (%d commands)...", len(suit.Cmds)))
			spCmds.Start()
		} else {
			fmt.Printf("  %s--> Running finalization scripts (%d commands)...%s\n", utils.ColorCyan, len(suit.Cmds), utils.ColorReset)
		}
		for idx, command := range suit.Cmds {
			fields := strings.Fields(command)
			cmdName := command
			if len(fields) > 0 {
				cmdName = filepath.Base(fields[0])
			}
			if spCmds != nil {
				spCmds.UpdateSubtext(fmt.Sprintf("[%d/%d] %s", idx+1, len(suit.Cmds), cmdName))
			} else {
				fmt.Printf("    [%d/%d] %s\n", idx+1, len(suit.Cmds), cmdName)
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
					if isInteractiveMode {
						utils.ExecTee(fullCmd, wardrobeLogFile)
					} else {
						utils.ExecLog(fullCmd, wardrobeLogFile)
					}
					continue
				}
			}
			if isInteractiveMode {
				utils.ExecTee(command, wardrobeLogFile)
			} else {
				utils.ExecLog(command, wardrobeLogFile)
			}
		}
		if spCmds != nil {
			spCmds.Success("Finalization completed")
		} else {
			fmt.Printf("  %s[OK] Finalization completed%s\n", utils.ColorGreen, utils.ColorReset)
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
	utils.ExecLog(cmd, wardrobeLogFile)
}
