package client

import (
	"github.com/c3systems/c3-go/rpc"
	"github.com/c3systems/go-substrate/p2p"
)

// TODO ...

// Config ...
type Config struct {
	//Chain ChainName,
	//DB DbConfig
	//Dev DevConfig
	P2P   *p2p.Config
	RPC   *rpc.Config
	Roles []string
	//Telemetry TelemetryConfig
	//Wasm WasmConfig
}

// BlockNumber ...
// note: required by sync.provideBlocks
type BlockNumber struct{}

// BlockRequest ...
// note: required by sync.provideBlocks
// see: https://github.com/polkadot-js/client/blob/master/packages/client-types/src/messages/BlockRequest.ts
type BlockRequest struct {
	ID        uint64
	FromValue *math.Big // note: or BlockNumber???
	Max       int
	From      int
	Direction string // note: create enums?
}

// BlockResponse ...
// note: required by sync.provideBlocks
// see: https://github.com/polkadot-js/client/blob/master/packages/client-types/src/messages/BlockResponse.ts
type BlockResponse struct {
	Blocks []interface{} // TODO: change...
	ID     uint64
}
