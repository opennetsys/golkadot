package types

import "github.com/c3systems/go-substrate/common/crypto"

// TODO
// https://github.com/polkadot-js/api/blob/master/packages/types/src/Header.ts

// Header ...
type Header struct {
	Digest         *Digest
	ExtrinsicsRoot *U8a
	Number         *Int
	ParentHash     *U8a
	StateRoot      *U8a
}

// NewHeader ...
func NewHeader() *Header {
	return &Header{}
}

// BlockNumber ...
func (h *Header) BlockNumber() *BlockNumber {
	return NewBlockNumber(h.Number)
}

// SetStateRoot ...
func (h *Header) SetStateRoot(stateRoot *U8a) {
	h.StateRoot = stateRoot
}

// SetExtrinsicsRoot ...
func (h *Header) SetExtrinsicsRoot(root *U8a) {
	h.ExtrinsicsRoot = root
}

// SetParentHash ...
func (h *Header) SetParentHash(hash *U8a) {
	h.ParentHash = hash
}

// Hash ...
func (h *Header) Hash() *Hash {
	return NewHash(crypto.NewBlake2b256(h.ToU8a())[:])
}

// ToU8a ...
func (h *Header) ToU8a() []uint8 {
	// TODO
	return nil
}
