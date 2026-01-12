package search

import (
	"testing"
)

func TestResultStruct(t *testing.T) {
	result := Result{
		Title:   "Test Title",
		URL:     "https://test.com",
		Snippet: "Test snippet",
	}

	if result.Title != "Test Title" {
		t.Errorf("Expected title 'Test Title', got '%s'", result.Title)
	}

	if result.URL != "https://test.com" {
		t.Errorf("Expected URL 'https://test.com', got '%s'", result.URL)
	}

	if result.Snippet != "Test snippet" {
		t.Errorf("Expected snippet 'Test snippet', got '%s'", result.Snippet)
	}
}

func TestResultTypeAlias(t *testing.T) {
	var r Result
	r.Title = "Test"
	r.URL = "https://example.com"
	r.Snippet = "Snippet"

	if r.Title != "Test" {
		t.Errorf("Expected title 'Test', got '%s'", r.Title)
	}
}
