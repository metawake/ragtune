# RagTune Makefile

.PHONY: build install test lint clean release-dry run help

# Build binary
build:
	go build -o ragtune ./cmd/ragtune

# Install to $GOPATH/bin
install:
	go install ./cmd/ragtune

# Run all tests
test:
	go test -v ./...

# Run tests with race detector
test-race:
	go test -race -v ./...

# Run linter
lint:
	golangci-lint run

# Clean build artifacts
clean:
	rm -f ragtune
	rm -rf dist/

# Dry run release (test goreleaser config)
release-dry:
	goreleaser release --snapshot --clean

# Run with sample data
run: build
	./ragtune --help

# Format code
fmt:
	go fmt ./...

# Download dependencies
deps:
	go mod download
	go mod tidy

# Show help
help:
	@echo "RagTune - EXPLAIN ANALYZE for RAG retrieval"
	@echo ""
	@echo "Usage:"
	@echo "  make build        Build binary"
	@echo "  make install      Install to GOPATH/bin"
	@echo "  make test         Run tests"
	@echo "  make lint         Run linter"
	@echo "  make clean        Clean build artifacts"
	@echo "  make release-dry  Test release build"
	@echo "  make fmt          Format code"
	@echo "  make deps         Download dependencies"
	@echo ""
	@echo "Release:"
	@echo "  git tag v0.1.0 && git push origin v0.1.0"


