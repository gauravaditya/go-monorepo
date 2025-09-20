# Makefile for go-monorepo: build, run, and local setup
SERVICES=core event consumer
swagger_api_file_dirs=$(foreach route,$(shell fd 'routes.go' internal/ | uniq), $(dir $(route)))
SWAGGER_DOCS_SERVICES_TARGET=$(foreach service,$(SERVICES),docs/$(service))

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
	go run ./cmd/event/main.go server --port=8081

consumer:
	go run ./cmd/consumer/main.go server --port=8082

clean:
	rm -rf bin

# Generate Swagger docs
swagger:
	swag init -g cmd/core/main.go -o docs

run/docs/%:
	@echo "Generating Swagger for $(@D) -> $(@F)... \n"

	$(eval api_path := $(subst run/docs/,internal/,$@))
	$(eval route_paths := $(filter internal/$(@F)/%,$(swagger_api_file_dirs)))
	$(foreach route,$(route_paths),swag init -o $(subst internal/,docs/,$(route)) -d $(route),$(api_path) -g routes.go -pdl 1 --parseInternal;)

$(SWAGGER_DOCS_SERVICES_TARGET): %: run/%

docs: $(SWAGGER_DOCS_SERVICES_TARGET)