# Makefile for go-monorepo: build, run, and local setup

.PHONY: all build up down core event consumer clean

all: build

build:
	go mod tidy
	go build -o bin/core ./cmd/core
	go build -o bin/event ./cmd/event
	go build -o bin/consumer ./cmd/consumer

up:
	docker-compose up -d

down:
	docker-compose down

core:
	go run ./cmd/core/main.go server --port=8080

event:
	go run ./cmd/event/main.go -port=8081

consumer:
	go run ./cmd/consumer/main.go -port=8082

clean:
	rm -rf bin

# Generate Swagger docs
swagger:
	swag init -g cmd/core/main.go -o docs
