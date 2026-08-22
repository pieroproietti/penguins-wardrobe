package cmd

import (
	"github.com/pieroproietti/penguins-wardrobe/pkg/tailor"
	"github.com/spf13/cobra"
)

func getCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Download or update the wardrobe costumes repository",
		RunE: func(cmd *cobra.Command, args []string) error {
			return tailor.Get()
		},
	}
}
