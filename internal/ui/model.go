package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zarazaex69/zuk/internal/search"
)

type state int

const (
	stateInput state = iota
	stateLoading
	stateResults
)

type Model struct {
	state       state
	query       string
	results     []search.Result
	selectedIdx int
	err         error
	width       int
	height      int
	viewport    viewport.Model
	ready       bool
	theme       Theme
}

func NewModel(themeName string) Model {
	return Model{
		state:  stateInput,
		query:  "",
		width:  80,
		height: 24,
		theme:  GetTheme(themeName),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
