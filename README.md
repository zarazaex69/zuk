# ZUK - DuckDuckGo CLI Search

[![Release](https://img.shields.io/github/v/release/zarazaex69/zuk?style=flat-square&logo=github&color=blue)](https://github.com/zarazaex69/zuk/releases)
[![Go Version](https://img.shields.io/github/go-mod/go-version/zarazaex69/zuk?style=flat-square&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-BSD-green?style=flat-square)](LICENSE)

Fast and lightweight command-line interface for DuckDuckGo search. Built with Go and Bubble Tea for a smooth terminal experience.

<p align="center">
  <img src="assets/logo.png" alt="ZUK Logo" width="200"/>
</p>

## Quick Start

### Install

```bash
# Using make
make install

# Or manually
go build -o zuk cmd/zuk/main.go
sudo mv zuk /usr/local/bin/
```

## Overview

ZUK provides a privacy-focused search experience directly in your terminal. No API keys, no tracking, just fast DuckDuckGo searches with an intuitive TUI interface.

## Key Features

- **Privacy First** - Uses DuckDuckGo Lite API
- **Interactive TUI** - Built with Bubble Tea
- **Fast & Lightweight** - Single binary, no dependencies
- **Browser Integration** - Open results in your default browser
- **Cross-Platform** - Linux, macOS, Windows support
- **Keyboard Navigation** - Vim-style keybindings

## Technology Stack

- **Go 1.23** - High-performance and fast compilation
- **Bubble Tea** - Modern TUI framework
- **goquery** - HTML parsing for search results
- **Make** - Build automation

## Usage

```bash
# Start ZUK
zuk

# Or using make
make build && ./bin/zuk
```

### Keyboard Shortcuts

- **Type** - Enter search query
- **Enter** - Execute search / Open selected result in browser
- **↑/↓ or j/k** - Navigate through results
- **Backspace** - Return to search input
- **Esc or q** - Quit application

## Development

### Prerequisites

- Go 1.23+
- Make (optional)

### Build from Source

```bash
# Clone repository
git clone https://github.com/zarazaex69/zuk.git
cd zuk

# Download dependencies
make deps

# Build binary
make build

# Run
./bin/zuk
```

### Project Structure

```
zuk/
├── cmd/
│   └── zuk/          # CLI entry point
├── internal/
│   ├── app/          # Application initialization
│   ├── ui/           # Bubble Tea UI components
│   │   ├── model.go  # State management
│   │   ├── view.go   # Rendering logic
│   │   └── update.go # Event handling
│   └── search/       # DuckDuckGo search logic
│       ├── search.go # API client
│       └── browser.go # Browser integration
├── assets/           # Resources (logo)
├── Makefile          # Build automation
└── go.mod            # Go dependencies
```

## Make Commands

```bash
make help      # Show available commands
make build     # Build the binary
make install   # Install to /usr/local/bin
make clean     # Remove build artifacts
make test      # Run tests
make lint      # Run linters
make deps      # Download dependencies
make tidy      # Tidy go modules
```

## How It Works

ZUK uses DuckDuckGo's Lite interface to perform searches:

1. Sends POST request to `https://lite.duckduckgo.com/lite/`
2. Parses HTML response using goquery
3. Extracts search results (title, URL, snippet)
4. Displays in interactive TUI
5. Opens selected URLs in default browser

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

BSD License - See LICENSE file for details

## Author

**zarazaex** - [GitHub](https://github.com/zarazaex69)

## Links

- GitHub: [github.com/zarazaex69/zuk](https://github.com/zarazaex69/zuk)
- Issues: [github.com/zarazaex69/zuk/issues](https://github.com/zarazaex69/zuk/issues)

## Acknowledgments

- Powered by [DuckDuckGo](https://duckduckgo.com)
- Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- HTML parsing by [goquery](https://github.com/PuerkitoBio/goquery)
