-include .env
export

.PHONY: run build clean generate test test-unit test-integration test-coverage

run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

test:
	go test -v ./...

test-unit:
	go test -v -short ./...

test-integration:
	@echo "Running integration tests (using ETHEREUM_MAINNET_RPC environment variable)"
	go test -v -run Integration ./test/integration/...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html	

clean:
	rm -rf bin/

generate:
	./scripts/generate-bindings.sh	

install:
	go mod download

lint:
	golangci-lint run