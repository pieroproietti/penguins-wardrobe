package tailor

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

// getWardrobeRoot returns ~/.wardrobe for the "real" user behind this
// elevated process.
//
// Priority order:
// 1. SUDO_USER  -- set by sudo, identifies the calling user reliably.
// 2. logname    -- reads the kernel audit loginuid; survives any number
//                  of 'su' invocations, unlike environment variables that
//                  each elevation mechanism may or may not preserve.
//                  Needed for distros without sudo (e.g. Quirinux/Devuan)
//                  where 'su' is the normal way to become root and
//                  os.UserHomeDir() returns /root instead of the real home.
// 3. firstHumanUser -- scan /etc/passwd for the first UID 1000-59999 with
//                       a valid login shell, last resort before giving up.
// 4. os.UserHomeDir() -- process HOME, works when not elevated at all.
func getWardrobeRoot() (string, error) {
	var homeDir string

	if sudoUser := os.Getenv("SUDO_USER"); sudoUser != "" {
		if u, err := user.Lookup(sudoUser); err == nil {
			homeDir = u.HomeDir
		}
	}

	if homeDir == "" {
		if out, err := exec.Command("logname").Output(); err == nil {
			login := strings.TrimSpace(string(out))
			if login != "" && login != "root" {
				if u, err := user.Lookup(login); err == nil {
					homeDir = u.HomeDir
				}
			}
		}
	}

	if homeDir == "" {
		if u := firstHumanUser(); u != nil {
			homeDir = u.HomeDir
		}
	}

	if homeDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("unable to determine home directory: %v", err)
		}
		homeDir = home
	}

	return filepath.Join(homeDir, ".wardrobe"), nil
}

// getWardrobeV2Dir returns ~/.wardrobe/v2 for the "real" user if present, or ~/.wardrobe.
func getWardrobeV2Dir() (string, error) {
	root, err := getWardrobeRoot()
	if err != nil {
		return "", err
	}
	v2 := filepath.Join(root, "v2")
	if _, err := os.Stat(v2); err == nil {
		return v2, nil
	}
	return root, nil
}

// firstHumanUser scans /etc/passwd for the first real (non-system) user:
// UID between 1000 and 59999 with a valid login shell.
func firstHumanUser() *user.User {
	f, err := os.Open("/etc/passwd")
	if err != nil {
		return nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ":")
		if len(fields) < 7 {
			continue
		}
		uid, err := strconv.Atoi(fields[2])
		if err != nil || uid < 1000 || uid >= 60000 {
			continue
		}
		shell := fields[6]
		if strings.HasSuffix(shell, "nologin") || strings.HasSuffix(shell, "/false") {
			continue
		}
		if u, err := user.Lookup(fields[0]); err == nil {
			return u
		}
	}
	return nil
}
