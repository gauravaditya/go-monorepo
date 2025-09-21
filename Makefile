# Makefile for go-monorepo: build, run, and local setup
GO_BIN = $(shell which go)
SERVICES=core event consumer
swagger_api_file_dirs=$(foreach route,$(shell fd 'routes.go' internal/ | uniq), $(dir $(route)))

VERSION ?= $(if $(CI_COMMIT_SHORT_SHA),$(CI_COMMIT_SHORT_SHA),$(shell git describe --tags --always --dirty))
BUILD_SHA ?= $(if $(CI_COMMIT_SHORT_SHA),$(CI_COMMIT_SHORT_SHA),$(shell git rev-parse --short HEAD))
BUILD_BRANCH ?= $(if $(CI_COMMIT_BRANCH),$(CI_COMMIT_BRANCH),$(shell git rev-parse --abbrev-ref HEAD))

VERSION_LD_FLAGS=latestVersion=$(VERSION) commitSHA=$(BUILD_SHA) commitBranch=$(BUILD_BRANCH)
BUILD_VERSION_LD_FLAGS ?= $(foreach arg,$(VERSION_LD_FLAGS),-X 'github.com/gauravaditya/go-monorepo/pkg/clicmd.$(arg)')

define gen_targets
$(foreach service,$(SERVICES),$(1)/$(service))
endef

DOCKER_BUILD_TARGETS=$(call gen_targets,docker)
SWAGGER_DOCS_TARGETS=$(call gen_targets,docs)
LOCAL_RUN_TARGETS=$(call gen_targets,local)

.PHONY: all build up down core event consumer clean

all: clean build

up:
	docker-compose up -d

down:
	docker-compose down

clean:
	@echo "Removing previous build artifacts..."
	@rm -rf bin

# Build all services
build: $(SERVICES)

$(SERVICES): %: build/%

build/%:
	@mkdir -p bin
	@echo "Building service $* with version $(VERSION)..."

	@CGO_ENABLED=$(CGO_ENABLED) $(GO_BIN) build \
	-ldflags "$(BUILD_VERSION_LD_FLAGS)" \
	-o bin/$* \
	./cmd/$*

# Build Docker images for services
docker: $(DOCKER_BUILD_TARGETS)

$(DOCKER_BUILD_TARGETS): %: run/%

run/docker/%:
	@echo "Building Docker image for service $* with version $(VERSION)..."
	@echo "  LD Flags: $(BUILD_VERSION_LD_FLAGS)"
	@docker build -t go-monorepo/$*:$(VERSION) . \
	--build-arg SERVICE=$* \
	--build-arg BUILD_VERSION_LD_FLAGS="$(BUILD_VERSION_LD_FLAGS)" 

# Generate Swagger docs
docs: $(SWAGGER_DOCS_TARGETS)

$(SWAGGER_DOCS_TARGETS): %: run/%

run/docs/%:
	@echo "Generating Swagger for $(@D) -> $(@F)... \n"

	$(eval api_path := internal/$(@F))
	$(eval route_paths := $(filter internal/$(@F)/%,$(swagger_api_file_dirs)))
	$(foreach route,$(route_paths),swag init -o $(subst internal/,docs/,$(route)) -d $(route),$(api_path) -g routes.go -pdl 1 --parseInternal;)

# Run services locally with dynamic port assignment
local: $(LOCAL_RUN_TARGETS)

$(LOCAL_RUN_TARGETS): %: run/%

run/local/%:
	@echo "Running service $* locally..."
	@read -p "Enter port for $*: " PORT; \
	echo "Starting $* on port $$PORT..."; \
	go run ./cmd/$*/main.go server --port=$$PORT