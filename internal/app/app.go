package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zarazaex69/zuk/internal/config"
	"github.com/zarazaex69/zuk/internal/ui"
)

func Run(themeName string, initialQuery string) error {
	// Load theme from config if not specified
	if themeName == "" {
		cfg, err := config.Load()
		if err == nil {
			themeName = cfg.Theme
		} else {
			themeName = "default"
		}
	}

	p := tea.NewProgram(ui.NewModel(themeName, initialQuery), tea.WithAltScreen())
	_, err := p.Run()
	return err
}
