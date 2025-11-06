# Colors
BLUE=\033[34m
NC=\033[0m # No Color

# Makefile for go-monorepo: build, run, and local setup
GO_BIN = $(shell which go)
SERVICES=core event consumer # list of services
swagger_api_file_dirs=$(foreach route,$(shell fd 'routes.go' internal/ | uniq), $(dir $(route)))

VERSION ?= $(if $(CI_COMMIT_SHORT_SHA),$(CI_COMMIT_SHORT_SHA),$(shell git describe --tags --always --dirty))
BUILD_SHA ?= $(if $(CI_COMMIT_SHORT_SHA),$(CI_COMMIT_SHORT_SHA),$(shell git rev-parse --short HEAD))
BUILD_BRANCH ?= $(if $(CI_COMMIT_BRANCH),$(CI_COMMIT_BRANCH),$(shell git rev-parse --abbrev-ref HEAD))

VERSION_LD_FLAGS=latestVersion=$(VERSION) commitSHA=$(BUILD_SHA) commitBranch=$(BUILD_BRANCH)
BUILD_VERSION_LD_FLAGS ?= $(foreach arg,$(VERSION_LD_FLAGS),-X 'github.com/gauravaditya/go-monorepo/pkg/clicmd.$(arg)')

define gen_targets
$(foreach service,$(SERVICES),$(1)/$(service))
endef

define help_message
	$(eval a1= $(shell echo $(1) | cut -d'/' -f1))
	$(eval a2= $(shell echo $(1) | cut -d'/' -f2))

	$(if $(findstring build/,$(1)),@printf "$(BLUE)%-40s$(NC) %s\n" "$(1)" " build $(a2) service",)
	$(if $(findstring docker/,$(1)),@printf "$(BLUE)%-40s$(NC) %s\n" "$(1)" " build docker image for $(a2) service",)
	$(if $(findstring local/,$(1)),@printf "$(BLUE)%-40s$(NC) %s\n" "$(1)" " run $(a2) service locally",)
	$(if $(findstring docs/,$(1)),@printf "$(BLUE)%-40s$(NC) %s\n" "$(1)" " generate swagger dosc for $(a2) service",)
endef

DOCKER_BUILD_TARGETS=$(call gen_targets,docker)
SWAGGER_DOCS_TARGETS=$(call gen_targets,docs)
LOCAL_RUN_TARGETS=$(call gen_targets,local)
BUILD_SERVICES=$(call gen_targets,build)

all_dynamic_targets=$(BUILD_SERVICES) $(LOCAL_RUN_TARGETS) $(SWAGGER_DOCS_TARGETS) $(DOCKER_BUILD_TARGETS)

.PHONY: all build up down core event consumer clean

help: ## prints help content for the makefile
	@echo
	@echo "Available targets:"
	@echo "-------------------"
	@grep -E '^[a-zA-Z._/\-]+:.*?## ' $(MAKEFILE_LIST) | sort | awk -F '[:|##]' '{printf "$(BLUE)%-40s$(NC) %s\n", $$1, $$4}'
	$(foreach target,$(all_dynamic_targets), $(call help_message,$(target)))

all: clean build ## run clean, build steps for all services

up: ## bring the containers up
	docker-compose up -d

down: ## bring the containers down
	docker-compose down

clean: ## remove older build artifacts
	@echo "Removing previous build artifacts..."
	@rm -rf bin


build: $(BUILD_SERVICES) ## Build all services

# $(SERVICES): %: build/%

build/%: ## build the target service
	@mkdir -p bin
	@echo "Building service $* with version $(VERSION)..."

	@CGO_ENABLED=$(CGO_ENABLED) $(GO_BIN) build \
	-ldflags "$(BUILD_VERSION_LD_FLAGS)" \
	-o bin/$* \
	./cmd/$*

docker: $(DOCKER_BUILD_TARGETS) ## Build Docker images for services

$(DOCKER_BUILD_TARGETS): %: run/%

run/docker/%:
	@echo "Building Docker image for service $* with version $(VERSION)..."
	@echo "	LD Flags: $(BUILD_VERSION_LD_FLAGS)"
	@docker build -t go-monorepo/$*:$(VERSION) . \
	--build-arg SERVICE=$* \
	--build-arg BUILD_VERSION_LD_FLAGS="$(BUILD_VERSION_LD_FLAGS)" 
	@docker tag go-monorepo/$*:$(VERSION) go-monorepo/$*:latest

docker/frontend:
	@echo "Building docker image for $(@F)..."
	@docker build -t go-monorepo/$(@F):$(VERSION) --file=$(@F)/Dockerfile $(@F)/
	@docker tag go-monorepo/$(@F):$(VERSION) go-monorepo/$(@F):latest

docs: $(SWAGGER_DOCS_TARGETS) ## Generate Swagger docs

$(SWAGGER_DOCS_TARGETS): %: run/%

run/docs/%:
	@echo "Generating Swagger for $(@D) -> $(@F)... \n"

	$(eval api_path := internal/$(@F))
	$(eval route_paths := $(filter internal/$(@F)/%,$(swagger_api_file_dirs)))
	$(foreach route,$(route_paths),swag init -o $(subst internal/,docs/,$(route)) -d $(route),$(api_path) -g routes.go -pdl 1 --parseInternal;)

local: $(LOCAL_RUN_TARGETS) ## Run services locally with dynamic port assignment

$(LOCAL_RUN_TARGETS): %: run/%

run/local/%:
	@echo "Running service $* locally..."
	@read -p "Enter port for $*: " PORT; \
	echo "Starting $* on port $$PORT..."; \
	go run ./cmd/$*/main.go server --port=$$PORT