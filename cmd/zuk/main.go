package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/zarazaex69/zuk/internal/app"
	"github.com/zarazaex69/zuk/internal/config"
	"github.com/zarazaex69/zuk/internal/ui"
)

func main() {
	var themeName string
	flag.StringVar(&themeName, "theme", "", "Set theme (default, monochrome, black, soft, nya)")
	flag.StringVar(&themeName, "t", "", "Set theme (short)")
	listThemes := flag.Bool("list-themes", false, "List available themes")
	flag.Parse()

	if *listThemes {
		printThemes()
		return
	}

	if themeName != "" {
		saveTheme(themeName)
	}

	query := strings.Join(flag.Args(), " ")

	if err := app.Run(themeName, query); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printThemes() {
	fmt.Println("Available themes:")
	for _, name := range ui.GetThemeNames() {
		theme := ui.GetTheme(name)
		fmt.Printf("  - %s: %s\n", name, theme.Name)
	}
}

func saveTheme(name string) {
	cfg := &config.Config{Theme: name}
	if err := cfg.Save(); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not save theme: %v\n", err)
	}
}
