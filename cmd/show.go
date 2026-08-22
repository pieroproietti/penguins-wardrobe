package cmd

import (
	"github.com/pieroproietti/penguins-wardrobe/pkg/tailor"
	"github.com/spf13/cobra"
)

func showCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show [costume]",
		Short: "Show details of a costume",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return tailor.Show(args[0])
		},
	}
}
