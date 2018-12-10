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

.PHONY test/common/hexutil:
test/common/hexutil:
	@go test -v common/hexutil/*.go

.PHONY test/common/stringutil:
test/common/stringutil:
	@go test -v common/stringutil/*.go
