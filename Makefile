SRC=$(word 2, $(MAKECMDGOALS))
SRC_BASE=$(basename $(SRC))

.PHONY: build run test migrate download mock

build:
	go build -o godate cmd/http/main.go

run:
	go run cmd/http/main.go

test:
	go test ./...

migrate:
	go run cmd/migration/main.go

download:
	go mod download

mock:
	mockgen -source=$(SRC) -destination=$(SRC_BASE)_mock.go -package=$(shell awk '/^package / {print $$2}' $(SRC))

# Prevent make from treating the argument as a target
%:
	@: