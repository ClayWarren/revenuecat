.PHONY: all fmt lint typecheck test build

all: fmt lint typecheck test build

fmt:
	go fmt ./...

lint:
	golangci-lint run

typecheck:
	go build -v -o /dev/null ./...

test:
	go test -v -race -cover ./...

build:
	go build -v ./...
