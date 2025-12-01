package search

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Result struct {
	Title   string
	URL     string
	Snippet string
}

func Search(query string) ([]Result, error) {
	formData := url.Values{}
	formData.Set("q", query)
	formData.Set("kl", "")
	formData.Set("df", "")

	req, err := http.NewRequest("POST", "https://lite.duckduckgo.com/lite/", strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:145.0) Gecko/20100101 Firefox/145.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "https://lite.duckduckgo.com")
	req.Header.Set("Referer", "https://lite.duckduckgo.com/")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform search: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return parseResults(string(body))
}

func parseResults(html string) ([]Result, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	var results []Result

	doc.Find("a.result-link").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		snippet := ""
		snippetElem := s.ParentsUntil("table").Next().Find("td.result-snippet")
		if snippetElem.Length() > 0 {
			snippet = strings.TrimSpace(snippetElem.Text())
		}

		results = append(results, Result{
			Title:   strings.TrimSpace(title),
			URL:     href,
			Snippet: snippet,
		})
	})

	return results, nil
}
