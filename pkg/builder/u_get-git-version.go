// Copyright 2026 Piero Proietti <piero.proietti@gmail.com>.
// All rights reserved.

package builder

import (
	"os/exec"
	"strings"
)

func getGitVersion() (string, string) {
	// 1. Get the nearest tag (baseVer).
	outTag, err := exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
	baseVer := strings.TrimPrefix(strings.TrimSpace(string(outTag)), "v")

	if err != nil || baseVer == "" {
		relNum := "1"
		if outCount, err := exec.Command("git", "rev-list", "--count", "HEAD").Output(); err == nil {
			if n := strings.TrimSpace(string(outCount)); n != "" {
				relNum = n
			}
		}
		return "0.1.0", relNum
	}

	// 2. Count commits since the tag (relNum)
	outRel, err := exec.Command("git", "rev-list", "--count", "HEAD", "--not", "--tags").Output()
	relNum := "1"
	if err == nil {
		relNum = strings.TrimSpace(string(outRel))
		if relNum == "0" {
			relNum = "1"
		}
	}

	// 3. Sanitize
	baseVer = strings.ReplaceAll(baseVer, "-", ".")
	baseVer = strings.ReplaceAll(baseVer, "_", ".")

	return baseVer, relNum
}
