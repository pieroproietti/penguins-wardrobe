// Copyright 2026 Piero Proietti <piero.proietti@gmail.com>.
// All rights reserved.

package cmd

import (
	"github.com/spf13/cobra"
)

const (
	remoteUserHost = "root@192.168.1.2"
	remotePkgPath  = "/eggs/"
)

var cleanExport bool

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export artifacts (pkg, log) to remote storage",
	Long: `The export command handles the transfer of produced artifacts
(distribution packages or log files) to configured remote servers,
automating cleanup and versioning.`,
}

func init() {
	exportCmd.PersistentFlags().BoolVar(&cleanExport, "clean", false, "Clean old versions on remote server before exporting")
	rootCmd.AddCommand(exportCmd)
}
