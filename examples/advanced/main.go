package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/zarazaex69/zuk/pkg/zuk"
)

func main() {
	// Create a custom HTTP client with timeout
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create client with custom HTTP client
	client := zuk.NewClientWithHTTP(httpClient)

	// Set custom user agent
	client.SetUserAgent("MyApp/1.0")

	// Search with options
	opts := &zuk.SearchOptions{
		Region:    "us-en", // US English
		TimeRange: "w",     // Past week
	}

	results, err := client.SearchWithOptions("golang news", opts)
	if err != nil {
		log.Fatalf("Search failed: %v", err)
	}

	// Print results
	fmt.Printf("Found %d results from the past week:\n\n", len(results))
	for i, result := range results {
		fmt.Printf("%d. %s\n", i+1, result.Title)
		fmt.Printf("   %s\n", result.URL)
		if result.Snippet != "" {
			fmt.Printf("   %s\n", result.Snippet)
		}
		fmt.Println()

		// Limit to first 5 results
		if i >= 4 {
			break
		}
	}
}
