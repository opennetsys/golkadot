package clienttypes

import (
	types "github.com/c3systems/go-substrate/types"
)

// TODO: https://github.com/polkadot-js/client/blob/master/packages/client-types/src/BlockData.ts

// BlockData ...
type BlockData struct {
	Hash   *types.Hash
	Header *types.Header
}

// NewBlockData ...
func NewBlockData(input interface{}) *BlockData {
	return &BlockData{}
}

// ToU8a ...
func (b *BlockData) ToU8a() []uint8 {
	// TODO
	return nil
}
