.PHONY: run build lint test proto mocks

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
