package tailor

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/pieroproietti/penguins-wardrobe/pkg/distro"
	"github.com/pieroproietti/penguins-wardrobe/pkg/utils"
	"gopkg.in/yaml.v3"
)

const wardrobeLogFile = "/var/log/wardrobe.log"

func logToFile(message string) {
	logPath := wardrobeLogFile
	if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
		// fallback
	}
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	f.WriteString(fmt.Sprintf("[%s] %s\n", timestamp, message))
}

func findYaml(costumePath string) string {
	candidates := []string{
		"index.yaml",
		"index.yml",
	}

	d := distro.NewDistro()
	if d != nil {
		if d.DistroID != "" {
			candidates = append(candidates, strings.ToLower(d.DistroID)+".yaml", strings.ToLower(d.DistroID)+".yml")
		}
		if d.DistroLike != "" {
			candidates = append(candidates, strings.ToLower(d.DistroLike)+".yaml", strings.ToLower(d.DistroLike)+".yml")
		}
		if d.FamilyID != "" {
			candidates = append(candidates, strings.ToLower(d.FamilyID)+".yaml", strings.ToLower(d.FamilyID)+".yml")
		}
		if strings.EqualFold(d.FamilyID, "archlinux") {
			candidates = append(candidates, "arch.yaml", "arch.yml")
		}
		if strings.EqualFold(d.FamilyID, "debian") {
			candidates = append(candidates, "debian.yaml", "debian.yml", "ubuntu.yaml", "devuan.yaml")
		}
	}

	// Standard distro fallbacks
	candidates = append(candidates, "debian.yaml", "debian.yml", "arch.yaml", "alpine.yaml", "fedora.yaml", "opensuse.yaml")

	seen := make(map[string]struct{})
	for _, c := range candidates {
		if _, ok := seen[c]; ok {
			continue
		}
		seen[c] = struct{}{}
		fullPath := filepath.Join(costumePath, c)
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath
		}
	}

	return ""
}

func loadSuit(yamlFile string) (*Suit, error) {
	if yamlFile == "" {
		return nil, fmt.Errorf("costume/accessory definition yaml file not found")
	}
	data, err := os.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}
	var suit Suit
	if err := yaml.Unmarshal(data, &suit); err != nil {
		return nil, err
	}
	suit.normalize()
	return &suit, nil
}

func getAvailablePackages() map[string]struct{} {
	available := make(map[string]struct{})
	if _, err := exec.LookPath("apt-cache"); err != nil {
		return nil
	}
	logToFile("Updating available packages database...")
	cmd := exec.Command("/usr/bin/apt-cache", "pkgnames")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return available
	}
	if err := cmd.Start(); err != nil {
		return available
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			available[line] = struct{}{}
		}
	}
	cmd.Wait()
	return available
}

// normalizePkgName strips the ":arch" multi-arch qualifier
func normalizePkgName(name string) string {
	name = strings.TrimSpace(name)
	if i := strings.Index(name, ":"); i != -1 {
		name = name[:i]
	}
	return name
}

// isInteractiveTerminal checks if stdin is connected to a real terminal.
// This is used to decide whether to show interactive prompts or fall back
// to noninteractive mode.
func isInteractiveTerminal() bool {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

// batchSize caps how many packages go into a single apt-get invocation.
// Installing in smaller batches keeps each dpkg transaction's trigger
// processing small, and each successfully completed batch is durably
// installed on disk, so a crash mid-way only loses the current batch.
const batchSize = 20

func cleanAptProgress(line string) string {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "Get:") {
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			return "Download " + fields[1]
		}
		return "Download..."
	}
	if strings.HasPrefix(line, "Unpacking ") {
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			return "Unpack " + fields[1]
		}
		return "Unpack..."
	}
	if strings.HasPrefix(line, "Setting up ") {
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			return "Setup " + fields[1]
		}
		return "Setup..."
	}
	if strings.HasPrefix(line, "Processing triggers for ") {
		pkg := strings.TrimPrefix(line, "Processing triggers for ")
		if i := strings.Index(pkg, " "); i != -1 {
			pkg = pkg[:i]
		}
		return "Trigger (" + pkg + ")"
	}
	if strings.HasPrefix(line, "Preparing to unpack ") {
		fields := strings.Fields(line)
		if len(fields) >= 4 {
			return "Prep " + fields[3]
		}
	}
	if strings.HasPrefix(line, "Selecting ") {
		fields := strings.Fields(line)
		if len(fields) >= 4 {
			return "Select " + fields[3]
		}
	}
	if strings.HasPrefix(line, "Hit:") || strings.HasPrefix(line, "Ign:") {
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			return "Fetch " + fields[1]
		}
	}
	return ""
}

