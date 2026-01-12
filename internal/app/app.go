package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zarazaex69/zuk/internal/config"
	"github.com/zarazaex69/zuk/internal/ui"
)

func Run(themeName, initialQuery string) error {
	if themeName == "" {
		themeName = loadThemeFromConfig()
	}

	p := tea.NewProgram(ui.NewModel(themeName, initialQuery), tea.WithAltScreen())
	_, err := p.Run()
	return err
}

func loadThemeFromConfig() string {
	cfg, err := config.Load()
	if err != nil {
		return "default"
	}
	return cfg.Theme
}
