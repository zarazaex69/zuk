<div align="center">

<img src="assets/logo.png" alt="ZUK Logo" width="200"/>

[![Release](https://img.shields.io/github/v/release/zarazaex69/zuk?style=flat-square&logo=github&logoColor=white&color=0D1117&labelColor=0D1117)](https://github.com/zarazaex69/zuk/releases)
![Golang](https://img.shields.io/badge/-Golang-0D1117?style=flat-square&logo=go&logoColor=00A7D0)
[![License](https://img.shields.io/badge/license-BSD--2--Clause-0D1117?style=flat-square&logo=open-source-initiative&logoColor=green&labelColor=0D1117)](LICENSE)


</div>


# About

Fast and lightweight command-line interface for DuckDuckGo search. Built with Go and Bubble Tea for a smooth terminal experience, have sdk - see [Links](#links).

# Fast Start

## Install

**Fast Install (Linux/macOS):**

```bash
curl -fsSL https://raw.githubusercontent.com/zarazaex69/zuk/refs/heads/master/install.sh | bash
```

## Usage

```bash
# Start ZUK
zuk

# Duckle with zuk 
zuk text

# Set theme
zuk -t nya
zuk --theme soft

# List available themes
zuk --list-themes
```

###  Keyboard Shortcuts

- **Type** - Enter search query
- **Enter** - Execute search / Open selected result in browser
- **↑/↓ or j/k** - Navigate through results
- **Backspace** - Return to search input
- **Esc or q** - Quit application

Theme preference is saved in `~/.config/zuk/config.json`

## How It Works

ZUK uses DuckDuckGo's Lite interface to perform searches:

1. Sends POST request to `https://lite.duckduckgo.com/lite/`
2. Parses HTML response using goquery
3. Extracts search results (title, URL, snippet)
4. Renders in viewport with theme-based styling
5. Tracks cursor position with auto-scrolling
6. Opens selected URLs in default browser


<div align="center">

---

### Contact

Telegram: [zarazaex](https://t.me/zarazaexe)
<br>
Email: [zarazaex@tuta.io](mailto:zarazaex@tuta.io)
<br>
Site: [zarazaex.xyz](https://zarazaex.xyz)

</div>
