# Installation Guide

## Quick Install

### Linux / macOS

```bash
curl -fsSL https://raw.githubusercontent.com/zarazaex69/zuk/main/install.sh | bash
```

This will:
1. Detect your platform (OS and architecture)
2. Download the latest release
3. Verify checksum
4. Install to `/usr/local/bin`

### Custom Installation Directory

```bash
INSTALL_DIR=$HOME/.local/bin curl -fsSL https://raw.githubusercontent.com/zarazaex69/zuk/main/install.sh | bash
```

## Manual Installation

### From GitHub Releases

1. Go to [Releases](https://github.com/zarazaex69/zuk/releases)
2. Download the appropriate binary for your platform:
   - `zuk-linux-amd64.tar.gz` - Linux x86_64
   - `zuk-linux-arm64.tar.gz` - Linux ARM64
   - `zuk-darwin-amd64.tar.gz` - macOS Intel
   - `zuk-darwin-arm64.tar.gz` - macOS Apple Silicon
   - `zuk-windows-amd64.exe.zip` - Windows

3. Extract and install:

**Linux/macOS:**
```bash
tar -xzf zuk-*.tar.gz
sudo mv zuk-* /usr/local/bin/zuk
chmod +x /usr/local/bin/zuk
```

**Windows:**
```powershell
# Extract zuk-windows-amd64.exe.zip
# Move zuk-windows-amd64.exe to a directory in your PATH
# Rename to zuk.exe
```

### From Source

**Prerequisites:**
- Go 1.23 or later
- Make (optional)

**Build:**

```bash
# Clone repository
git clone https://github.com/zarazaex69/zuk.git
cd zuk

# Using Make
make build
sudo make install

# Or manually
go build -o zuk cmd/zuk/main.go
sudo mv zuk /usr/local/bin/
```

## Verify Installation

```bash
zuk --list-themes
```

You should see a list of available themes.

## Uninstall

```bash
sudo rm /usr/local/bin/zuk
rm -rf ~/.config/zuk
```

## Troubleshooting

### Command not found

Make sure `/usr/local/bin` is in your PATH:

```bash
echo $PATH
```

If not, add to your shell profile (`~/.bashrc`, `~/.zshrc`, etc.):

```bash
export PATH="/usr/local/bin:$PATH"
```

### Permission denied

If you get permission errors during installation:

```bash
# Install to user directory instead
INSTALL_DIR=$HOME/.local/bin curl -fsSL https://raw.githubusercontent.com/zarazaex69/zuk/main/install.sh | bash

# Add to PATH
export PATH="$HOME/.local/bin:$PATH"
```

### Checksum verification failed

If checksum verification fails, you can:
1. Try downloading again
2. Manually verify the checksum from the release page
3. Report the issue on GitHub

## Platform Support

| Platform | Architecture | Status |
|----------|-------------|--------|
| Linux | x86_64 | ✅ Supported |
| Linux | ARM64 | ✅ Supported |
| macOS | Intel | ✅ Supported |
| macOS | Apple Silicon | ✅ Supported |
| Windows | x86_64 | ✅ Supported |

## Next Steps

After installation:
1. Run `zuk` to start searching
2. Try different themes: `zuk -t nya`
3. Check out the [README](../README.md) for more features
