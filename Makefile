all: build

.PHONY: deps
deps:
	@go mod vendor

.PHONY: build
build:
	@go build

.PHONY test/wasm:
test/wasm:
	@go test -v wasm/*.go