// installWithRetries installs packages, falling back to one-by-one on failure
func installWithRetries(packages []string, retries int) []string {
	return installPackagesImpl(packages, retries, false)
}

// installNoRecommends installs packages with --no-install-recommends
func installNoRecommends(packages []string) []string {
	return installPackagesImpl(packages, 3, true)
}

func installPackagesImpl(packages []string, retries int, noRecommends bool) []string {
	if len(packages) == 0 {
		return nil
	}
	if _, err := exec.LookPath("apt-get"); err != nil {
		printAiPrompt(packages)
		return nil
	}

	available := getAvailablePackages()
	var toInstall []string
	var missing []string
	if available != nil {
		for _, pkg := range packages {
			cleanPkg := normalizePkgName(pkg)
			if _, ok := available[cleanPkg]; ok {
				toInstall = append(toInstall, pkg)
			} else {
				missing = append(missing, pkg)
			}
		}
	} else {
		toInstall = packages
	}

	if len(missing) > 0 {
		logToFile(fmt.Sprintf("WARNING: %d packages skipped (not found): %v", len(missing), missing))
	}
	if len(toInstall) == 0 {
		logToFile("No valid packages to install.")
		return missing
	}

	flags := "-y"
	if noRecommends {
		flags = "-y --no-install-recommends"
	}

	totalBatches := (len(toInstall) + batchSize - 1) / batchSize
	if totalBatches > 1 {
		logToFile(fmt.Sprintf("Installing %d packages in %d batches of up to %d...", len(toInstall), totalBatches, batchSize))
	}

	var failed []string
	for start := 0; start < len(toInstall); start += batchSize {
		end := start + batchSize
		if end > len(toInstall) {
			end = len(toInstall)
		}
		batch := toInstall[start:end]
		batchNum := start/batchSize + 1
		if ss := utils.GetSplitScreen(); ss != nil && totalBatches > 1 {
			ss.SetAction("Installing packages (batch %d/%d)...", batchNum, totalBatches)
		}
		logToFile(fmt.Sprintf("Batch %d/%d (packages %d-%d of %d): %v", batchNum, totalBatches, start+1, end, len(toInstall), batch))
		failed = append(failed, installBatchWithFallback(batch, retries, flags)...)
	}
	return append(missing, failed...)
}

// installBatchWithFallback installs a single batch in bulk, falling back
// to one-by-one installation within the batch if the bulk call fails.
func installBatchWithFallback(batch []string, retries int, flags string) []string {
	// License-prompt packages must never go through the noninteractive
	// path: their preinst aborts and poisons dpkg for every later batch.
	var clean []string
	for _, p := range batch {
		if !isLicensePrompt(p) {
			clean = append(clean, p)
		}
	}
	batch = clean
	if len(batch) == 0 {
		return nil
	}
	
	// Use readline frontend if we have an interactive terminal
	debconfFrontend := "readline"
	if !isInteractiveTerminal() {
		debconfFrontend = "noninteractive"
	}
	
	pkgString := strings.Join(batch, " ")
	cmd := fmt.Sprintf("UCF_FORCE_CONFFOLD=1 DEBIAN_FRONTEND=%s apt-get install -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' %s %s", debconfFrontend, flags, pkgString)
	logToFile(fmt.Sprintf("Installing batch of %d packages: %s", len(batch), pkgString))
	if err := utils.ExecTee(cmd, wardrobeLogFile); err == nil {
		logToFile("✅ Batch installed.")
		return nil
	}

	// Heal dpkg state before retrying
	healDpkgState()

	logToFile("⚠️  Retrying package by package to isolate failures...")
	pending := batch
	for attempt := 1; attempt <= retries && len(pending) > 0; attempt++ {
		var stillFailing []string
		for _, pkg := range pending {
			if ss := utils.GetSplitScreen(); ss != nil {
				ss.SetAction("Retrying package: %s (attempt %d/%d)", pkg, attempt, retries)
			}
			singleCmd := fmt.Sprintf("UCF_FORCE_CONFFOLD=1 DEBIAN_FRONTEND=%s apt-get install -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' %s %s", debconfFrontend, flags, pkg)
			if err := utils.ExecTee(singleCmd, wardrobeLogFile); err != nil {
				// Double-check with dpkg before believing the failure
				if isPackageInstalled(pkg) {
					logToFile(fmt.Sprintf("ℹ️  apt-get reported an error installing %s, but dpkg confirms it is installed correctly.", pkg))
				} else {
					stillFailing = append(stillFailing, pkg)
				}
			}
		}
		pending = stillFailing
		if len(pending) > 0 && attempt < retries {
			logToFile(fmt.Sprintf("⚠️  %d packages still failing after attempt %d/%d, retrying: %v", len(pending), attempt, retries, pending))
		}
	}

	if len(pending) > 0 {
		logToFile(fmt.Sprintf("⚠️  %d packages could not be installed: %v", len(pending), pending))
	} else {
		logToFile("✅ All packages in batch installed successfully (one by one).")
	}
	return pending
}

