package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wardrobe",
	Short: "penguins-wardrobe: dress up your Linux distribution with costumes and configurations",
	Long: `penguins-wardrobe is a lightweight tool to manage and apply system configurations,
desktop environments, and themes ("costumes") to Linux distributions.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(getCmd())
	rootCmd.AddCommand(listCmd())
	rootCmd.AddCommand(showCmd())
	rootCmd.AddCommand(wearCmd())
	rootCmd.AddCommand(versionCmd())
}
