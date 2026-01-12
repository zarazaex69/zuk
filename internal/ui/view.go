package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type styles struct {
	logo          lipgloss.Style
	subtitle      lipgloss.Style
	separator     lipgloss.Style
	searchLabel   lipgloss.Style
	query         lipgloss.Style
	cursor        lipgloss.Style
	helpBar       lipgloss.Style
	helpKey       lipgloss.Style
	loading       lipgloss.Style
	count         lipgloss.Style
	scroll        lipgloss.Style
	err           lipgloss.Style
	gray          lipgloss.Style
	marker        lipgloss.Style
	number        lipgloss.Style
	title         lipgloss.Style
	titleSelected lipgloss.Style
	url           lipgloss.Style
	snippet       lipgloss.Style
}

func (m Model) styles() styles {
	return styles{
		logo:          lipgloss.NewStyle().Foreground(m.theme.Logo).Bold(true),
		subtitle:      lipgloss.NewStyle().Foreground(m.theme.Gray),
		separator:     lipgloss.NewStyle().Foreground(m.theme.Separator),
		searchLabel:   lipgloss.NewStyle().Foreground(m.theme.Cyan).Bold(true),
		query:         lipgloss.NewStyle().Foreground(m.theme.White),
		cursor:        lipgloss.NewStyle().Foreground(m.theme.Yellow).Bold(true),
		helpBar:       lipgloss.NewStyle().Foreground(m.theme.Gray),
		helpKey:       lipgloss.NewStyle().Foreground(m.theme.Yellow).Bold(true),
		loading:       lipgloss.NewStyle().Foreground(m.theme.Cyan).Bold(true),
		count:         lipgloss.NewStyle().Foreground(m.theme.Green),
		scroll:        lipgloss.NewStyle().Foreground(m.theme.Gray).Italic(true),
		err:           lipgloss.NewStyle().Foreground(m.theme.Red).Bold(true),
		gray:          lipgloss.NewStyle().Foreground(m.theme.Gray),
		marker:        lipgloss.NewStyle().Foreground(m.theme.Yellow).Bold(true),
		number:        lipgloss.NewStyle().Foreground(m.theme.Gray),
		title:         lipgloss.NewStyle().Foreground(m.theme.White).Bold(true),
		titleSelected: lipgloss.NewStyle().Foreground(m.theme.Yellow).Bold(true),
		url:           lipgloss.NewStyle().Foreground(m.theme.Cyan),
		snippet:       lipgloss.NewStyle().Foreground(m.theme.Gray),
	}
}

func (m Model) separator() string {
	return m.styles().separator.Render(strings.Repeat("─", m.width-4))
}

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
		return m.viewResults()
	}
	return ""
}

func (m Model) viewInput() string {
	s := m.styles()
	sep := m.separator()

	var b strings.Builder
	b.WriteString("\n  " + s.logo.Render("🦆 ZUK") + s.subtitle.Render(" - DuckDuckGo CLI Search") + "\n\n")
	b.WriteString("  " + sep + "\n\n")
	b.WriteString(s.searchLabel.Render("  Search: ") + s.query.Render(m.query) + s.cursor.Render("█") + "\n\n")
	b.WriteString("  " + sep + "\n\n")
	b.WriteString(s.helpBar.Render("  Press ") + s.helpKey.Render("Enter") + s.helpBar.Render(" to search, ") + s.helpKey.Render("Esc") + s.helpBar.Render(" to quit") + "\n")

	return b.String()
}

func (m Model) viewLoading() string {
	s := m.styles()
	sep := m.separator()

	var b strings.Builder
	b.WriteString("\n  " + s.logo.Render("🦆 ZUK") + "\n\n")
	b.WriteString("  " + sep + "\n\n")
	b.WriteString(s.loading.Render("  🔍 Searching for: ") + s.query.Render(m.query) + s.loading.Render("...") + "\n")

	return b.String()
}

func (m Model) viewResults() string {
	s := m.styles()
	sep := m.separator()

	var b strings.Builder
	b.WriteString("\n  " + s.logo.Render("🦆 ZUK") + s.count.Render(fmt.Sprintf(" (%d results)", len(m.results))) + "\n")
	b.WriteString("  " + sep + "\n")
	b.WriteString(m.viewport.View() + "\n")
	b.WriteString("  " + sep + "\n")

	scrollInfo := s.scroll.Render(fmt.Sprintf("  [%d/%d]", m.selectedIdx+1, len(m.results)))
	if m.viewport.TotalLineCount() > m.viewport.Height {
		scrollInfo += s.scroll.Render(fmt.Sprintf(" %.0f%%", m.viewport.ScrollPercent()*100))
	}
	b.WriteString(scrollInfo + "\n")

	b.WriteString(s.helpBar.Render("  ") +
		s.helpKey.Render("↑/↓") + s.helpBar.Render(" navigate  ") +
		s.helpKey.Render("Enter") + s.helpBar.Render(" open  ") +
		s.helpKey.Render("Backspace") + s.helpBar.Render(" new search  ") +
		s.helpKey.Render("q") + s.helpBar.Render(" quit"))

	return b.String()
}

func (m Model) renderResultsList() string {
	s := m.styles()

	if m.err != nil {
		return s.err.Render("  Error: " + m.err.Error())
	}

	if len(m.results) == 0 {
		return s.gray.Render("  No results found.")
	}

	maxWidth := max(m.width-8, 40)

	var b strings.Builder
	for i, result := range m.results {
		isSelected := i == m.selectedIdx

		marker := "   "
		if isSelected {
			marker = s.marker.Render(" → ")
		}

		titleStyle := s.title
		if isSelected {
			titleStyle = s.titleSelected
		}

		b.WriteString(marker + s.number.Render(fmt.Sprintf("%2d. ", i+1)) + titleStyle.Render(truncate(result.Title, maxWidth-10)) + "\n")
		b.WriteString(s.url.Render("      "+truncate(result.URL, maxWidth-6)) + "\n")

		if result.Snippet != "" {
			b.WriteString(s.snippet.Render("      "+truncate(result.Snippet, maxWidth-6)) + "\n")
		}

		b.WriteString("\n")
	}

	return b.String()
}

func truncate(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return string(runes[:maxLen])
	}
	return string(runes[:maxLen-3]) + "..."
}
