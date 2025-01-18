SHELL := /bin/bash

GOLANGCI_LINT="golangci-lint"
GORELEASER="goreleaser"

.PHONY: lint
lint:
	$(GOLANGCI_LINT) run

bin: ./cmd ./internal ./pkg
	$(GORELEASER) release --snapshot --clean