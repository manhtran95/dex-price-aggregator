.PHONY: run build test clean

run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

install:
	go mod download

lint:
	golangci-lint run