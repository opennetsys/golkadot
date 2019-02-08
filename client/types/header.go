package clienttypes

import (
	"errors"
	"math"
	"math/big"

	pcrypto "github.com/opennetsys/godot/common/crypto"
)

// NewHeader ...
func NewHeader(h *Header, sessionValidators []*AccountID) (*Header, error) {
	hdr := &Header{}

	if h != nil {
		hdr.BlockNumber = h.BlockNumber
		hdr.ParentHash = h.ParentHash
		hdr.Number = h.Number
		hdr.StateRoot = h.StateRoot
		hdr.ExtrinsicsRoot = h.ExtrinsicsRoot
		hdr.Digest = h.Digest
		hdr.Author = h.Author
	}

	if sessionValidators != nil && len(sessionValidators) > 0 {
		if hdr.Digest == nil {
			return nil, errors.New("nil digest")
		}
		if hdr.Digest.Logs == nil {
			return nil, errors.New("nil digest logs")
		}
		digestItem, ok := hdr.Digest.Logs[Seal]
		if !ok {
			return nil, errors.New("no seal item")
		}
		slotItem, ok := digestItem.(SealObj)
		if !ok {
			return nil, errors.New("seal item is not a seal object")
		}

		idx := math.Mod(float64(slotItem.Slot), float64(len(sessionValidators)))
		hdr.Author = sessionValidators[int(idx)]
	}

	return hdr, nil
}

// GetBlockNumber ...
func (h *Header) GetBlockNumber() *big.Int {
	return h.Number
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

// Hash ...
func (h *Header) Hash() *pcrypto.Blake2b256Hash {
	return pcrypto.NewBlake2b256(h.ToU8a())
}

// ToU8a ...
func (h *Header) ToU8a() []uint8 {
	// TODO
	return nil
}
