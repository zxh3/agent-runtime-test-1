.PHONY: build run test clean build-amd64

# Variables
BINARY_NAME=agent-runtime
GO=go
GOFMT=gofmt
GOLINT=golangci-lint

# Default target
all: build

# Build the application
build:
	$(GO) mod tidy
	$(GO) build -o $(BINARY_NAME) ./cmd/agent_workshop_runtime

# Build for AMD64 architecture
build-amd64:
	$(GO) mod tidy
	GOOS=linux GOARCH=amd64 $(GO) build -o $(BINARY_NAME)-amd64 ./cmd/agent_workshop_runtime

# Run the application
run: build
	./$(BINARY_NAME)

# Run the application with hot reload (requires air)
dev:
	air

# Run tests
test:
	$(GO) test -v ./...

# Run tests with coverage
test-coverage:
	$(GO) test -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out

# Format code
fmt:
	$(GOFMT) -w .

# Lint code
lint:
	$(GOLINT) run

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)-amd64
	rm -f coverage.out

# Install development dependencies
install-dev:
	$(GO) install github.com/cosmtrek/air@latest
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Help target
help:
	@echo "Available targets:"
	@echo "  build         - Build the application"
	@echo "  build-amd64   - Build the application for AMD64 architecture"
	@echo "  run           - Build and run the application"
	@echo "  dev           - Run with hot reload (requires air)"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  fmt           - Format code"
	@echo "  lint          - Lint code"
	@echo "  clean         - Clean build artifacts"
	@echo "  install-dev   - Install development dependencies"
