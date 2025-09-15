
# go-monorepo

This is a Go monorepo project with a scalable structure for multiple microservices, Kafka, and Postgres integration.

## Project Structure

```
go-monorepo/
├── README.md
├── go.mod
├── Makefile
├── docker-compose.yml
├── api/                # Shared API definitions (add as needed)
├── cmd/
│   ├── core/           # Core service (static HTML, webhook, event data)
│   ├── event/          # Event producer service (REST API, Kafka producer)
│   └── consumer/       # Kafka consumer service (webhook to core, storage)
├── deployments/
│   └── k8s/            # Kubernetes manifests (add as needed)
├── internal/
│   ├── core/           # Core service logic
│   ├── event/          # Event service logic
│   └── consumer/       # Consumer service logic
└── bin/                # Built binaries (gitignored)
```

## Directory Overview

- **cmd/**: Entry points for each microservice.
- **internal/**: Service-specific business logic and config.
- **api/**: Shared API definitions (e.g., Go types, protobufs).
- **deployments/**: Deployment configurations (e.g., Kubernetes YAML).
- **bin/**: Compiled binaries (not committed).

## Local Development & Usage

### Prerequisites
- Go 1.21+
- Docker & Docker Compose

### Setup & Run

1. Start Kafka and Postgres:
	```sh
	make up
	```
2. Build all services:
	```sh
	make build
	```
3. Run services (in separate terminals):
	```sh
	make core
	make event
	make consumer
	```

### Makefile Targets
- `make up` — Start Kafka and Postgres
- `make down` — Stop all containers
- `make build` — Build all Go services
- `make core` — Run the core service (default port 8080)
- `make event` — Run the event service (default port 8081)
- `make consumer` — Run the consumer service (default port 8082)
- `make clean` — Remove built binaries

### Service Ports
- Core: `:8080`
- Event: `:8081`
- Consumer: `:8082`

### Configuration
- Each service uses an embedded YAML config in its `internal/<service>/config.yaml`.
- Update these files to match your local setup if needed.

### Healthchecks
- (Recommended) Add `/health` endpoints to each service for orchestration and debugging.

## License

MIT
