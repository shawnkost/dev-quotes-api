# Makefile for Clean Dev Quotes API

.PHONY: swag run build tidy fmt

# Generate Swagger docs
swag:
	swag init -g cmd/server/main.go

# Run the app
run:
	go run ./cmd/server/main.go

# Build the binary
build:
	go build -o bin/dev-quotes-api ./cmd/server

# Tidy up go.mod/go.sum
tidy:
	go mod tidy

# Format code
fmt:
	go fmt ./...
