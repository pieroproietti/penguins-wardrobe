// Copyright 2026 Piero Proietti <piero.proietti@gmail.com>.
// All rights reserved.

package builder

import (
	"fmt"
	"os"
	"text/template"
)

func writeTemplate(tmplPath string, destPath string, data RecipeData) error {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return fmt.Errorf("unable to read template %s: %w", tmplPath, err)
	}

	f, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("unable to create destination file %s: %w", destPath, err)
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		return fmt.Errorf("error generating template %s: %w", tmplPath, err)
	}

	err = f.Sync()
	if err != nil {
		return fmt.Errorf("error syncing file %s: %w", destPath, err)
	}

	return nil
}
