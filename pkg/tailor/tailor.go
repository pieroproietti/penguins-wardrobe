package tailor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pieroproietti/penguins-wardrobe/pkg/utils"
)

func Show(costumeName string) error {
	v2Dir, err := getWardrobeV2Dir()
	if err != nil {
		return err
	}

	costumeDir := filepath.Join(v2Dir, "costumes", costumeName)
	if _, err := os.Stat(costumeDir); os.IsNotExist(err) {
		if strings.HasPrefix(costumeName, "accessories/") || strings.HasPrefix(costumeName, "costumes/") {
			costumeDir = filepath.Join(v2Dir, costumeName)
		} else {
			accDir := filepath.Join(v2Dir, "accessories", costumeName)
			if _, errAcc := os.Stat(accDir); errAcc == nil {
				costumeDir = accDir
			}
		}
	}

	yamlPath := findYaml(costumeDir)
	if yamlPath == "" {
		return fmt.Errorf("costume '%s' not found in %s", costumeName, v2Dir)
	}

	suit, err := loadSuit(yamlPath)
	if err != nil {
		return err
	}

	isAcc := strings.Contains(costumeDir, "/accessories/")
	icon := "👗"
	titleType := "COSTUME"
	if isAcc {
		icon = "👝"
		titleType = "ACCESSORY"
	}

	versionStr := ""
	if suit.Release != "" {
		versionStr = fmt.Sprintf(" (v%s)", suit.Release)
	}

	utils.PrintBanner(icon, fmt.Sprintf("%s: %s%s", titleType, suit.Name, versionStr), suit.Description)
	if suit.Author != "" {
		fmt.Printf("  %-16s: %s\n", "Autore", suit.Author)
	}
	if len(suit.Distributions) > 0 {
		fmt.Printf("  %-16s: %s\n", "Distribuzioni", strings.Join(suit.Distributions, ", "))
	}
	if len(suit.Accessories) > 0 {
		fmt.Printf("  %-16s: %s\n", "Accessori", strings.Join(suit.Accessories, ", "))
	}
	if len(suit.Packages) > 0 {
		limit := 5
		if len(suit.Packages) < limit {
			limit = len(suit.Packages)
		}
		preview := strings.Join(suit.Packages[:limit], ", ")
		if len(suit.Packages) > limit {
			preview += "..."
		}
		fmt.Printf("  %-16s: %d pacchetti (%s)\n", "Pacchetti", len(suit.Packages), preview)
	}
	if len(suit.Cmds) > 0 {
		fmt.Printf("  %-16s: %d comandi di finalizzazione\n", "Comandi", len(suit.Cmds))
	}
	fmt.Println()
	return nil
}
