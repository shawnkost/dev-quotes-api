# Makefile for Clean Dev Quotes API

.PHONY: swag run build tidy fmt docker-build docker-run docker-compose-up docker-compose-down

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

# Docker commands
docker-build:
	docker build -t dev-quotes-api .

docker-run:
	docker run -p 8080:8080 dev-quotes-api

docker-compose-up:
	docker-compose up -d

docker-compose-down:
	docker-compose down

# Development setup
setup: tidy swag build

# Clean up
clean:
	rm -rf bin/
	rm -rf docs/
