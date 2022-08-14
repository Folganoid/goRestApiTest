.PHONY: build
build:
	@echo "  >  Building binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build ./cmd/apiserver
.DEFAULT_GOAL := build