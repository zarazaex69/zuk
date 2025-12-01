package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zarazaex69/zuk/internal/search"
)

type searchResultMsg struct {
	results []search.Result
	err     error
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.state {
		case stateInput:
			return m.updateInput(msg)
		case stateResults:
			return m.updateResults(msg)
		}

	case searchResultMsg:
		m.state = stateResults
		m.results = msg.results
		m.err = msg.err
		m.selectedIdx = 0
		return m, nil
	}

	return m, nil
}

func (m Model) updateInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "esc":
		return m, tea.Quit

	case "enter":
		if m.query != "" {
			m.state = stateLoading
			return m, m.performSearch()
		}

	case "backspace":
		if len(m.query) > 0 {
			m.query = m.query[:len(m.query)-1]
		}

	default:
		m.query += msg.String()
	}

	return m, nil
}

func (m Model) updateResults(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q", "esc":
		return m, tea.Quit

	case "up", "k":
		if m.selectedIdx > 0 {
			m.selectedIdx--
		}

	case "down", "j":
		if m.selectedIdx < len(m.results)-1 {
			m.selectedIdx++
		}

	case "enter":
		if len(m.results) > 0 {
			search.OpenBrowser(m.results[m.selectedIdx].URL)
		}

	case "backspace":
		m.state = stateInput
		m.query = ""
		m.results = nil
		m.selectedIdx = 0
	}

	return m, nil
}

func (m Model) performSearch() tea.Cmd {
	return func() tea.Msg {
		results, err := search.Search(m.query)
		return searchResultMsg{results: results, err: err}
	}
}
