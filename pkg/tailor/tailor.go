package tailor

import (
	"path/filepath"

	"github.com/pieroproietti/penguins-wardrobe/pkg/utils"
)

func Show(costumeName string) error {
	root, err := getWardrobeRoot()
	if err != nil {
		return err
	}

	costumeDir := filepath.Join(root, "costumes", costumeName)
	yamlPath := findYaml(costumeDir)
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
