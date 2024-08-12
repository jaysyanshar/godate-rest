.PHONY: build run test migrate

build:
	go mod download
	go build -o godate cmd/http/main.go

run:
	go mod download
	go run cmd/http/main.go

test:
	go mod download
	go test ./...

migrate:
	go mod download
	go run cmd/migration/main.go