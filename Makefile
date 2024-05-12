GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")

all: build

.PHONY: build
build: 
	$(GO) mod tidy

.PHONY: test
test: 
	IPC_WAIT=0 $(GO) test -bench=. -benchmem

.PHONY: benchdata
benchdata:
	IPC_WAIT=0 $(GO) test -bench . -benchmem ./... | gobenchdata --json bench.json

.PHONY: tools
tools:
	$(GO) install go.bobheadxi.dev/gobenchdata@latest

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: fmt-check
fmt-check:
	@diff=$$($(GOFMT) -d $(GOFILES)); \
  if [ -n "$$diff" ]; then \
    echo "Please run 'make fmt' and commit the result:"; \
    echo "$${diff}"; \
    exit 1; \
  fi;
