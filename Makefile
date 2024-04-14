define setup_env
	$(eval ENV_FILE := ./deploy/env/.env.prod)
	@echo "- setup env $(ENV_FILE)"
	$(eval include ./deploy/env/.env.prod)
	$(eval export)
endef

setup-prod-env:
	$(call setup_env)

LOCAL_BIN:=$(CURDIR)/bin

CUR_MIGRATION_DIR=$(MIGRATION_DIR)
MIGRATION_DSN="host=$(PG_HOST) port=$(PG_PORT) dbname=$(POSTGRES_DB) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) sslmode=disable"

app-start:
	docker-compose --env-file deploy/env/.env.prod -f docker-compose.prod.yaml up -d --build

app-down:
	docker-compose --env-file deploy/env/.env.prod -f docker-compose.prod.yaml down -v

app-restart:
	make app-down
	make app-start

local-start-app:
	docker-compose --env-file deploy/env/.env.local -f docker-compose.local.yaml up -d --build

local-down-app:
	docker-compose --env-file deploy/env/.env.local -f docker-compose.local.yaml down -v

local-app-start:
	go run ./cmd/grpc_server/main.go --config=./deploy/env/.env.local

lint:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

fix-imports:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goimports -w .

test:
	docker-compose --env-file deploy/env/.env.test -f docker-compose.e2e.yaml up -d --build
	docker wait avito-test-e2e-1
	docker logs avito-test-e2e-1
	docker-compose --env-file deploy/env/.env.test -f docker-compose.e2e.yaml down -v

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2
	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports@v0.18.0
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v2@latest

migration-status:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${CUR_MIGRATION_DIR} postgres ${MIGRATION_DSN} status -v

migration-up:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${CUR_MIGRATION_DIR} postgres ${MIGRATION_DSN} up -v

migration-down:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${CUR_MIGRATION_DIR} postgres ${MIGRATION_DSN} down -v

create-migration:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${CUR_MIGRATION_DIR} create testdata sql