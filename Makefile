# Define the binary output directory
BINARY_DIR := bin
BINARY_NAME := main
SOURCE_FILES := cmd/main.go
TEST_DIR := ./test

# Build the application
build:
	@go build -o $(BINARY_DIR)/$(BINARY_NAME) $(SOURCE_FILES)

# Run the application
run: build
	@./$(BINARY_DIR)/$(BINARY_NAME)

# Run the application in development mode using Air
run-dev:
	@air -c .air.toml

# Run tests
test:
	@go test -v ./...  # Run all tests in the project

# Clean build artifacts
clean:
	@rm -rf $(BINARY_DIR)/*

# Run tests with coverage
test-cov:
	@go test -cover ./...  # Run all tests with coverage
