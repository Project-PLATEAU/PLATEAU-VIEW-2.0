lint:
	golangci-lint run --fix

test:
	go test -race -v ./...

build:
	go build ./cmd/reearth-cms

run-app:
	go run ./cmd/reearth-cms

run-db:
	docker compose up -d reearth-cms-mongo

gql:
	go generate ./internal/adapter/gql

oapi:
	go generate ./pkg/integrationapi && go generate ./internal/adapter/integration

i18n:
	go run github.com/reearth/reearthx/tools i18n-extract -l en,ja -o i18n

i18n-ci:
	go run github.com/reearth/reearthx/tools i18n-extract -l en,ja -o i18n -f

.PHONY: lint test build run-app run-db gql i18n i18n-ci oapi
