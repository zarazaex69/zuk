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

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		headerHeight := 4
		footerHeight := 2
		verticalMargins := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMargins)
			m.viewport.YPosition = headerHeight
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMargins
		}

		if m.state == stateResults {
			m.viewport.SetContent(m.renderResultsList())
		}

		return m, nil

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
		m.viewport.SetContent(m.renderResultsList())
		m.viewport.GotoTop()
		return m, nil
	}

	return m, cmd
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
			runes := []rune(m.query)
			m.query = string(runes[:len(runes)-1])
		}

	default:
		if len(msg.String()) == 1 || msg.Type == tea.KeySpace {
			m.query += msg.String()
		}
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
			m.viewport.SetContent(m.renderResultsList())
			m.ensureSelectedVisible()
		}

	case "down", "j":
		if m.selectedIdx < len(m.results)-1 {
			m.selectedIdx++
			m.viewport.SetContent(m.renderResultsList())
			m.ensureSelectedVisible()
		}

	case "enter":
		if len(m.results) > 0 && m.selectedIdx < len(m.results) {
			search.OpenBrowser(m.results[m.selectedIdx].URL)
		}

	case "backspace":
		m.state = stateInput
		m.query = ""
		m.results = nil
		m.selectedIdx = 0
		m.viewport.GotoTop()
	}

	return m, nil
}

func (m *Model) ensureSelectedVisible() {
	// Each result takes approximately 4 lines
	linesPerResult := 4
	selectedLine := m.selectedIdx * linesPerResult

	if selectedLine < m.viewport.YOffset {
		m.viewport.SetYOffset(selectedLine)
	} else if selectedLine >= m.viewport.YOffset+m.viewport.Height {
		m.viewport.SetYOffset(selectedLine - m.viewport.Height + linesPerResult)
	}
}

func (m Model) performSearch() tea.Cmd {
	return func() tea.Msg {
		results, err := search.Search(m.query)
		return searchResultMsg{results: results, err: err}
	}
}
