package cmd

import (
	"github.com/pieroproietti/penguins-wardrobe/pkg/tailor"
	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List available costumes",
		RunE: func(cmd *cobra.Command, args []string) error {
			return tailor.List()
		},
	}
}
