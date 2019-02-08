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

.PHONY: test/rpc
test/rpc:
	@go test -v rpc/*.go

.PHONY: test/p2p
test/p2p:
	@go test -v p2p/*.go

.PHONY: test/types
test/types:
	@go test -v types/*.go

.PHONY: test/runtime
test/runtime:
	@go test -v runtime/*.go $(ARGS)

.PHONY: test/client
test/client:
	@go test -v client/*.go

.PHONY: test/clientdb
test/clientdb:
	@go test -v client/db/*.go

.PHONY: test/clientchain
test/clientchain:
	@go test -v client/chain/*.go

.PHONY: test/client/chainloader
test/client/chainloader:
	@go test -v client/chains/loader/*.go

.PHONY: test/common/all
test/common/all:
	@go test -v $$(go list ./... | grep common/)

.PHONY: test/common
test/common:
	@go test -v common/common*.go

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

.PHONY: move/repo
move/repo:
	@find . -type f -name '*.go' -not \( -path './.git/*' -o -path './vendor/*' \) -exec sed -i 's|github.com/c3systems/go-substrate/|github.com/opennetsys/go-substrate/|g' {} \;

.PHONY: move/name
move/name:
	@find . -type f -name '*.go' -not \( -path './.git/*' -o -path './vendor/*' \) -exec sed -i 's|github.com/opennetsys/go-substrate/|github.com/opennetsys/godot/|g' {} \;
