package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zarazaex69/zuk/internal/app"
	"github.com/zarazaex69/zuk/internal/config"
	"github.com/zarazaex69/zuk/internal/ui"
)

func main() {
	themeFlag := flag.String("theme", "", "Set theme (default, monochrome, black, soft, nya)")
	themeFlagShort := flag.String("t", "", "Set theme (short)")
	listThemes := flag.Bool("list-themes", false, "List available themes")
	flag.Parse()

	// Handle theme listing
	if *listThemes {
		fmt.Println("Available themes:")
		for _, name := range ui.GetThemeNames() {
			theme := ui.GetTheme(name)
			fmt.Printf("  - %s: %s\n", name, theme.Name)
		}
		return
	}

	// Determine theme
	themeName := *themeFlag
	if themeName == "" {
		themeName = *themeFlagShort
	}

	// Save theme if specified
	if themeName != "" {
		cfg := &config.Config{Theme: themeName}
		if err := cfg.Save(); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Could not save theme: %v\n", err)
		}
	}

	// Get search query from remaining arguments
	var initialQuery string
	if len(flag.Args()) > 0 {
		initialQuery = flag.Args()[0]
	}

	// Run app
	if err := app.Run(themeName, initialQuery); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
