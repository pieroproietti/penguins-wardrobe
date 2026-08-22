// Copyright 2026 Piero Proietti <piero.proietti@gmail.com>.
// All rights reserved.

package builder

import (
	"fmt"
	"strings"
	"time"

	"github.com/pieroproietti/penguins-wardrobe/pkg/context"
	"github.com/pieroproietti/penguins-wardrobe/pkg/distro"
	"github.com/pieroproietti/penguins-wardrobe/pkg/utils"
)

func LogBuild(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	utils.LogNormal("[build] %s", msg)
}

func getDebianDepends(arch string) string {
	return "git, rsync"
}

func HandleBuild(d *distro.Distro) {
	// 1. Data preparation
	ctx := context.Detect()
	baseVer, relNum := getGitVersion()
	dist := strings.ToLower(d.DistroLike)
	now := time.Now()
	arch := getDebianArch()
	data := RecipeData{
		BaseVersion: baseVer,
		Rel:         relNum,
		Date:        now.Format(time.RFC1123Z),
		RpmDate:     now.Format("Mon Jan 02 2006"),
		Arch:        arch,
		Depends:     getDebianDepends(arch),
	}

	// 2. staging
	staging(ctx)

	// 3. addBuildRecipe
	recipe(ctx, dist, data)

	// 4. Packager
	packager(ctx, dist, data)
}
