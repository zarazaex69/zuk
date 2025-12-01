package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}

	switch m.state {
	case stateInput:
		return m.viewInput()
	case stateLoading:
		return m.viewLoading()
	case stateResults:
		return m.viewResultsWithViewport()
	}
	return ""
}

func (m Model) viewInput() string {
	var b strings.Builder

	// Styles based on theme
	logoStyle := lipgloss.NewStyle().Foreground(m.theme.Logo).Bold(true)
	subtitleStyle := lipgloss.NewStyle().Foreground(m.theme.Gray)
	separatorStyle := lipgloss.NewStyle().Foreground(m.theme.Separator)
	searchLabelStyle := lipgloss.NewStyle().Foreground(m.theme.Cyan).Bold(true)
	queryStyle := lipgloss.NewStyle().Foreground(m.theme.White)
	cursorStyle := lipgloss.NewStyle().Foreground(m.theme.Yellow).Bold(true)
	helpBarStyle := lipgloss.NewStyle().Foreground(m.theme.Gray)
	helpKeyStyle := lipgloss.NewStyle().Foreground(m.theme.Yellow).Bold(true)

	// Header
	logo := logoStyle.Render("🦆 ZUK")
	subtitle := subtitleStyle.Render(" - DuckDuckGo CLI Search")
	b.WriteString("\n  " + logo + subtitle + "\n\n")

	// Separator
	sep := separatorStyle.Render(strings.Repeat("─", m.width-4))
	b.WriteString("  " + sep + "\n\n")

	// Search input
	label := searchLabelStyle.Render("  Search: ")
	query := queryStyle.Render(m.query)
	cursor := cursorStyle.Render("█")
	b.WriteString(label + query + cursor + "\n\n")

	// Separator
	b.WriteString("  " + sep + "\n\n")

	// Help
	help := helpBarStyle.Render("  Press ") +
		helpKeyStyle.Render("Enter") +
		helpBarStyle.Render(" to search, ") +
		helpKeyStyle.Render("Esc") +
		helpBarStyle.Render(" to quit")
	b.WriteString(help + "\n")

	return b.String()
}

func (m Model) viewLoading() string {
	var b strings.Builder

	logoStyle := lipgloss.NewStyle().Foreground(m.theme.Logo).Bold(true)
	separatorStyle := lipgloss.NewStyle().Foreground(m.theme.Separator)
	loadingStyle := lipgloss.NewStyle().Foreground(m.theme.Cyan).Bold(true)
	queryStyle := lipgloss.NewStyle().Foreground(m.theme.White)

	logo := logoStyle.Render("🦆 ZUK")
	b.WriteString("\n  " + logo + "\n\n")

	sep := separatorStyle.Render(strings.Repeat("─", m.width-4))
	b.WriteString("  " + sep + "\n\n")

	loading := loadingStyle.Render("  🔍 Searching for: ") +
		queryStyle.Render(m.query) +
		loadingStyle.Render("...")
	b.WriteString(loading + "\n")

	return b.String()
}

func (m Model) viewResultsWithViewport() string {
	var b strings.Builder

	logoStyle := lipgloss.NewStyle().Foreground(m.theme.Logo).Bold(true)
	countStyle := lipgloss.NewStyle().Foreground(m.theme.Green)
	separatorStyle := lipgloss.NewStyle().Foreground(m.theme.Separator)
	scrollStyle := lipgloss.NewStyle().Foreground(m.theme.Gray).Italic(true)
	helpBarStyle := lipgloss.NewStyle().Foreground(m.theme.Gray)
	helpKeyStyle := lipgloss.NewStyle().Foreground(m.theme.Yellow).Bold(true)

	// Header
	logo := logoStyle.Render("🦆 ZUK")
	resultCount := countStyle.Render(fmt.Sprintf(" (%d results)", len(m.results)))
	b.WriteString("\n  " + logo + resultCount + "\n")

	// Separator
	sep := separatorStyle.Render(strings.Repeat("─", m.width-4))
	b.WriteString("  " + sep + "\n")

	// Viewport content
	b.WriteString(m.viewport.View())
	b.WriteString("\n")

	// Footer separator
	b.WriteString("  " + sep + "\n")

	// Scroll info
	scrollPercent := scrollStyle.Render(fmt.Sprintf("  [%d/%d]", m.selectedIdx+1, len(m.results)))
	if m.viewport.TotalLineCount() > m.viewport.Height {
		scrollPercent += scrollStyle.Render(fmt.Sprintf(" %.0f%%", m.viewport.ScrollPercent()*100))
	}
	b.WriteString(scrollPercent + "\n")

	// Help bar
	help := helpBarStyle.Render("  ") +
		helpKeyStyle.Render("↑/↓") + helpBarStyle.Render(" navigate  ") +
		helpKeyStyle.Render("Enter") + helpBarStyle.Render(" open  ") +
		helpKeyStyle.Render("Backspace") + helpBarStyle.Render(" new search  ") +
		helpKeyStyle.Render("q") + helpBarStyle.Render(" quit")
	b.WriteString(help)

	return b.String()
}

func (m Model) renderResultsList() string {
	errorStyle := lipgloss.NewStyle().Foreground(m.theme.Red).Bold(true)
	grayStyle := lipgloss.NewStyle().Foreground(m.theme.Gray)
	markerStyle := lipgloss.NewStyle().Foreground(m.theme.Yellow).Bold(true)
	numberStyle := lipgloss.NewStyle().Foreground(m.theme.Gray)
	titleStyle := lipgloss.NewStyle().Foreground(m.theme.White).Bold(true)
	titleSelectedStyle := lipgloss.NewStyle().Foreground(m.theme.Yellow).Bold(true)
	urlStyle := lipgloss.NewStyle().Foreground(m.theme.Cyan)
	snippetStyle := lipgloss.NewStyle().Foreground(m.theme.Gray)

	if m.err != nil {
		return errorStyle.Render("  Error: " + m.err.Error())
	}

	if len(m.results) == 0 {
		return grayStyle.Render("  No results found.")
	}

	var b strings.Builder
	maxWidth := m.width - 8
	if maxWidth < 40 {
		maxWidth = 40
	}

	for i, result := range m.results {
		isSelected := i == m.selectedIdx

		// Selection marker
		var marker string
		if isSelected {
			marker = markerStyle.Render(" → ")
		} else {
			marker = "   "
		}

		// Number
		num := numberStyle.Render(fmt.Sprintf("%2d. ", i+1))

		// Title
		title := truncateRunes(result.Title, maxWidth-10)
		var titleStyled string
		if isSelected {
			titleStyled = titleSelectedStyle.Render(title)
		} else {
			titleStyled = titleStyle.Render(title)
		}

		b.WriteString(marker + num + titleStyled + "\n")

		// URL
		url := truncateRunes(result.URL, maxWidth-6)
		urlStyled := urlStyle.Render("      " + url)
		b.WriteString(urlStyled + "\n")

		// Snippet
		if result.Snippet != "" {
			snippet := truncateRunes(result.Snippet, maxWidth-6)
			snippetStyled := snippetStyle.Render("      " + snippet)
			b.WriteString(snippetStyled + "\n")
		}

		b.WriteString("\n")
	}

	return b.String()
}

func truncateRunes(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return string(runes[:maxLen])
	}
	return string(runes[:maxLen-3]) + "..."
}
