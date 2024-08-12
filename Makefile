.PHONY: build run test migrate

build:
	go build -o godate cmd/http/main.go

run:
	go run cmd/http/main.go

test:
	go test ./...

migrate:
	go run cmd/migration/main.go