package ui

import (
	"strings"
	"testing"

	"github.com/zarazaex69/zuk/internal/search"
)

func TestTruncate(t *testing.T) {
	tests := []struct {
		input    string
		maxLen   int
		expected string
	}{
		{"short", 10, "short"},
		{"exactly ten", 11, "exactly ten"},
		{"this is a very long string", 10, "this is..."},
		{"abc", 3, "abc"},
		{"abcd", 3, "abc"},
		{"", 5, ""},
		{"test", 2, "te"},
	}

	for _, tt := range tests {
		result := truncate(tt.input, tt.maxLen)
		if result != tt.expected {
			t.Errorf("truncate(%q, %d) = %q, want %q", tt.input, tt.maxLen, result, tt.expected)
		}
	}
}

func TestViewInput(t *testing.T) {
	m := NewModel("default", "")
	m.ready = true
	m.width = 80
	m.query = "test query"

	view := m.viewInput()

	if !strings.Contains(view, "ZUK") {
		t.Error("Input view should contain ZUK header")
	}

	if !strings.Contains(view, "test query") {
		t.Error("Input view should contain the query")
	}

	if !strings.Contains(view, "Enter") {
		t.Error("Input view should contain instructions")
	}
}

func TestViewLoading(t *testing.T) {
	m := NewModel("default", "test")
	m.ready = true
	m.width = 80

	view := m.viewLoading()

	if !strings.Contains(view, "Searching") {
		t.Error("Loading view should contain 'Searching' text")
	}
}

func TestRenderResultsListEmpty(t *testing.T) {
	m := NewModel("default", "")
	m.ready = true
	m.width = 80
	m.results = []search.Result{}

	view := m.renderResultsList()

	if !strings.Contains(view, "No results") {
		t.Error("Empty results should show 'No results' message")
	}
}

func TestRenderResultsListWithData(t *testing.T) {
	m := NewModel("default", "")
	m.ready = true
	m.width = 100
	m.results = []search.Result{
		{
			Title:   "Test Result 1",
			URL:     "https://example.com/1",
			Snippet: "This is a test snippet",
		},
		{
			Title:   "Test Result 2",
			URL:     "https://example.com/2",
			Snippet: "Another test snippet",
		},
	}

	view := m.renderResultsList()

	if !strings.Contains(view, "Test Result 1") {
		t.Error("Results should contain first result title")
	}

	if !strings.Contains(view, "https://example.com/1") {
		t.Error("Results should contain first result URL")
	}
}

func TestRenderResultsListWithSelection(t *testing.T) {
	m := NewModel("default", "")
	m.ready = true
	m.width = 100
	m.selectedIdx = 1
	m.results = []search.Result{
		{Title: "Result 1", URL: "https://example.com/1"},
		{Title: "Result 2", URL: "https://example.com/2"},
	}

	view := m.renderResultsList()

	lines := strings.Split(view, "\n")
	hasArrow := false
	for _, line := range lines {
		if strings.Contains(line, "→") && strings.Contains(line, "Result 2") {
			hasArrow = true
			break
		}
	}

	if !hasArrow {
		t.Error("Selected result should have arrow indicator")
	}
}

func TestRenderResultsListWithError(t *testing.T) {
	m := NewModel("default", "")
	m.ready = true
	m.width = 80
	m.err = &testError{msg: "test error"}

	view := m.renderResultsList()

	if !strings.Contains(view, "Error") {
		t.Error("Error view should contain 'Error' text")
	}

	if !strings.Contains(view, "test error") {
		t.Error("Error view should contain error message")
	}
}

type testError struct {
	msg string
}

func (e *testError) Error() string {
	return e.msg
}
