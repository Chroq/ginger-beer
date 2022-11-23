.PHONY: build run test-coverage

all: build

build:
	@mkdir -p bin
	@go build -o bin/main -v cmd/app/main.go

run:
	@go run cmd/app/main.go $(ARGS)

test-coverage:
	@go clean -testcache
	@go test --coverprofile=coverage.out ./... > /dev/null
	@go tool cover -func=coverage.out | grep total