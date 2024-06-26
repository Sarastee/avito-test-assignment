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

local-db-start:
	docker-compose --env-file deploy/env/.env.local -f docker-compose.local.yaml up -d --build

local-db-down:
	docker-compose --env-file deploy/env/.env.local -f docker-compose.local.yaml down -v

local-app-start:
	go run ./cmd/service/main.go --config=./deploy/env/.env.local

lint:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

fix-imports:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goimports -w .

swagger:
	mkdir -p pkg/swagger
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/swag init -g ./cmd/service/main.go -o ./pkg/swagger
	$(LOCAL_BIN)/statik -src=pkg/swagger/ -include='*.css,*.html,*.js,*.json,*.png'


test:
	docker-compose --env-file deploy/env/.env.test -f docker-compose.e2e.yaml up -d --build
	make wait_container
	make show_logs
	docker-compose --env-file deploy/env/.env.test -f docker-compose.e2e.yaml down -v
        
wait_container:
	@if [ "$$CI" = "true" ]; then \
  		CONTAINER_NAME="avito-test_e2e_1"; \
  	else \
  	    CONTAINER_NAME="avito-test-e2e-1"; \
	fi; \
	docker wait "$$CONTAINER_NAME"

show_logs:
	@if [ "$$CI" = "true" ]; then \
  		CONTAINER_NAME="avito-test_e2e_1"; \
  	else \
  	    CONTAINER_NAME="avito-test-e2e-1"; \
	fi; \
	docker logs "$$CONTAINER_NAME"

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2
	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports@v0.18.0
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v2@latest
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@latest
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@v0.1.7

migration-status:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${CUR_MIGRATION_DIR} postgres ${MIGRATION_DSN} status -v

migration-up:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${CUR_MIGRATION_DIR} postgres ${MIGRATION_DSN} up -v

migration-down:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${CUR_MIGRATION_DIR} postgres ${MIGRATION_DSN} down -v

create-migration:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${CUR_MIGRATION_DIR} create testdata sql