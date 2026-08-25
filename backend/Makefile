.PHONY: run build lint test proto mocks migrate-up migrate-down migrate-create sqlc

run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

lint:
	golangci-lint run ./...

test:
	go test ./... -v

proto:
	buf generate

mocks:
	mockery

MIGRATIONS_PATH=migrations

migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down 1

migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)

sqlc:
	sqlc generate
