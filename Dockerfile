# Stage 1: Build the Go application
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod tidy
COPY . .

ARG SERVICE
ARG BUILD_VERSION_LD_FLAGS

RUN CGO_ENABLED=0 BUILD_VERSION_LD_FLAGS=${BUILD_VERSION_LD_FLAGS} make ${SERVICE}

# Stage 2: Create a minimal image to run the application
FROM alpine:latest

ARG SERVICE
COPY --from=builder /app/bin/${SERVICE} /bin/service
EXPOSE 80

ENTRYPOINT [ "/bin/service" ]

CMD [ "server", "--host", "0.0.0.0", "--port", "80" ]