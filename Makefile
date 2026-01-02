.PHONY: help build install clean test lint deps tidy bench

BINARY_NAME=mo
VERSION?=0.1.0
BUILD_DIR=bin
INSTALL_DIR=/usr/local/bin

LDFLAGS=-ldflags "-w -s -X main.Version=$(VERSION)"

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the CLI binary
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/mo

install: build ## Install the CLI to system
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)..."
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/
	@echo "Installation complete. Run '$(BINARY_NAME)' to get started."

bench: ## Run API benchmarks
	@echo "Building benchmark..."
	@go build -o $(BUILD_DIR)/bench ./cmd/bench
	@$(BUILD_DIR)/bench $(ARGS)

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

test: ## Run tests
	@echo "Running tests..."
	@go test -v -race ./...

lint: ## Run linters
	@echo "Running linters..."
	@golangci-lint run ./...

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download

tidy: ## Tidy go modules
	@echo "Tidying go modules..."
	@go mod tidy
