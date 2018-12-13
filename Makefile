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

.PHONY test/common/assert:
test/common/assert:
	@go test -v common/assert/*.go

.PHONY test/common/ext:
test/common/ext:
	@go test -v common/ext/*.go

.PHONY test/common/bn:
test/common/bn:
	@go test -v common/bn/*.go

.PHONY test/common/chainspec:
test/common/chainspec:
	@go test -v common/chainspec/*.go

.PHONY test/common/u8compact:
test/common/u8compact:
	@go test -v common/u8compact/*.go
