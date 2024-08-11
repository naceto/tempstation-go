ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: build-tools
build-tools:
	go install github.com/air-verse/air@latest && \
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
	go install github.com/jackc/tern/v2@latest && \
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: run-sqlc
run-sqlc:
	@cd db && sqlc generate

.PHONY: run-tern
run-tern:
	@echo "Running tern with POSTGRES_PASSWORD=$(POSTGRES_PASSWORD)"
	@cd db/migrations && \
	POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) tern migrate

.PHONY: new-migration
new-tern-migration:
	@echo "Running tern with POSTGRES_PASSWORD=$(POSTGRES_PASSWORD)"
	@cd db/migrations && \
	POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) tern new $(name)

.PHONY: run
run:
	@air --build.cmd "go build -o bin/tempstation/main cmd/tempstation/main.go" --build.bin "./bin/tempstation/main"

.PHONY: openapi
openapi:
	@oapi-codegen --config=api/common/config.yaml api/common/openapi.yaml
	@oapi-codegen --config=api/generic/config.yaml api/generic/openapi.yaml
	@oapi-codegen --config=api/sensors/config.yaml api/sensors/openapi.yaml
	@cp api/common/openapi.yaml web/common/openapi.yaml
	@cp api/generic/openapi.yaml web/generic/openapi.yaml
	@cp api/sensors/openapi.yaml web/sensors/openapi.yaml

.PHONY: run-dblab
run-dblab:
	dblab --host localhost --user tempstation_admin --db tempstation_db --pass $(POSTGRES_PASSWORD) --ssl disable --port 5432 --driver postgres --limit 50

# goverter
# mockery
