.PHONY: help build install clean test lint deps tidy

BINARY_NAME=zuk
VERSION?=0.1.0
BUILD_DIR=bin
INSTALL_DIR=/usr/local/bin

LDFLAGS=-ldflags "-w -s -X main.Version=$(VERSION)"

help: 
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: 
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/zuk

install: build 
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)..."
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/
	@echo "Installation complete. Run '$(BINARY_NAME)' to get started."

clean: 
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

test: 
	@echo "Running tests..."
	@go test -v -race ./...

lint: 
	@echo "Running linters..."
	@golangci-lint run ./...

deps:
	@echo "Downloading dependencies..."
	@go mod download

tidy: 
	@echo "Tidying go modules..."
	@go mod tidy
