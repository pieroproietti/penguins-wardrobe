package tailor

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/pieroproietti/penguins-wardrobe/pkg/utils"
)

// setupRepositories applica la sezione "repositories" della forma annidata
func setupRepositories(repos *Repositories, suitName string) {
	if repos == nil {
		return
	}

	if len(repos.SourcesList) > 0 {
		if err := enableAptComponents(repos.SourcesList); err != nil {
			logToFile(WarnPrefix(suitName) + "sources.list: " + err.Error())
		}
	}

	if len(repos.SourcesListD) > 0 {
		logToFile(WarnPrefix(suitName) + "running third-party repository setup commands...")
		for idx, command := range repos.SourcesListD {
			if ss := utils.GetSplitScreen(); ss != nil {
				ss.SetAction(fmt.Sprintf("Third-party repo [%d/%d]", idx+1, len(repos.SourcesListD)))
			}
			if err := utils.ExecTee(command, wardrobeLogFile); err != nil {
				logToFile(WarnPrefix(suitName) + "repository command failed: " + command + ": " + err.Error())
			}
		}
	}

	if repos.Update {
		logToFile(WarnPrefix(suitName) + "apt-get update...")
		if ss := utils.GetSplitScreen(); ss != nil {
			ss.SetAction("Updating package index (apt-get update)...")
		}
		_ = utils.ExecTee("apt-get update", wardrobeLogFile)
	}

	if repos.Upgrade {
		logToFile(WarnPrefix(suitName) + "apt-get upgrade...")
		if ss := utils.GetSplitScreen(); ss != nil {
			ss.SetAction("Upgrading packages (apt-get upgrade)...")
		}
		_ = utils.ExecTee("UCF_FORCE_CONFFOLD=1 DEBIAN_FRONTEND=readline apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' upgrade -y", wardrobeLogFile)
	}
}

// WarnPrefix genera un prefisso omogeneo per i log di questo pacchetto.
func WarnPrefix(suitName string) string {
	return "[" + suitName + "] "
}

// enableAptComponents assicura che i componenti richiesti siano abilitati
func enableAptComponents(components []string) error {
	const path = "/etc/apt/sources.list"

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	debLineRe := regexp.MustCompile(`^(deb|deb-src)\s+(\S+)\s+(\S+)\s+(.*)$`)

	lines := strings.Split(string(data), "\n")
	changed := false
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		m := debLineRe.FindStringSubmatch(trimmed)
		if m == nil {
			continue
		}
		existing := strings.Fields(m[4])
		existingSet := make(map[string]struct{}, len(existing))
		for _, c := range existing {
			existingSet[c] = struct{}{}
		}

		added := false
		for _, want := range components {
			if _, ok := existingSet[want]; !ok {
				existing = append(existing, want)
				existingSet[want] = struct{}{}
				added = true
			}
		}

		if added {
			lines[i] = strings.Join([]string{m[1], m[2], m[3], strings.Join(existing, " ")}, " ")
			changed = true
		}
	}

	if !changed {
		return nil
	}

	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}
