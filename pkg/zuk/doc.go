// Package zuk provides a simple client for searching DuckDuckGo.
//
// This package allows you to perform privacy-focused web searches
// using DuckDuckGo's Lite interface without requiring an API key.
//
// Basic usage:
//
//	client := zuk.NewClient()
//	results, err := client.Search("golang tutorial")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, result := range results {
//		fmt.Printf("%s\n%s\n%s\n\n", result.Title, result.URL, result.Snippet)
//	}
//
// Advanced usage with options:
//
//	client := zuk.NewClient()
//	opts := &zuk.SearchOptions{
//		Region:    "us-en",
//		TimeRange: "w", // Past week
//	}
//	results, err := client.SearchWithOptions("golang news", opts)
//
// The package respects DuckDuckGo's privacy principles and does not
// track or store any search queries.
package zuk
