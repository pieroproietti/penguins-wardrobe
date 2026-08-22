package tailor

import (
	"fmt"
	"os"

	"github.com/pieroproietti/penguins-wardrobe/pkg/utils"
)

func Get() error {
	root, err := getWardrobeRoot()
	if err != nil {
		return err
	}

	if _, err := os.Stat(root); os.IsNotExist(err) {
		utils.LogNormal("Downloading costumes repository to %s...", root)
		cmd := fmt.Sprintf("git clone https://github.com/pieroproietti/penguins-wardrobe.git %s", root)
		return utils.Exec(cmd)
	}

	utils.LogNormal("Wardrobe already present in %s. To update, use 'git -C %s pull'.", root, root)
	return nil
}
