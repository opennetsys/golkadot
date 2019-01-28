package types

import (
	clienttypes "github.com/c3systems/go-substrate/client/types"
	"github.com/c3systems/go-substrate/types"
)

// ChainJSON ...
// TODO
type ChainJSON struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	BootNodes []string         `json:"bootNodes"`
	Genesis   ChainJSONGenesis `json:"genesis"`
}

// ChainJSONGenesis ...
type ChainJSONGenesis struct {
	Raw ChainJSONGenesisRaw `json:"raw"`
}

// ChainJSONGenesisRaw ...
type ChainJSONGenesisRaw map[string]string

// ChainGenesis ...
type ChainGenesis struct {
	Block  *clienttypes.Data
	Code   []uint8
	Header *types.Header
}
