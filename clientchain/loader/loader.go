package clientchainloader

import (
	"github.com/c3systems/go-substrate/client"
	clientchaintypes "github.com/c3systems/go-substrate/clientchain/types"
)

// TODO: https://github.com/polkadot-js/client/blob/master/packages/client-chains/src/loader.ts

// Loader ...
type Loader struct {
	chain *clientchaintypes.ChainJSON
}

// NewLoader ...
func NewLoader(config *client.Config) *Loader {
	// TODO
	return &Loader{}
}

// Chain ...
func (l *Loader) Chain() *clientchaintypes.ChainJSON {
	return l.chain
}
