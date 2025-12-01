package ui

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	switch m.state {
	case stateInput:
		return m.viewInput()
	case stateLoading:
		return m.viewLoading()
	case stateResults:
		return m.viewResults()
	}
	return ""
}

func (m Model) viewInput() string {
	var b strings.Builder
	b.WriteString("\n  🦆 ZUK - DuckDuckGo CLI Search\n\n")
	b.WriteString(fmt.Sprintf("  Search: %s█\n\n", m.query))
	b.WriteString("  Press Enter to search, Esc to quit\n")
	return b.String()
}

func (m Model) viewLoading() string {
	return "\n  Searching... 🔍\n"
}

func (m Model) viewResults() string {
	var b strings.Builder

	if m.err != nil {
		b.WriteString(fmt.Sprintf("\n  Error: %v\n\n", m.err))
		b.WriteString("  Press Backspace to search again, Esc to quit\n")
		return b.String()
	}

	b.WriteString("\n  🦆 Search Results\n\n")

	if len(m.results) == 0 {
		b.WriteString("  No results found.\n\n")
		b.WriteString("  Press Backspace to search again, Esc to quit\n")
		return b.String()
	}

	for i, result := range m.results {
		cursor := "  "
		if i == m.selectedIdx {
			cursor = "→ "
		}

		b.WriteString(fmt.Sprintf("%s%d. %s\n", cursor, i+1, result.Title))
		b.WriteString(fmt.Sprintf("   %s\n", result.URL))
		if result.Snippet != "" {
			b.WriteString(fmt.Sprintf("   %s\n", result.Snippet))
		}
		b.WriteString("\n")
	}

	b.WriteString("  ↑/↓ or j/k: Navigate | Enter: Open | Backspace: New search | Esc/q: Quit\n")

	return b.String()
}
