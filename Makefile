.PHONY: run build test clean generate

run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

generate:
	./scripts/generate-bindings.sh	

install:
	go mod download

lint:
	golangci-lint run