#!/bin/bash
set -e

echo "Building service1..."
go build -o bin/service1 ./cmd/service1

echo "Building service2..."
go build -o bin/service2 ./cmd/service2

echo "Build complete."
