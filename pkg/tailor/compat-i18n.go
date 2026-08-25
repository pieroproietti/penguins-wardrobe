package tailor

import (
	"fmt"
	"os"
	"strings"
)

// systemLanguage devuelve el código de idioma (2 letras) del sistema, o "en".
func systemLanguage() string {
	for _, v := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		if loc := os.Getenv(v); len(loc) >= 2 {
			return strings.ToLower(loc[:2])
		}
	}
	if data, err := os.ReadFile("/etc/default/locale"); err == nil {
		for _, line := range strings.Split(string(data), "\n") {
			if strings.HasPrefix(line, "LANG=") {
				loc := strings.Trim(strings.TrimPrefix(line, "LANG="), "\"")
				if len(loc) >= 2 {
					return strings.ToLower(loc[:2])
				}
			}
		}
	}
	return "en"
}

// currentDistroName devuelve p.ej. "debian trixie", "ubuntu noble", "fedora".
func currentDistroName() string {
	id, codename := "", ""
	if data, err := os.ReadFile("/etc/os-release"); err == nil {
		for _, line := range strings.Split(string(data), "\n") {
			if strings.HasPrefix(line, "ID=") {
				id = strings.Trim(strings.TrimPrefix(line, "ID="), "\"")
			}
			if strings.HasPrefix(line, "VERSION_CODENAME=") {
				codename = strings.Trim(strings.TrimPrefix(line, "VERSION_CODENAME="), "\"")
			}
		}
	}
	if codename != "" {
		return id + " " + codename
	}
	return id
}

// incompatibleDistroMessage avisa, en el idioma del sistema, de que el traje
// no se ejecutará porque el sistema no es compatible.
func incompatibleDistroMessage(costume string, supported []string, current string) string {
	msgs := map[string]string{
		"es": "El traje '%s' no es compatible con este sistema (%s). Solo puede ejecutarse en: %s. El wardrobe no se ejecutará y el sistema no será modificado.",
		"gl": "O traxe '%s' non é compatible con este sistema (%s). Só pode executarse en: %s. O wardrobe non se executará e o sistema non será modificado.",
		"ru": "Костюм '%s' несовместим с этой системой (%s). Он может работать только на: %s. Wardrobe не будет запущен, и система не будет изменена.",
		"de": "Das Kostüm '%s' ist mit diesem System (%s) nicht kompatibel. Ausführung nur möglich auf: %s. Die Wardrobe wird nicht ausgeführt und das System nicht verändert.",
		"pt": "O traje '%s' não é compatível com este sistema (%s). Só pode ser executado em: %s. O wardrobe não será executado e o sistema não será modificado.",
		"it": "Il costume '%s' non è compatibile con questo sistema (%s). Può essere eseguito solo su: %s. Il wardrobe non verrà eseguito e il sistema non sarà modificato.",
		"hu": "A '%s' jelmez nem kompatibilis ezzel a rendszerrel (%s). Csak a következőkön futtatható: %s. A wardrobe nem fut, és a rendszer nem módosul.",
		"fr": "Le costume '%s' n'est pas compatible avec ce système (%s). Il ne peut s'exécuter que sur : %s. Le wardrobe ne s'exécutera pas et le système ne sera pas modifié.",
		"en": "The costume '%s' is not compatible with this system (%s). It can only run on: %s. The wardrobe will not run and the system will not be modified.",
	}
	msg, ok := msgs[systemLanguage()]
	if !ok {
		msg = msgs["en"]
	}
	return fmt.Sprintf(msg, costume, current, strings.Join(supported, ", "))
}
