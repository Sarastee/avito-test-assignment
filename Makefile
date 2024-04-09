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
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3
	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports@v0.18.0
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v2@latest

#generate:
#	mkdir -p pkg/swagger
#	make generate-user-api
#	$(LOCAL_BIN)/statik -src=pkg/swagger -include='*.css,*.html,*.js,*.json,*.png'

#generate-user-api:
#	mkdir -p pkg/user_v1
#	protoc --proto_path api/user_v1 \
#	--proto_path vendor.protogen \
#	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
#	--plugin=protoc-gen-go=bin/protoc-gen-go \
#	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
#	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
#	--grpc-gateway_out=pkg/user_v1 --grpc-gateway_opt=paths=source_relative \
#	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
#	--validate_out lang=go:pkg/user_v1 --validate_opt=paths=source_relative \
#	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
#	--openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
#	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
#	api/user_v1/user.proto

#vendor-proto:
#	@if [ ! -d vendor.protogen/google ]; then \
#		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
#		mkdir -p  vendor.protogen/google/ &&\
#		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
#		rm -rf vendor.protogen/googleapis ;\
#	fi
#	@if [ ! -d vendor.protogen/validate ]; then \
#		mkdir -p vendor.protogen/validate &&\
#		git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
#		mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
#		rm -rf vendor.protogen/protoc-gen-validate ;\
#	fi
#	@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
#		mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
#		git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
#		mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
#		rm -rf vendor.protogen/openapiv2 ;\
#	fi

migration-status:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${CUR_MIGRATION_DIR} postgres ${MIGRATION_DSN} status -v

migration-up:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${CUR_MIGRATION_DIR} postgres ${MIGRATION_DSN} up -v

migration-down:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${CUR_MIGRATION_DIR} postgres ${MIGRATION_DSN} down -v

create-migration:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${CUR_MIGRATION_DIR} create $(migration_name) sql

local-create-new-migration: setup-local-env create-migration

local-migration-status: setup-local-env migration-status

local-migration-up: setup-local-env migration-up

local-migration-down: setup-local-env migration-down

prod-migration-status: setup-prod-env migration-status

prod-migration-up: setup-prod-env migration-up

prod-migration-down: setup-prod-env migration-down

local-create-new-migration: setup-local-env create-migration