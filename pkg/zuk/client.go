package zuk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Client represents a DuckDuckGo search client
type Client struct {
	httpClient *http.Client
	userAgent  string
}

// Result represents a single search result
type Result struct {
	Title   string
	URL     string
	Snippet string
}

// SearchOptions contains options for search requests
type SearchOptions struct {
	Region    string // e.g., "us-en", "ru-ru"
	TimeRange string // "d" (day), "w" (week), "m" (month), "y" (year)
}

// NewClient creates a new DuckDuckGo search client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
		userAgent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:145.0) Gecko/20100101 Firefox/145.0",
	}
}

// NewClientWithHTTP creates a client with custom HTTP client
func NewClientWithHTTP(httpClient *http.Client) *Client {
	return &Client{
		httpClient: httpClient,
		userAgent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:145.0) Gecko/20100101 Firefox/145.0",
	}
}

// SetUserAgent sets custom user agent
func (c *Client) SetUserAgent(ua string) {
	c.userAgent = ua
}

// Search performs a search query and returns results
func (c *Client) Search(query string) ([]Result, error) {
	return c.SearchWithOptions(query, nil)
}

// SearchWithOptions performs a search with custom options
func (c *Client) SearchWithOptions(query string, opts *SearchOptions) ([]Result, error) {
	formData := url.Values{}
	formData.Set("q", query)

	if opts != nil {
		if opts.Region != "" {
			formData.Set("kl", opts.Region)
		}
		if opts.TimeRange != "" {
			formData.Set("df", opts.TimeRange)
		}
	} else {
		formData.Set("kl", "")
		formData.Set("df", "")
	}

	req, err := http.NewRequest("POST", "https://lite.duckduckgo.com/lite/", strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "https://lite.duckduckgo.com")
	req.Header.Set("Referer", "https://lite.duckduckgo.com/")

	resp, err := c.httpClient.Do(req)
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
