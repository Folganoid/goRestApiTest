.PHONY: build
build:
	@echo "  >  Building binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build ./cmd/apiserver

.DEFAULT_GOAL := build

.PHONY: test
test:
	@echo "  >  Test..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -v -race -timeout 30s ./...