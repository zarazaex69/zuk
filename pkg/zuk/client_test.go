package zuk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Fatal("NewClient returned nil")
	}
	if client.httpClient == nil {
		t.Error("httpClient is nil")
	}
	if client.userAgent == "" {
		t.Error("userAgent is empty")
	}
}

func TestSetUserAgent(t *testing.T) {
	client := NewClient()
	customUA := "CustomBot/1.0"
	client.SetUserAgent(customUA)

	if client.userAgent != customUA {
		t.Errorf("Expected user agent %q, got %q", customUA, client.userAgent)
	}
}

func TestSearchWithMockServer(t *testing.T) {
	mockHTML := `
		<html>
			<body>
				<a class="result-link" href="https://example.com">Example Title</a>
				<td class="result-snippet">Example snippet</td>
			</body>
		</html>
	`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockHTML))
	}))
	defer server.Close()

	// Note: This test would need to modify the URL in the client
	// For now, it's a placeholder showing the structure
}

func TestParseResults(t *testing.T) {
	html := `
		<html>
			<body>
				<a class="result-link" href="https://example.com">Test Title</a>
				<td class="result-snippet">Test snippet</td>
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

	if results[0].Title != "Test Title" {
		t.Errorf("Expected title 'Test Title', got %q", results[0].Title)
	}

	if results[0].URL != "https://example.com" {
		t.Errorf("Expected URL 'https://example.com', got %q", results[0].URL)
	}

	if results[0].Snippet != "Test snippet" {
		t.Errorf("Expected snippet 'Test snippet', got %q", results[0].Snippet)
	}
}

func TestSearchWithOptions(t *testing.T) {
	client := NewClient()

	opts := &SearchOptions{
		Region:    "us-en",
		TimeRange: "w",
	}

	// This would make a real request, so we skip in unit tests
	// In integration tests, you would test this
	_ = client
	_ = opts
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

func TestParseResultsMultiple(t *testing.T) {
	html := `
		<html>
			<body>
				<a class="result-link" href="https://example1.com">Title 1</a>
				<td class="result-snippet">Snippet 1</td>
				<a class="result-link" href="https://example2.com">Title 2</a>
				<td class="result-snippet">Snippet 2</td>
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
		t.Errorf("Expected first title 'Title 1', got %q", results[0].Title)
	}

	if results[1].URL != "https://example2.com" {
		t.Errorf("Expected second URL 'https://example2.com', got %q", results[1].URL)
	}
}