// isPackageInstalled reports whether dpkg considers pkg to be correctly and fully installed.
func isPackageInstalled(pkg string) bool {
	installed, err := currentlyInstalledPackages()
	if err != nil {
		return false
	}
	_, ok := installed[normalizePkgName(pkg)]
	return ok
}

// installInteractive installs packages without suppressing debconf prompts.
func installInteractive(packages []string) []string {
	if len(packages) == 0 {
		return nil
	}

	available := getAvailablePackages()
	var toInstall []string
	var missing []string
	if available != nil {
		for _, pkg := range packages {
			cleanPkg := normalizePkgName(pkg)
			if _, ok := available[cleanPkg]; ok {
				toInstall = append(toInstall, pkg)
			} else {
				missing = append(missing, pkg)
			}
		}
	} else {
		toInstall = packages
	}

	if len(missing) > 0 {
		logToFile(fmt.Sprintf("WARNING: %d interactive packages skipped (not found): %v", len(missing), missing))
	}
	if len(toInstall) == 0 {
		return missing
	}

	// Use readline frontend for interactive packages so prompts are shown
	debconfFrontend := "readline"
	if !isInteractiveTerminal() {
		debconfFrontend = "noninteractive"
	}

	pkgString := strings.Join(toInstall, " ")
	cmd := fmt.Sprintf("DEBIAN_FRONTEND=%s apt-get install -o Dpkg::Options::='--force-confold' -y %s", debconfFrontend, pkgString)
	logToFile(fmt.Sprintf("Installing interactive packages: %s", pkgString))
	if err := utils.ExecTee(cmd, wardrobeLogFile); err != nil {
		var stillFailing []string
		for _, pkg := range toInstall {
			if !isPackageInstalled(pkg) {
				stillFailing = append(stillFailing, pkg)
			}
		}
		if len(stillFailing) > 0 {
			logToFile(fmt.Sprintf("⚠️  Some interactive packages could not be installed: %v", stillFailing))
		}
		return append(missing, stillFailing...)
	}
	return missing
}

// removePackages removes packages that the vendor does not want on the system.
func removePackages(packages []string) {
	if len(packages) == 0 {
		return
	}

	kernel := currentKernelPackage()
	var safe []string
	for _, p := range packages {
		if kernel != "" && normalizePkgName(p) == kernel {
			logToFile(fmt.Sprintf("⚠️  Refusing to remove the running kernel package: %s", p))
			continue
		}
		safe = append(safe, p)
	}
	if len(safe) == 0 {
		return
	}

	pkgString := strings.Join(safe, " ")
	cmd := fmt.Sprintf("DEBIAN_FRONTEND=readline apt-get remove -o Dpkg::Options::='--force-confold' -y %s", pkgString)
	logToFile(fmt.Sprintf("Removing packages: %s", pkgString))
	if err := utils.ExecTee(cmd, wardrobeLogFile); err != nil {
		logToFile(fmt.Sprintf("⚠️  Some packages could not be removed: %v", err))
	}
	_ = utils.ExecTee("DEBIAN_FRONTEND=readline apt-get autoremove -o Dpkg::Options::='--force-confold' -y", wardrobeLogFile)
}

