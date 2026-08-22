// Copyright 2026 Piero Proietti <piero.proietti@gmail.com>.
// All rights reserved.

package cmd

import (
	"os"
	"path/filepath"

	"github.com/pieroproietti/penguins-wardrobe/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var docTarget string

var genDocsCmd = &cobra.Command{
	Use:    "_gen_docs",
	Hidden: true,
	Short:  "Generate Markdown files, Man pages, and autocompletion scripts",
	Run: func(cmd *cobra.Command, args []string) {
		utils.LogNormal("[gen_docs] Generating documentation in folder: %s", docTarget)
		mdDir := filepath.Join(docTarget, "md")
		manDir := filepath.Join(docTarget, "man")
		compDir := filepath.Join(docTarget, "completion")

		dirs := []string{mdDir, manDir, compDir}
		for _, dir := range dirs {
			if err := os.MkdirAll(dir, 0755); err != nil {
				utils.LogError("Unable to create directory %s: %v", dir, err)
				os.Exit(1)
			}
		}

		if err := doc.GenMarkdownTree(rootCmd, mdDir); err != nil {
			utils.LogError("Markdown generation failed: %v", err)
			os.Exit(1)
		}

		header := &doc.GenManHeader{
			Title:   "WARDROBE",
			Section: "1",
		}
		if err := doc.GenManTree(rootCmd, header, manDir); err != nil {
			utils.LogError("Man pages generation failed: %v", err)
			os.Exit(1)
		}

		rootCmd.GenBashCompletionFile(filepath.Join(compDir, "wardrobe.bash"))
		rootCmd.GenZshCompletionFile(filepath.Join(compDir, "wardrobe.zsh"))
		rootCmd.GenFishCompletionFile(filepath.Join(compDir, "wardrobe.fish"), true)

		utils.LogSuccess("[gen_docs] Documentation and completions generated successfully.")
	},
}

func init() {
	genDocsCmd.Flags().StringVar(&docTarget, "target", "./docs", "Target directory for generation")
	rootCmd.AddCommand(genDocsCmd)
}
