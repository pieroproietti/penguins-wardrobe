// Copyright 2026 Piero Proietti <piero.proietti@gmail.com>.
// All rights reserved.

package builder

import (
	"fmt"
	"io"
	"os"
)

// moveFile moves a file safely even across different filesystems (e.g. tmpfs -> ext4)
func moveFile(sourcePath, destPath string) error {
	// 1. Try native rename
	err := os.Rename(sourcePath, destPath)
	if err == nil {
		return nil
	}

	// 2. Fall back to copy and delete
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("unable to open source: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("unable to create destination: %v", err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("error during copy: %v", err)
	}

	inputFile.Close()

	return os.Remove(sourcePath)
}
