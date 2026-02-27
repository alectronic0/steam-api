.PHONY: compare_steam_libs get_steam_api_list list_steam_apps web test lint fmt security coverage help

help:
	@echo "Steam API - Available Commands"
	@echo "=============================="
	@echo "  make web                  - Start HTTP server"
	@echo "  make test                 - Run all tests"
	@echo "  make test-coverage        - Run tests with coverage report"
	@echo "  make lint                 - Run golangci-lint"
	@echo "  make fmt                  - Check code formatting"
	@echo "  make fmt-fix              - Auto-fix code formatting"
	@echo "  make security             - Run gosec security scanner"
	@echo "  make vet                  - Run go vet"
	@echo "  make ci                   - Run all CI checks (test, lint, fmt, vet)"
	@echo "  make compare_steam_libs   - Run comparison tool"
	@echo "  make get_steam_api_list   - Get Steam API list"
	@echo "  make list_steam_apps      - List Steam apps"
	@echo "  make docker-build         - Build Docker image"
	@echo "  make docker-run           - Run in Docker"
	@echo "  make clean                - Remove build artifacts"

compare_steam_libs:
	go run cmd/compare_steam_libs/main.go

get_steam_api_list:
	go run cmd/get_steam_api_list/main.go

list_steam_apps:
	go run cmd/list_steam_apps/main.go

web:
	go run cmd/web/main.go

# Testing
test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Code Quality
lint:
	golangci-lint run --timeout=5m

fmt:
	@gofmt -l .
	@if [ -n "$$(gofmt -l .)" ]; then echo "❌ Formatting issues found. Run 'make fmt-fix' to fix."; exit 1; fi
	@echo "✅ Code formatting is correct"

fmt-fix:
	gofmt -w .

vet:
	go vet ./...

# Security
security:
	gosec ./...

# CI Pipeline
ci: test fmt vet lint security
	@echo "✅ All CI checks passed!"

# Docker
docker-build:
	docker build -t steam-api:latest .

docker-run:
	docker-compose up -d

docker-stop:
	docker-compose down

# Dependencies
mod-tidy:
	go mod tidy

mod-download:
	go mod download

# Cleanup
clean:
	rm -f coverage.out coverage.html
	go clean -cache -testcache
	@echo "✅ Cleanup complete"