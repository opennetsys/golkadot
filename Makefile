all: build

.PHONY: deps
deps:
	@GO111MODULE=on go mod vendor

.PHONY: build
build:
	@go build

.PHONY: start
start:
	@go run cmd/node/node.go

.PHONY: test
test:
	@go test -v ./... && echo "ALL PASS" || echo "FAILURE"

.PHONY: test/wasm
test/wasm:
	@go test -v wasm/*.go

.PHONY: test/runtime
test/runtime:
	@go test -v runtime/*.go $(ARGS)

.PHONY: test/clientdb
test/clientdb:
	@go test -v clientdb/*.go

.PHONY: test/common
test/common:
	@go test -v $$(go list ./... | grep common/)

.PHONY: test/common/hexutil
test/common/hexutil:
	@go test -v common/hexutil/*.go

.PHONY: test/common/stringutil
test/common/stringutil:
	@go test -v common/stringutil/*.go

.PHONY: test/common/assert
test/common/assert:
	@go test -v common/assert/*.go

.PHONY: test/common/ext
test/common/ext:
	@go test -v common/ext/*.go

.PHONY: test/common/bnutil
test/common/bnutil:
	@go test -v common/bnutil/*.go

.PHONY: test/common/chainspec
test/common/chainspec:
	@go test -v common/chainspec/*.go

.PHONY: test/common/u8compact
test/common/u8compact:
	@go test -v common/u8compact/*.go

.PHONY: test/common/u8util
test/common/u8util:
	@go test -v common/u8util/*.go

.PHONY: test/common/triecodec
test/common/triecodec:
	@go test -v common/triecodec/*.go $(ARGS)

.PHONY: test/common/triehash
test/common/triehash:
	@go test -v common/triehash/*.go $(ARGS)

.PHONY: test/common/triedb
test/common/triedb:
	@go test -v common/triedb/*.go $(ARGS)

.PHONY: test/common/crypto
test/common/crypto:
	@go test -v common/crypto/*.go

.PHONY: test/common/mathutil
test/common/mathutil:
	@go test -v common/mathutil/*.go

.PHONY: test/common/db
test/common/db:
	@go test -v common/db/*.go

.PHONY: test/common/fileflatdb
test/common/fileflatdb:
	@go test -v common/fileflatdb/*.go

.PHONY: test/common/diskdb
test/common/diskdb:
	@go test -v common/diskdb/*.go
