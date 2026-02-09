# Makefile - helper targets

.PHONY: generate build run

# regenerate swagger docs using go:generate in cmd/server
generate:
	go generate ./cmd/server

build:
	go build ./...

run:
	go run ./cmd/server
