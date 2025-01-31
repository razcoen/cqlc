SHELL=/bin/bash

include ./cassandra.mk

GOLANGCI_LINT=golangci-lint
GORELEASER=goreleaser
GOIMPORTS=goimports

.PHONY: lint
lint:
	$(GOLANGCI_LINT) run

.PHONY: generate-and-format
generate-and-format:
	go generate ./...
	go fmt ./...
	$(GOIMPORTS) -w $(shell find . -type f -name '*.go')
	$(GOLANGCI_LINT) run --fix

bin: ./cmd ./internal ./pkg
	$(GORELEASER) release --snapshot --clean
