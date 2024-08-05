.PHONY: build-tools
build-tools:
	go install github.com/air-verse/air@latest && \
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
	go install github.com/jackc/tern/v2@latest

# Target to run sqlc
run-sqlc:
	cd db && sqlc generate

# File containing the PostgreSQL password
PASSWORD_FILE := postgresql_password.txt

# Read the password from the file and export it as an environment variable
export POSTGRES_PASSWORD := $(shell cat $(PASSWORD_FILE))

# Target to run tern
run-tern:
	@echo "Running tern with POSTGRES_PASSWORD=$(POSTGRES_PASSWORD)"
	@cd db/migrations && \
	POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) tern migrate


# Target to create new tern migration
tern-migration:
	@echo "Running tern with POSTGRES_PASSWORD=$(POSTGRES_PASSWORD)"
	@cd db/migrations && \
	POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) tern new $(name)
