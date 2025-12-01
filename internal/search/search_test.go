package search

import (
	"strings"
	"testing"
)

func TestParseResults(t *testing.T) {
	html := `
		<html>
			<body>
				<a class="result-link" href="https://example.com">Example Title</a>
				<td class="result-snippet">Example snippet text</td>
			</body>
		</html>
	`

	results, err := parseResults(html)
	if err != nil {
		t.Fatalf("parseResults failed: %v", err)
	}

	if len(results) == 0 {
		t.Fatal("Expected at least one result")
	}

	if results[0].Title != "Example Title" {
		t.Errorf("Expected title 'Example Title', got '%s'", results[0].Title)
	}

	if results[0].URL != "https://example.com" {
		t.Errorf("Expected URL 'https://example.com', got '%s'", results[0].URL)
	}
}

func TestParseResultsEmpty(t *testing.T) {
	html := `<html><body></body></html>`

	results, err := parseResults(html)
	if err != nil {
		t.Fatalf("parseResults failed: %v", err)
	}

	if len(results) != 0 {
		t.Errorf("Expected 0 results, got %d", len(results))
	}
}

func TestParseResultsInvalidHTML(t *testing.T) {
	html := `<html><body><a class="result-link">No href</a></body></html>`

	results, err := parseResults(html)
	if err != nil {
		t.Fatalf("parseResults failed: %v", err)
	}

	if len(results) != 0 {
		t.Errorf("Expected 0 results for link without href, got %d", len(results))
	}
}

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

func TestParseResultsMultiple(t *testing.T) {
	html := `
		<html>
			<body>
				<table>
					<tr>
						<td><a class="result-link" href="https://example1.com">Title 1</a></td>
					</tr>
					<tr>
						<td class="result-snippet">Snippet 1</td>
					</tr>
					<tr>
						<td><a class="result-link" href="https://example2.com">Title 2</a></td>
					</tr>
					<tr>
						<td class="result-snippet">Snippet 2</td>
					</tr>
				</table>
			</body>
		</html>
	`

	results, err := parseResults(html)
	if err != nil {
		t.Fatalf("parseResults failed: %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("Expected 2 results, got %d", len(results))
	}

	if results[0].Title != "Title 1" {
		t.Errorf("Expected first title 'Title 1', got '%s'", results[0].Title)
	}

	if results[1].URL != "https://example2.com" {
		t.Errorf("Expected second URL 'https://example2.com', got '%s'", results[1].URL)
	}
}

func TestParseResultsWhitespace(t *testing.T) {
	html := `
		<html>
			<body>
				<a class="result-link" href="https://example.com">  Title with spaces  </a>
				<td class="result-snippet">  Snippet with spaces  </td>
			</body>
		</html>
	`

	results, err := parseResults(html)
	if err != nil {
		t.Fatalf("parseResults failed: %v", err)
	}

	if len(results) == 0 {
		t.Fatal("Expected at least one result")
	}

	if strings.HasPrefix(results[0].Title, " ") || strings.HasSuffix(results[0].Title, " ") {
		t.Errorf("Title should be trimmed, got '%s'", results[0].Title)
	}
}
