// Copyright 2026 Piero Proietti <piero.proietti@gmail.com>.
// All rights reserved.

package main

import (
	"fmt"
	"os"

	"github.com/pieroproietti/penguins-wardrobe/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
