package clienttypes

import pcrypto "github.com/opennetsys/go-substrate/common/crypto"

// TODO
// https://github.com/polkadot-js/api/blob/master/packages/types/src/Header.ts

// NewHeader ...
func NewHeader() *Header {
	return &Header{}
}

// GetBlockNumber ...
func (h *Header) GetBlockNumber() *BlockNumber {
	return NewBlockNumber(h.Number)
}

// SetStateRoot ...
func (h *Header) SetStateRoot(stateRoot *pcrypto.Blake2b256Hash) {
	h.StateRoot = stateRoot
}

// SetExtrinsicsRoot ...
func (h *Header) SetExtrinsicsRoot(root *pcrypto.Blake2b256Hash) {
	h.ExtrinsicsRoot = root
}

// SetParentHash ...
func (h *Header) SetParentHash(hash *pcrypto.Blake2b256Hash) {
	h.ParentHash = hash
}

// GetHash ...
func (h *Header) GetHash() *pcrypto.Blake2b256Hash {
	return h.Hash
}

// ToU8a ...
func (h *Header) ToU8a() []uint8 {
	// TODO
	return nil
}
