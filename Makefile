# === Configuration ===

# Tools
GOLANGCI_LINT := golangci-lint

# Scripts
MIGRATE_UP_SCRIPT := scripts/migrate-up.sh

# === Phony targets ===
.PHONY: all lint test migrate

# Default target
all: lint test

# === Linting ===
lint:
	@echo "Running linters..."
	$(GOLANGCI_LINT) run ./...

# === Tests ===
test:
	@echo "Running tests..."
	go test -v ./...

# === Migrations ===
migrate:
	@echo "Applying migrations..."
	@$(MIGRATE_UP_SCRIPT)
