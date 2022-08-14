.PHONY: build
build:
	@echo "  >  Building binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build ./cmd/apiserver

.PHONY: test
test:
	@echo "  >  Test..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -v -race -timeout 30s ./...


.DEFAULT_GOAL := build