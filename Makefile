build:
	@go build -o bin/main cmd/main.go

run:
	@go run cmd/main.go

lint:
	@golangci-lint run ./...

.PHONY: build run lint