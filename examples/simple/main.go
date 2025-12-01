package main

import (
	"fmt"
	"log"

	"github.com/zarazaex69/zuk/pkg/zuk"
)

func main() {
	// Create a new client
	client := zuk.NewClient()

	// Perform a search
	results, err := client.Search("golang tutorial")
	if err != nil {
		log.Fatalf("Search failed: %v", err)
	}

	// Print results
	fmt.Printf("Found %d results:\n\n", len(results))
	for i, result := range results {
		fmt.Printf("%d. %s\n", i+1, result.Title)
		fmt.Printf("   %s\n", result.URL)
		if result.Snippet != "" {
			fmt.Printf("   %s\n", result.Snippet)
		}
		fmt.Println()
	}
}
