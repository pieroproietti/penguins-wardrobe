// Copyright 2026 Piero Proietti <piero.proietti@gmail.com>.
// All rights reserved.

package utils

import (
	"fmt"
	"os"
)

// Colori e stili ANSI
const (
	ColorBold    = "\033[1m"
	ColorDim     = "\033[2m"
	ColorBlue    = "\033[1;34m"
	ColorCyan    = "\033[36m"
	ColorGreen   = "\033[1;32m"
	ColorRed     = "\033[1;31m"
	ColorReset   = "\033[0m"
	ColorYellow  = "\033[33m"
	ColorMagenta = "\033[35m"
	ColorWhite   = "\033[1;37m"
)

// DisableColors permette di disattivare i colori.
var DisableColors bool

func init() {
	// Auto-rilevamento: se os.Stdout NON è un terminale (es. rediretto in un file log o pipe),
	// spegne i colori in automatico per evitare caratteri ANSI sporchi nel testo.
	stat, _ := os.Stdout.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		DisableColors = true
	}
}

// colorize restituisce il codice ANSI solo se i colori sono abilitati,
// altrimenti restituisce una stringa vuota, mantenendo il log pulito.
func colorize(colorCode string) string {
	if DisableColors {
		return ""
	}
	return colorCode
}

// --- SISTEMA DI LOGGING CENTRALIZZATO ---

// LogNormal stampa un messaggio informativo con il tag [wardrobe]
func LogNormal(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Printf("%s[wardrobe]%s %s\n", colorize(ColorCyan), colorize(ColorReset), msg)
}

// LogSuccess stampa un messaggio di successo
func LogSuccess(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Printf("%s[wardrobe]%s %s\n", colorize(ColorGreen), colorize(ColorReset), msg)
}

// LogWarning stampa un messaggio di avviso
func LogWarning(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Printf("%s[WARNING]%s %s\n", colorize(ColorYellow), colorize(ColorReset), msg)
}

// LogError stampa un messaggio di errore sullo STANDARD ERROR
func LogError(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stderr, "\n%s[ERROR]%s %s\n", colorize(ColorRed), colorize(ColorReset), msg)
}

// Fatal stampa un errore ed esce con codice 1
func Fatal(format string, a ...interface{}) {
	LogError(format, a...)
	os.Exit(1)
}

// --- HELPER VISUALI PER FORMATTAZIONE E SEZIONI ---

const sectionDivider = "============================================================"

// PrintBanner stampa un'intestazione principale incorniciata
func PrintBanner(icon, title, subtitle string) {
	fmt.Println()
	fmt.Printf("%s%s%s\n", colorize(ColorCyan), sectionDivider, colorize(ColorReset))
	fmt.Printf("  %s %s%s%s\n", icon, colorize(ColorBold+ColorWhite), title, colorize(ColorReset))
	if subtitle != "" {
		fmt.Printf("  %s%s%s\n", colorize(ColorDim), subtitle, colorize(ColorReset))
	}
	fmt.Printf("%s%s%s\n", colorize(ColorCyan), sectionDivider, colorize(ColorReset))
}

// PrintSection stampa un separatore di sezione principale
func PrintSection(icon, title string) {
	fmt.Println()
	fmt.Printf("%s%s%s\n", colorize(ColorCyan), sectionDivider, colorize(ColorReset))
	fmt.Printf("  %s %s%s%s\n", icon, colorize(ColorBold+ColorWhite), title, colorize(ColorReset))
	fmt.Printf("%s%s%s\n", colorize(ColorCyan), sectionDivider, colorize(ColorReset))
}

// PrintSubSection stampa un'intestazione di sotto-sezione (es. per singoli accessori)
func PrintSubSection(icon, title string) {
	fmt.Println()
	fmt.Printf("%s%s%s %s%s%s\n", colorize(ColorCyan), icon, colorize(ColorReset), colorize(ColorBold), title, colorize(ColorReset))
}

// PrintSummaryBox stampa il riepilogo finale formattato
func PrintSummaryBox(title string, rows [][2]string) {
	fmt.Println()
	fmt.Printf("%s%s%s\n", colorize(ColorGreen), sectionDivider, colorize(ColorReset))
	fmt.Printf("  %s%s%s\n", colorize(ColorBold+ColorGreen), title, colorize(ColorReset))
	fmt.Printf("%s%s%s\n", colorize(ColorGreen), sectionDivider, colorize(ColorReset))
	for _, row := range rows {
		fmt.Printf("  %-20s: %s%s%s\n", row[0], colorize(ColorWhite), row[1], colorize(ColorReset))
	}
	fmt.Printf("%s%s%s\n", colorize(ColorGreen), sectionDivider, colorize(ColorReset))
}
