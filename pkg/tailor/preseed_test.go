package tailor

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindPreseed(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "wardrobe-preseed-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// No preseed file
	if found := findPreseed(tempDir); found != "" {
		t.Errorf("expected empty string, got %s", found)
	}

	// Should ignore other names like index.preseed or preseed.cfg
	indexPreseed := filepath.Join(tempDir, "index.preseed")
	if err := os.WriteFile(indexPreseed, []byte("lightdm shared/default-x-display-manager select lightdm\n"), 0644); err != nil {
		t.Fatalf("failed to write index.preseed: %v", err)
	}
	if found := findPreseed(tempDir); found != "" {
		t.Errorf("expected empty string for index.preseed, got %s", found)
	}

	// Must match packages.preseed exclusively
	packagesPreseed := filepath.Join(tempDir, "packages.preseed")
	if err := os.WriteFile(packagesPreseed, []byte("lightdm shared/default-x-display-manager select lightdm\n"), 0644); err != nil {
		t.Fatalf("failed to write packages.preseed: %v", err)
	}

	if found := findPreseed(tempDir); found != packagesPreseed {
		t.Errorf("expected %s, got %s", packagesPreseed, found)
	}
}
