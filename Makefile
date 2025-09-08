GOCMD?=go
GOTEST?=$(GOCMD) test
GOMOD?=$(GOCMD) mod
GOLANGCI?=golangci-lint
DB_URL?=$(DATABASE_URL)

.PHONY: tidy test lint ci migrate

tidy:
	$(GOMOD) tidy

test:
	$(GOTEST) ./... -race -count=1

lint: tidy
	$(GOLANGCI) run

ci: tidy lint test

migrate:
	@echo "Apply migrations (requires psql)"
	psql "$(DB_URL)" -f migrations/0001_init.sql


