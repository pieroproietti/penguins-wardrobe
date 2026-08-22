// Copyright 2026 Piero Proietti <piero.proietti@gmail.com>.
// All rights reserved.

package builder

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pieroproietti/penguins-wardrobe/pkg/context"
	"github.com/pieroproietti/penguins-wardrobe/pkg/utils"
)

// recipe writes the control file (PKGBUILD, SPEC, etc.) into the staging area
func recipe(ctx context.RuntimeContext, dist string, data RecipeData) {
	utils.LogNormal("[build] Recipe: writing recipe for %s...", dist)

	stage := ctx.StageDir

	switch dist {
	case "alpine":
		writeAPKBUILD(ctx, stage, data)

	case "arch", "manjaro":
		writePKGBUILD(ctx, stage, dist, data)

	case "fedora", "opensuse":
		writeSpecFile(ctx, stage, dist, data)

	default:
		writeDebianFiles(ctx, stage, data)
	}
}

// writeAPKBUILD writes the APKBUILD for alpine
func writeAPKBUILD(ctx context.RuntimeContext, stage string, data RecipeData) error {
	tmplPath := filepath.Join(ctx.ProjRoot, "pkg/builder/templates/alpine.tmpl")
	destPath := filepath.Join(stage, "APKBUILD")
	return writeTemplate(tmplPath, destPath, data)
}

// writePKGBUILD writes the PKGBUILD for arch/manjaro
func writePKGBUILD(ctx context.RuntimeContext, stage string, dist string, data RecipeData) error {
	tmplName := fmt.Sprintf("%s.tmpl", dist)
	tmplPath := filepath.Join(ctx.ProjRoot, "pkg/builder/templates", tmplName)
	destPath := filepath.Join(stage, "PKGBUILD")
	return writeTemplate(tmplPath, destPath, data)
}

// writeSpecFile writes penguins-wardrobe.spec for fedora/opensuse
func writeSpecFile(ctx context.RuntimeContext, stage string, dist string, data RecipeData) error {
	tmplName := fmt.Sprintf("%s.tmpl", dist)
	tmplPath := filepath.Join(ctx.ProjRoot, "pkg/builder/templates", tmplName)
	destPath := filepath.Join(stage, "penguins-wardrobe.spec")
	return writeTemplate(tmplPath, destPath, data)
}

// writeDebianFiles sets up the DEBIAN directory and creates control, rules, compat, copyright, changelog
func writeDebianFiles(ctx context.RuntimeContext, stage string, data RecipeData) error {
	debianDir := filepath.Join(stage, "DEBIAN")
	os.MkdirAll(debianDir, 0755)

	files := map[string]string{
		"control.tmpl":   "control",
		"rules.tmpl":     "rules",
		"compat.tmpl":    "compat",
		"copyright.tmpl": "copyright",
		"changelog.tmpl": "changelog",
	}

	for tmplName, destName := range files {
		tmplPath := filepath.Join(ctx.ProjRoot, "pkg/builder/templates/debian", tmplName)
		destPath := filepath.Join(debianDir, destName)

		utils.LogNormal("--> Writing template: %s -> %s", tmplName, destPath)

		err := writeTemplate(tmplPath, destPath, data)
		if err != nil {
			return err
		}
	}
	return nil
}
