package messages

import "math/big"

// Started ...
type Started struct {
	BestHash   []uint8
	BestNumber *big.Int
}

// NewStarted ...
func NewStarted(bestHash []uint8, bestNumber *big.Int) *Started {
	return &Started{
		BestHash:   bestHash,
		BestNumber: bestNumber,
	}
}
