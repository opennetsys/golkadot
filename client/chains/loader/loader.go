package loader

import (
	clientchaintypes "github.com/c3systems/go-substrate/client/chains/types"
	clienttypes "github.com/c3systems/go-substrate/client/types"
)

// TODO: https://github.com/polkadot-js/client/blob/master/packages/client-chains/src/loader.ts

// Loader ...
type Loader struct {
	chain *clientchaintypes.ChainJSON
}

// NewLoader ...
// TODO: config loader?
func NewLoader(config *clienttypes.ConfigClient) *Loader {
	// TODO
	return &Loader{}
}

// Chain ...
func (l *Loader) Chain() *clientchaintypes.ChainJSON {
	return l.chain
}
