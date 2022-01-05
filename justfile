set dotenv-load := false

ROOT := justfile_directory()

# By default, print the list of recipes
_default:
    @just --list

# Run go mod tidy
tidy:
    go mod tidy

# Format all source files
fmt:
    gofumpt -extra -w .

# Run all the tests
test:
    go test ./...

# Lint all source files
lint:
    golangci-lint run --fix

# Run unit tests and linting in one go
check: tidy fmt test lint

# Set up for a new day
new day:
    go run scripts/new.go {{ day }}
