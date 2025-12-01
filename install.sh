#!/usr/bin/env bash
set -e

# ZUK CLI Installation Script
# Usage: curl -fsSL https://raw.githubusercontent.com/zarazaex69/zuk/main/install.sh | bash

REPO="zarazaex69/zuk"
BINARY_NAME="zuk"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Logging functions
log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

log_step() {
    echo -e "${CYAN}[STEP]${NC} $1"
}

# Print banner
print_banner() {
    echo -e "${CYAN}"
    echo "  🦆 ZUK - DuckDuckGo CLI"
    echo "  ======================="
    echo -e "${NC}"
}

# Detect OS and architecture
detect_platform() {
    log_step "Detecting platform..."
    
    local os=$(uname -s | tr '[:upper:]' '[:lower:]')
    local arch=$(uname -m)

    case "$os" in
        linux)
            OS="linux"
            ;;
        darwin)
            OS="darwin"
            ;;
        *)
            log_error "Unsupported operating system: $os"
            exit 1
            ;;
    esac

    case "$arch" in
        x86_64|amd64)
            ARCH="amd64"
            ;;
        aarch64|arm64)
            ARCH="arm64"
            ;;
        *)
            log_error "Unsupported architecture: $arch"
            exit 1
            ;;
    esac

    PLATFORM="${OS}-${ARCH}"
    log_info "Detected platform: $PLATFORM"
}

# Get latest release version
get_latest_version() {
    log_step "Fetching latest release..."
    
    LATEST_VERSION=$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    
    if [ -z "$LATEST_VERSION" ]; then
        log_error "Failed to fetch latest version"
        exit 1
    fi
    
    log_info "Latest version: $LATEST_VERSION"
}

# Download and install binary
install_binary() {
    log_step "Installing ZUK..."
    
    local binary_name="${BINARY_NAME}-${PLATFORM}"
    local archive_name="${binary_name}.tar.gz"
    local download_url="https://github.com/${REPO}/releases/download/${LATEST_VERSION}/${archive_name}"
    local checksum_url="${download_url}.sha256"

    log_info "Downloading $archive_name..."

    # Create temporary directory
    local tmp_dir=$(mktemp -d)
    cd "$tmp_dir"

    # Download binary archive
    if ! curl -fsSL -o "$archive_name" "$download_url"; then
        log_error "Failed to download binary"
        rm -rf "$tmp_dir"
        exit 1
    fi

    # Download checksum
    if ! curl -fsSL -o "${archive_name}.sha256" "$checksum_url"; then
        log_warn "Failed to download checksum, skipping verification"
    else
        log_info "Verifying checksum..."
        if command -v sha256sum >/dev/null 2>&1; then
            sha256sum -c "${archive_name}.sha256" || {
                log_error "Checksum verification failed"
                rm -rf "$tmp_dir"
                exit 1
            }
        else
            log_warn "sha256sum not found, skipping checksum verification"
        fi
    fi

    # Extract archive
    log_info "Extracting archive..."
    tar -xzf "$archive_name"

    # Install binary
    log_info "Installing to $INSTALL_DIR..."
    if [ -w "$INSTALL_DIR" ]; then
        mv "$binary_name" "$INSTALL_DIR/$BINARY_NAME"
        chmod +x "$INSTALL_DIR/$BINARY_NAME"
    else
        log_info "Requesting sudo access to install to $INSTALL_DIR..."
        sudo mv "$binary_name" "$INSTALL_DIR/$BINARY_NAME"
        sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"
    fi

    # Cleanup
    cd - >/dev/null
    rm -rf "$tmp_dir"
    
    log_info "Installation complete!"
}

# Verify installation
verify_installation() {
    log_step "Verifying installation..."
    
    if command -v "$BINARY_NAME" >/dev/null 2>&1; then
        log_info "✓ ZUK installed successfully!"
    else
        log_error "Installation failed: $BINARY_NAME not found in PATH"
        log_info "Make sure $INSTALL_DIR is in your PATH"
        exit 1
    fi
}

# Print usage instructions
print_usage() {
    echo ""
    echo -e "${GREEN}Quick Start:${NC}"
    echo "  zuk                    # Start interactive search"
    echo "  zuk -t nya             # Use Nya theme"
    echo "  zuk --list-themes      # List available themes"
    echo ""
    echo -e "${GREEN}Themes:${NC}"
    echo "  default, monochrome, black, soft, nya"
    echo ""
    echo -e "${GREEN}Documentation:${NC}"
    echo "  https://github.com/${REPO}"
    echo ""
}

# Main installation flow
main() {
    print_banner
    detect_platform
    get_latest_version
    install_binary
    verify_installation
    print_usage
}

main "$@"
