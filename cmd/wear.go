package cmd

import (
	"github.com/pieroproietti/penguins-wardrobe/pkg/tailor"
	"github.com/spf13/cobra"
)

func wearCmd() *cobra.Command {
	var noAcc bool
	var noFirm bool

	cmd := &cobra.Command{
		Use:   "wear [costume]",
		Short: "Wear a costume from the wardrobe",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return tailor.Wear(args[0], noAcc, noFirm)
		},
	}

	cmd.Flags().BoolVar(&noAcc, "no-acc", false, "Do not install accessories")
	cmd.Flags().BoolVar(&noFirm, "no-firm", false, "Do not install firmware")

	return cmd
}
