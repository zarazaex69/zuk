package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zarazaex69/zuk/internal/ui"
)

func Run() error {
	p := tea.NewProgram(ui.NewModel(), tea.WithAltScreen())
	_, err := p.Run()
	return err
}
