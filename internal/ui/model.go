package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zarazaex69/zuk/internal/search"
)

type searchResultMsg struct {
	results []search.Result
	err     error
}

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

func NewModel(themeName string, initialQuery string) Model {
	m := Model{
		state:  stateInput,
		query:  initialQuery,
		width:  80,
		height: 24,
		theme:  GetTheme(themeName),
	}

	// If initial query provided, start in loading state
	if initialQuery != "" {
		m.state = stateLoading
	}

	return m
}

func (m Model) Init() tea.Cmd {
	// If we have an initial query, start searching immediately
	if m.query != "" && m.state == stateLoading {
		return m.performSearch()
	}
	return nil
}

func (m Model) performSearch() tea.Cmd {
	return func() tea.Msg {
		results, err := search.Search(m.query)
		return searchResultMsg{results: results, err: err}
	}
}
