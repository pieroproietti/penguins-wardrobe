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

	utils.LogNormal(utils.ColorCyan+"Costume: %s"+utils.ColorReset, suit.Name)
	utils.LogNormal("Descrizione: %s", suit.Description)
	if len(suit.Distributions) > 0 {
		utils.LogNormal("Distribuzioni: %v", suit.Distributions)
	}
	utils.LogNormal("Pacchetti: %v", suit.Packages)
	if len(suit.Accessories) > 0 {
		utils.LogNormal("Accessori: %v", suit.Accessories)
	}
	if len(suit.Cmds) > 0 {
		utils.LogNormal("Comandi: %v", suit.Cmds)
	}
	return nil
}
