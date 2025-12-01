package search

import (
	"runtime"
	"testing"
)

func TestOpenBrowser(t *testing.T) {
	// This test just ensures the function doesn't panic
	// Actual browser opening is hard to test in CI
	url := "https://example.com"

	err := OpenBrowser(url)

	// On some systems without display, this might fail, which is okay
	if err != nil && runtime.GOOS == "linux" {
		t.Logf("OpenBrowser failed on Linux (expected in headless environment): %v", err)
	}
}

func TestOpenBrowserEmpty(t *testing.T) {
	err := OpenBrowser("")

	// Should not panic with empty URL
	if err != nil {
		t.Logf("OpenBrowser with empty URL returned error: %v", err)
	}
}
