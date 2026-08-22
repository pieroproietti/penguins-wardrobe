// Copyright 2026 Piero Proietti <piero.proietti@gmail.com>.
// All rights reserved.

package context

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	EnvCI   = "ci"
	EnvVM   = "vm"
	EnvHost = "host"
)

type RuntimeContext struct {
	EnvType      string
	ProjRoot     string
	BaseBuildDir string
	StageDir     string
}

func isVirtual() bool {
	out, _ := exec.Command("systemd-detect-virt").Output()
	return strings.TrimSpace(string(out)) != "none"
}

func Detect() RuntimeContext {
	ctx := RuntimeContext{}

	if ctx.ProjRoot = os.Getenv("GITHUB_WORKSPACE"); ctx.ProjRoot == "" {
		if ctx.ProjRoot = os.Getenv("PROJ_ROOT"); ctx.ProjRoot == "" {
			cwd, _ := os.Getwd()
			ctx.ProjRoot, _ = filepath.Abs(cwd)
		}
	}

	if ctx.BaseBuildDir = os.Getenv("BUILD_DIR"); ctx.BaseBuildDir == "" {
		ctx.BaseBuildDir = "/tmp/wardrobe-build-dir"
	}

	if ctx.StageDir = os.Getenv("STAGE_DIR"); ctx.StageDir == "" {
		ctx.StageDir = "/tmp/wardrobe-stage-dir"
	}

	switch {
	case os.Getenv("GITHUB_ACTIONS") == "true" || os.Getenv("CI") == "true":
		ctx.EnvType = EnvCI
	case isVirtual():
		ctx.EnvType = EnvVM
	default:
		ctx.EnvType = EnvHost
	}

	return ctx
}
