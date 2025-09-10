# Define variables for better maintainability
BINARY_NAME = kd-gen
BUILD_DIR = ./build
MAIN_FILE = main.go
GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)

# Default target
all: fmt lint sec build test

# Format code with gofmt
fmt:
	@echo "Running gofmt..."
	@gofmt -w . || { echo "gofmt failed"; exit 1; }
	@echo "Done."

# Format imports with goimports
goimports:
	@echo "Running goimports..."
	@goimports -w . || { echo "goimports failed"; exit 1; }
	@echo "Done."

# Lint code with golangci-lint
lint:
	@echo "Running lint..."
	@golangci-lint run ./... || { echo "lint failed"; exit 1; }
	@echo "Done."

# Security check with gosec
sec:
	@echo "Running security check..."
	@gosec ./... || { echo "security check failed"; exit 1; }
	@echo "Done."

# Build the binary (depends on main.go)
build: $(MAIN_FILE)
	@echo "Building $(BINARY_NAME) for $(GOOS)/$(GOARCH)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE) || { echo "build failed"; exit 1; }
	@chmod +x $(BUILD_DIR)/$(BINARY_NAME)
	@ls -l $(BUILD_DIR)/$(BINARY_NAME)
	@echo "Build successful"

# Clean previous builds
clean:
	@echo "Cleaning up previous builds..."
	@rm -rf $(BUILD_DIR)
	@mkdir -p $(BUILD_DIR)

# Test the binary (run it to verify)
test: build
	@echo "Testing $(BINARY_NAME)..."
	@$(BUILD_DIR)/$(BINARY_NAME) || { echo "Test failed"; exit 1; }
	@echo "Test successful"

# Help target to list available commands
help:
	@echo "Available targets:"
	@echo "  all      - Run fmt, lint, sec, build, and test (default)"
	@echo "  fmt      - Format code with gofmt"
	@echo "  goimports- Format imports with goimports"
	@echo "  lint     - Run golangci-lint"
	@echo "  sec      - Run gosec for security checks"
	@echo "  build    - Build the binary"
	@echo "  clean    - Remove build artifacts"
	@echo "  test     - Build and test the binary"
	@echo "  help     - Show this help message"

# Phony targets to avoid conflicts with files
.PHONY: all fmt goimports lint sec build clean test help