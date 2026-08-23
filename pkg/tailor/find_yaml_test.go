package tailor

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindYamlAndLoadSuit(t *testing.T) {
	v2Dir, err := getWardrobeV2Dir()
	if err != nil {
		t.Fatalf("getWardrobeV2Dir error: %v", err)
	}

	if _, err := os.Stat(v2Dir); os.IsNotExist(err) {
		// Fallback to local repo v2 for test suite execution
		if localV2, err := filepath.Abs("../../v2"); err == nil {
			if _, err := os.Stat(localV2); err == nil {
				v2Dir = localV2
			}
		}
	}

	// Test costume colibri
	colibriDir := filepath.Join(v2Dir, "costumes", "colibri")
	colibriYaml := findYaml(colibriDir)
	if colibriYaml == "" {
		t.Fatalf("findYaml failed for colibri at %s", colibriDir)
	}
	colibriSuit, err := loadSuit(colibriYaml)
	if err != nil {
		t.Fatalf("loadSuit failed for colibri: %v", err)
	}
	if len(colibriSuit.Accessories) != 2 {
		t.Errorf("expected 2 accessories, got %d: %v", len(colibriSuit.Accessories), colibriSuit.Accessories)
	}

	// Test accessory eggs-dev
	eggsDevDir := filepath.Join(v2Dir, "accessories", "eggs-dev")
	eggsDevYaml := findYaml(eggsDevDir)
	if eggsDevYaml == "" {
		t.Fatalf("findYaml failed for eggs-dev at %s", eggsDevDir)
	}
	eggsDevSuit, err := loadSuit(eggsDevYaml)
	if err != nil {
		t.Fatalf("loadSuit failed for eggs-dev: %v", err)
	}
	if eggsDevSuit.Name != "eggs-dev" {
		t.Errorf("expected name 'eggs-dev', got '%s'", eggsDevSuit.Name)
	}
	if len(eggsDevSuit.Packages) == 0 {
		t.Errorf("expected packages in eggs-dev, got none")
	}

	// Test accessory base
	baseDir := filepath.Join(v2Dir, "accessories", "base")
	baseYaml := findYaml(baseDir)
	if baseYaml == "" {
		t.Fatalf("findYaml failed for base at %s", baseDir)
	}
	baseSuit, err := loadSuit(baseYaml)
	if err != nil {
		t.Fatalf("loadSuit failed for base: %v", err)
	}
	if baseSuit.Name != "base" {
		t.Errorf("expected name 'base', got '%s'", baseSuit.Name)
	}
}
