define setup_env
	$(eval ENV_FILE := ./deploy/env/.env.local)
	@echo "- setup env $(ENV_FILE)"
	$(eval include ./deploy/env/.env.local)
	$(eval export)
endef

setup-local-env:
	$(call setup_env,local)

setup-prod-env:
	$(call setup_env,prod)

LOCAL_BIN:=$(CURDIR)/bin

CUR_MIGRATION_DIR=$(MIGRATION_DIR)
MIGRATION_DSN="host=$(PG_HOST) port=$(PG_PORT) dbname=$(POSTGRES_DB) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) sslmode=disable"

local-start-app:
	docker-compose --env-file deploy/env/.env.local -f docker-compose.local.yaml up -d --build

local-down-app:
	docker-compose --env-file deploy/env/.env.local -f docker-compose.local.yaml down -v

prod-start-app:
	docker-compose --env-file deploy/env/.env.prod -f docker-compose.prod.yaml up -d --build

prod-down-app:
	docker-compose --env-file deploy/env/.env.prod -f docker-compose.prod.yaml down -v

app-start:
	go run ./cmd/grpc_server/main.go --config=./deploy/env/.env.local

lint:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

fix-imports:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goimports -w .

test:
	go clean -testcache
	go test ./... -covermode count -coverpkg=github.com/sarastee/auth/internal/service/...,github.com/sarastee/auth/internal/api/... -count 5

test-coverage:
	go clean -testcache
	go test ./... -coverprofile=coverage.tmp.out -covermode count -coverpkg=github.com/sarastee/auth/internal/service/...,github.com/sarastee/auth/internal/api/... -count 5
	grep -v 'mocks\|config' coverage.tmp.out  > coverage.out
	rm coverage.tmp.out
	go tool cover -html=coverage.out;
	go tool cover -func=./coverage.out | grep "total";
	grep -sqFx "/coverage.out" .gitignore || echo "/coverage.out" >> .gitignore

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

local-create-new-migration: setup-local-env create-migration

local-migration-status: setup-local-env migration-status

local-migration-up: setup-local-env migration-up

local-migration-down: setup-local-env migration-down

prod-migration-status: setup-prod-env migration-status

prod-migration-up: setup-prod-env migration-up

prod-migration-down: setup-prod-env migration-down

local-create-new-migration: setup-local-env create-migration