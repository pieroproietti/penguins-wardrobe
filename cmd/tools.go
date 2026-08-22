// Copyright 2026 Piero Proietti <piero.proietti@gmail.com>.
// All rights reserved.

package cmd

import (
	"github.com/spf13/cobra"
)

var toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "Useful tools for maintenance and system management",
	Long: `A suite of auxiliary tools provided by wardrobe for the management,
packaging, and inspection of the system.`,
}

func init() {
	rootCmd.AddCommand(toolsCmd)
}
