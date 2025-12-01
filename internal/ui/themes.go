package ui

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	Name      string
	Logo      lipgloss.Color
	Cyan      lipgloss.Color
	Yellow    lipgloss.Color
	White     lipgloss.Color
	Gray      lipgloss.Color
	Red       lipgloss.Color
	Green     lipgloss.Color
	Blue      lipgloss.Color
	Magenta   lipgloss.Color
	BgDark    lipgloss.Color
	Separator lipgloss.Color
}

var themes = map[string]Theme{
	"default": {
		Name:      "Default",
		Logo:      lipgloss.Color("208"), // Orange
		Cyan:      lipgloss.Color("43"),
		Yellow:    lipgloss.Color("220"),
		White:     lipgloss.Color("255"),
		Gray:      lipgloss.Color("245"),
		Red:       lipgloss.Color("196"),
		Green:     lipgloss.Color("42"),
		Blue:      lipgloss.Color("39"),
		Magenta:   lipgloss.Color("201"),
		BgDark:    lipgloss.Color("236"),
		Separator: lipgloss.Color("238"),
	},
	"monochrome": {
		Name:      "Monochrome",
		Logo:      lipgloss.Color("255"), // White
		Cyan:      lipgloss.Color("255"),
		Yellow:    lipgloss.Color("255"),
		White:     lipgloss.Color("255"),
		Gray:      lipgloss.Color("245"),
		Red:       lipgloss.Color("255"),
		Green:     lipgloss.Color("255"),
		Blue:      lipgloss.Color("255"),
		Magenta:   lipgloss.Color("255"),
		BgDark:    lipgloss.Color("0"),
		Separator: lipgloss.Color("238"),
	},
	"black": {
		Name:      "Black",
		Logo:      lipgloss.Color("15"), // Bright white
		Cyan:      lipgloss.Color("14"), // Bright cyan
		Yellow:    lipgloss.Color("11"), // Bright yellow
		White:     lipgloss.Color("15"), // Bright white
		Gray:      lipgloss.Color("8"),  // Dark gray
		Red:       lipgloss.Color("9"),  // Bright red
		Green:     lipgloss.Color("10"), // Bright green
		Blue:      lipgloss.Color("12"), // Bright blue
		Magenta:   lipgloss.Color("13"), // Bright magenta
		BgDark:    lipgloss.Color("0"),  // Black
		Separator: lipgloss.Color("8"),
	},
	"soft": {
		Name:      "Soft",
		Logo:      lipgloss.Color("173"), // Soft orange
		Cyan:      lipgloss.Color("116"), // Soft cyan
		Yellow:    lipgloss.Color("186"), // Soft yellow
		White:     lipgloss.Color("252"), // Soft white
		Gray:      lipgloss.Color("243"), // Soft gray
		Red:       lipgloss.Color("167"), // Soft red
		Green:     lipgloss.Color("108"), // Soft green
		Blue:      lipgloss.Color("110"), // Soft blue
		Magenta:   lipgloss.Color("175"), // Soft magenta
		BgDark:    lipgloss.Color("235"),
		Separator: lipgloss.Color("240"),
	},
	"nya": {
		Name:      "Nya",
		Logo:      lipgloss.Color("213"), // Pink
		Cyan:      lipgloss.Color("117"), // Light blue
		Yellow:    lipgloss.Color("229"), // Cream
		White:     lipgloss.Color("255"), // White
		Gray:      lipgloss.Color("246"), // Light gray
		Red:       lipgloss.Color("210"), // Light pink/red
		Green:     lipgloss.Color("121"), // Mint green
		Blue:      lipgloss.Color("153"), // Lavender blue
		Magenta:   lipgloss.Color("219"), // Light magenta
		BgDark:    lipgloss.Color("234"), // Very dark gray
		Separator: lipgloss.Color("239"), // Medium gray
	},
}

func GetTheme(name string) Theme {
	if theme, ok := themes[name]; ok {
		return theme
	}
	return themes["default"]
}

func GetThemeNames() []string {
	return []string{"default", "monochrome", "black", "soft", "nya"}
}
