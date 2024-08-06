ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: build-tools
build-tools:
	go install github.com/air-verse/air@latest && \
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
	go install github.com/jackc/tern/v2@latest

.PHONY: run-sqlc
run-sqlc:
	@cd db && sqlc generate

.PHONY: run-tern
run-tern:
	@echo "Running tern with POSTGRES_PASSWORD=$(POSTGRES_PASSWORD)"
	@cd db/migrations && \
	POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) tern migrate

.PHONY: run-tern-migration
tern-migration:
	@echo "Running tern with POSTGRES_PASSWORD=$(POSTGRES_PASSWORD)"
	@cd db/migrations && \
	POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) tern new $(name)

.PHONY: run
run:
	@air --build.cmd "go build -o bin/tempstation/main cmd/tempstation/main.go" --build.bin "./bin/tempstation/main"
