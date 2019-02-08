package clientchaintypes

import (
	clienttypes "github.com/opennetsys/golkadot/client/types"
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
	Block  *clienttypes.BlockData
	Code   []uint8
	Header *clienttypes.Header
}
