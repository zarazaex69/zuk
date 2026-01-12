package ui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zarazaex69/zuk/internal/search"
)

func TestNewModel(t *testing.T) {
	m := NewModel("default", "")

	if m.state != stateInput {
		t.Errorf("Expected initial state to be stateInput, got %v", m.state)
	}

	if m.query != "" {
		t.Errorf("Expected empty query, got %q", m.query)
	}

	if m.selectedIdx != 0 {
		t.Errorf("Expected selectedIdx to be 0, got %d", m.selectedIdx)
	}
}

func TestNewModelWithQuery(t *testing.T) {
	m := NewModel("default", "test query")

	if m.state != stateLoading {
		t.Errorf("Expected state to be stateLoading with initial query, got %v", m.state)
	}

	if m.query != "test query" {
		t.Errorf("Expected query 'test query', got %q", m.query)
	}
}

func TestModelInit(t *testing.T) {
	m := NewModel("default", "")
	cmd := m.Init()

	if cmd != nil {
		t.Error("Init should return nil command without initial query")
	}
}

func TestUpdateWindowSize(t *testing.T) {
	m := NewModel("default", "")
	msg := tea.WindowSizeMsg{Width: 100, Height: 50}

	updated, _ := m.Update(msg)
	updatedModel := updated.(Model)

	if updatedModel.width != 100 {
		t.Errorf("Expected width 100, got %d", updatedModel.width)
	}

	if updatedModel.height != 50 {
		t.Errorf("Expected height 50, got %d", updatedModel.height)
	}
}

func TestUpdateSearchResult(t *testing.T) {
	m := NewModel("default", "")
	m.state = stateLoading

	results := []search.Result{
		{Title: "Test", URL: "https://example.com"},
	}

	msg := searchResultMsg{results: results, err: nil}
	updated, _ := m.Update(msg)
	updatedModel := updated.(Model)

	if updatedModel.state != stateResults {
		t.Errorf("Expected state to be stateResults, got %v", updatedModel.state)
	}

	if len(updatedModel.results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(updatedModel.results))
	}

	if updatedModel.selectedIdx != 0 {
		t.Errorf("Expected selectedIdx to be reset to 0, got %d", updatedModel.selectedIdx)
	}
}

func TestUpdateInputBackspace(t *testing.T) {
	m := NewModel("default", "")
	m.query = "test"

	msg := tea.KeyMsg{Type: tea.KeyBackspace}

	updated, _ := m.updateInput(msg)
	updatedModel := updated.(Model)

	if updatedModel.query != "tes" {
		t.Errorf("Expected query 'tes', got %q", updatedModel.query)
	}
}

func TestUpdateInputTyping(t *testing.T) {
	m := NewModel("default", "")
	m.query = "test"

	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}}

	updated, _ := m.updateInput(msg)
	updatedModel := updated.(Model)

	if updatedModel.query != "testa" {
		t.Errorf("Expected query 'testa', got %q", updatedModel.query)
	}
}

func TestUpdateResultsNavigation(t *testing.T) {
	m := NewModel("default", "")
	m.state = stateResults
	m.height = 20
	m.results = []search.Result{
		{Title: "Result 1"},
		{Title: "Result 2"},
		{Title: "Result 3"},
	}

	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}}
	updated, _ := m.updateResults(msg)
	updatedModel := updated.(Model)

	if updatedModel.selectedIdx != 1 {
		t.Errorf("Expected selectedIdx 1 after down, got %d", updatedModel.selectedIdx)
	}

	msg = tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'k'}}
	updated, _ = updatedModel.updateResults(msg)
	updatedModel = updated.(Model)

	if updatedModel.selectedIdx != 0 {
		t.Errorf("Expected selectedIdx 0 after up, got %d", updatedModel.selectedIdx)
	}
}

func TestUpdateResultsBackspace(t *testing.T) {
	m := NewModel("default", "")
	m.state = stateResults
	m.query = "test"
	m.results = []search.Result{{Title: "Result"}}
	m.selectedIdx = 1

	msg := tea.KeyMsg{Type: tea.KeyBackspace}
	updated, _ := m.updateResults(msg)
	updatedModel := updated.(Model)

	if updatedModel.state != stateInput {
		t.Errorf("Expected state stateInput, got %v", updatedModel.state)
	}

	if updatedModel.query != "" {
		t.Errorf("Expected empty query, got %q", updatedModel.query)
	}

	if len(updatedModel.results) != 0 {
		t.Errorf("Expected no results, got %d", len(updatedModel.results))
	}

	if updatedModel.selectedIdx != 0 {
		t.Errorf("Expected selectedIdx 0, got %d", updatedModel.selectedIdx)
	}
}

func TestGetTheme(t *testing.T) {
	theme := GetTheme("default")
	if theme.Name != "Default" {
		t.Errorf("Expected theme name 'Default', got %q", theme.Name)
	}

	theme = GetTheme("nonexistent")
	if theme.Name != "Default" {
		t.Errorf("Expected fallback to default theme, got %q", theme.Name)
	}
}

func TestGetThemeNames(t *testing.T) {
	names := GetThemeNames()
	if len(names) == 0 {
		t.Error("Expected at least one theme name")
	}

	found := false
	for _, name := range names {
		if name == "default" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected 'default' theme in theme names")
	}
}