func printAiPrompt(packages []string) {
	d := distro.NewDistro()
	logToFile(fmt.Sprintf("System %s detected (Non-Debian). Generating prompt and AIPrompt.txt file...", d.DistroLike))

	gpuCmd := "lspci -k | grep -A 2 -E 'VGA|3D'"
	gpuInfo, _ := exec.Command("sh", "-c", gpuCmd).Output()
	sessionCmd := "ls /usr/share/xsessions/ 2>/dev/null"
	sessions, _ := exec.Command("sh", "-c", sessionCmd).Output()

	var sb strings.Builder
	sb.WriteString("\n--- AI ASSISTANT PROMPT ---\n")
	sb.WriteString(fmt.Sprintf("I am using %s (base %s).\n", d.DistroID, d.DistroLike))
	sb.WriteString(fmt.Sprintf("I need to install and configure these packages:\n%s\n\n", strings.Join(packages, " ")))
	sb.WriteString("HARDWARE INFO (for video drivers and KMS):\n")
	if len(gpuInfo) > 0 {
		sb.WriteString(string(gpuInfo))
	} else {
		sb.WriteString("No VGA info found (pciutils not installed?).\n")
	}
	sb.WriteString("\nAVAILABLE DESKTOP SESSIONS:\n")
	if len(sessions) > 0 {
		sb.WriteString(string(sessions))
	} else {
		sb.WriteString("No sessions found in /usr/share/xsessions/\n")
	}
	sb.WriteString("\nPlease give me the exact command to install the equivalent packages on this distro and the steps needed to configure LightDM correctly.\n")
	sb.WriteString("----------------------------------------\n")

	promptContent := sb.String()
	utils.LogNormal("\n%s%s%s", utils.ColorCyan, promptContent, utils.ColorReset)

	userHome, _ := os.UserHomeDir()
	sudoUser := os.Getenv("SUDO_USER")
	if sudoUser != "" {
		userHome = filepath.Join("/home", sudoUser)
	}

	promptFile := filepath.Join(userHome, "AIPrompt.txt")
	err := os.WriteFile(promptFile, []byte(promptContent), 0644)
	if err != nil {
		logToFile(fmt.Sprintf("Error creating AIPrompt.txt: %v", err))
	} else {
		if sudoUser != "" {
			utils.Exec(fmt.Sprintf("chown %s:%s %s", sudoUser, sudoUser, promptFile))
		}
		logToFile(fmt.Sprintf("✅ AIPrompt.txt file generated at: %s", promptFile))
		utils.LogNormal("Prompt file generated in Home: %s%s%s\n", utils.ColorYellow, promptFile, utils.ColorReset)
	}
}

// licensePromptPackages holds suit.PackagesInteractive: packages whose
// preinst asks a license question that cannot be answered noninteractively.
var licensePromptPackages []string

// SetLicensePromptPackages is called by Wear() before any install starts.
func SetLicensePromptPackages(pkgs []string) { licensePromptPackages = pkgs }

func isLicensePrompt(pkg string) bool {
	c := normalizePkgName(pkg)
	for _, p := range licensePromptPackages {
		if normalizePkgName(p) == c {
			return true
		}
	}
	return false
}

// healDpkgState repairs a poisoned dpkg state.
func healDpkgState() {
	// First try to configure what we can without interaction
	utils.Exec("DEBIAN_FRONTEND=noninteractive dpkg --configure -a --force-confold")
	utils.Exec("DEBIAN_FRONTEND=noninteractive apt-get install -f -y")
	
	// Then try with readline if we have a terminal
	if isInteractiveTerminal() {
		utils.Exec("DEBIAN_FRONTEND=readline dpkg --configure -a")
	}
	
	for _, p := range licensePromptPackages {
		if !isPackageInstalled(p) {
			logToFile(fmt.Sprintf("⚠️  Purging half-configured license package %s so the rest of the system can heal...", p))
			utils.Exec(fmt.Sprintf("DEBIAN_FRONTEND=noninteractive dpkg --purge --force-remove-reinstreq --force-depends %s", p))
		}
	}
	utils.Exec("DEBIAN_FRONTEND=noninteractive apt-get install -f -y")
}