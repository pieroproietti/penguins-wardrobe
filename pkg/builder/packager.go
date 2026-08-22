// Copyright 2026 Piero Proietti <piero.proietti@gmail.com>.
// All rights reserved.

package builder

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pieroproietti/penguins-wardrobe/pkg/context"
	"github.com/pieroproietti/penguins-wardrobe/pkg/utils"
)

func packager(ctx context.RuntimeContext, dist string, data RecipeData) {
	utils.LogNormal("[build] Packager: building package for %s...", dist)

	stage := ctx.StageDir
	var cmd *exec.Cmd
	var pkgFileName string

	switch dist {
	case "alpine":
		pkgFileName = fmt.Sprintf("penguins-wardrobe-%s-r%s.apk", data.BaseVersion, data.Rel)

		apkOutDir := filepath.Join(stage, "APK")
		os.MkdirAll(apkOutDir, 0755)

		cmd = exec.Command("abuild", "-fr")
		cmd.Dir = stage
		cmd.Env = append(os.Environ(), fmt.Sprintf("REPODEST=%s", apkOutDir))

	case "arch", "manjaro":
		pkgFileName = fmt.Sprintf("penguins-wardrobe-%s-%s-x86_64.pkg.tar.zst", data.BaseVersion, data.Rel)
		cmd = exec.Command("makepkg", "-s", "-f", "--noconfirm")
		cmd.Dir = stage

		absStage, err := filepath.Abs(stage)
		if err != nil {
			absStage = stage
		}
		cmd.Env = append(os.Environ(),
			fmt.Sprintf("PKGDEST=%s", absStage),
			"PKGEXT=.pkg.tar.zst",
		)

	case "debian":
		pkgFileName = fmt.Sprintf("penguins-wardrobe_%s-%s_%s.deb", data.BaseVersion, data.Rel, getDebianArch())
		finalPath := filepath.Join(ctx.ProjRoot, pkgFileName)
		cmd = exec.Command("dpkg-deb", "--root-owner-group", "--build", stage, finalPath)

	case "fedora", "opensuse":
		pkgFileName = fmt.Sprintf("penguins-wardrobe-%s-%s.x86_64.rpm", data.BaseVersion, data.Rel)
		rpmOutDir := filepath.Join(stage, "RPMS")
		os.MkdirAll(rpmOutDir, 0755)

		specFile := filepath.Join(stage, "penguins-wardrobe.spec")

		cmd = exec.Command("rpmbuild", "-bb",
			"--define", fmt.Sprintf("_stagedir %s", stage),
			"--define", fmt.Sprintf("_rpmdir %s", rpmOutDir),
			specFile,
		)

	default:
		utils.LogWarning("Distro %s not specifically handled in packager, defaulting to debian format", dist)
		pkgFileName = fmt.Sprintf("penguins-wardrobe_%s-%s_%s.deb", data.BaseVersion, data.Rel, getDebianArch())
		finalPath := filepath.Join(ctx.ProjRoot, pkgFileName)
		cmd = exec.Command("dpkg-deb", "--root-owner-group", "--build", stage, finalPath)
	}

	utils.LogNormal("[build] Running command: %s", cmd.String())
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		utils.LogError("Packager failure: %v", err)
		return
	}

	finalDest := filepath.Join(ctx.ProjRoot, pkgFileName)

	if dist != "debian" && dist != "generic" {
		var generatedPkg string

		if dist == "arch" || dist == "manjaro" {
			matches, _ := filepath.Glob(filepath.Join(stage, "*.pkg.tar.zst"))
			if len(matches) > 0 {
				generatedPkg = matches[0]
			}
		} else if dist == "fedora" || dist == "opensuse" {
			matches, _ := filepath.Glob(filepath.Join(stage, "RPMS", "*", "*.rpm"))
			if len(matches) > 0 {
				generatedPkg = matches[0]
			}
		} else if dist == "alpine" {
			filepath.Walk(filepath.Join(stage, "APK"), func(path string, info os.FileInfo, e error) error {
				if e == nil && !info.IsDir() && strings.HasSuffix(info.Name(), ".apk") {
					generatedPkg = path
				}
				return nil
			})
		}

		if generatedPkg != "" {
			err := moveFile(generatedPkg, finalDest)
			if err != nil {
				utils.LogError("Error moving package: %v", err)
				return
			}
		} else {
			utils.LogError("Critical error: packager finished without errors, but package not found in stage!")
			return
		}
	}

	utils.LogSuccess("Package %s: %s, created in: %s", dist, pkgFileName, ctx.ProjRoot)
}

func getDebianArch() string {
	arch := runtime.GOARCH
	if arch == "386" {
		return "i386"
	}
	return arch
}
