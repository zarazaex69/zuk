# 🪿 ZUK Library Examples

This directory contains examples of using ZUK as a Go library.

## 🦆 Simple Example

Basic usage with default settings:

```bash
cd simple
go run main.go
```

## 🦆  Advanced Example

Custom HTTP client, user agent, and search options:

```bash
cd advanced
go run main.go
```

## 🦆 API Reference

### Creating a Client

```go
// Default client
client := zuk.NewClient()

// Custom HTTP client
httpClient := &http.Client{Timeout: 10 * time.Second}
client := zuk.NewClientWithHTTP(httpClient)

// Set custom user agent
client.SetUserAgent("MyApp/1.0")
```

### Searching

```go
// Simple search
results, err := client.Search("golang")

// Search with options
opts := &zuk.SearchOptions{
    Region:    "us-en",  // Region code
    TimeRange: "w",      // d=day, w=week, m=month, y=year
}
results, err := client.SearchWithOptions("golang", opts)
```

### Result Structure

```go
type Result struct {
    Title   string  // Result title
    URL     string  // Result URL
    Snippet string  // Result description/snippet
}
```

## 🦆 Region Codes

Common region codes:
- `us-en` - United States (English)
- `uk-en` - United Kingdom
- `ru-ru` - Russia
- `de-de` - Germany
- `fr-fr` - France
- `jp-jp` - Japan

## 🦆 Time Ranges

- `d` - Past day
- `w` - Past week
- `m` - Past month
- `y` - Past year
- `` (empty) - Any time
