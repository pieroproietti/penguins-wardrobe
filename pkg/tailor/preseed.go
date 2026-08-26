package tailor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pieroproietti/penguins-wardrobe/pkg/utils"
)

// findPreseed looks specifically and exclusively for a "packages.preseed" file
// inside the given costume or accessory directory.
func findPreseed(dir string) string {
	preseedFile := filepath.Join(dir, "packages.preseed")
	if info, err := os.Stat(preseedFile); err == nil && !info.IsDir() {
		return preseedFile
	}
	return ""
}

// applyPreseed loads answers into the debconf database using debconf-set-selections.
func applyPreseed(preseedFile, suitName string) error {
	if preseedFile == "" {
		return nil
	}

	// Only applicable on Debian/Devuan/Ubuntu systems where apt-get or debconf is present
	if _, err := exec.LookPath("apt-get"); err != nil {
		return nil
	}

	// Ensure debconf-set-selections is available (from debconf-utils)
	if _, err := exec.LookPath("debconf-set-selections"); err != nil {
		logToFile(WarnPrefix(suitName) + "debconf-set-selections not found, installing debconf-utils...")
		_ = utils.ExecLog("DEBIAN_FRONTEND=noninteractive apt-get install -o Dpkg::Options::='--force-confold' -o Dpkg::Use-Pty=0 -y debconf-utils", wardrobeLogFile)
	}

	logToFile(WarnPrefix(suitName) + "applying debconf preseed: " + preseedFile)
	cmd := fmt.Sprintf("debconf-set-selections %s", preseedFile)
	return utils.ExecLog(cmd, wardrobeLogFile)
}
