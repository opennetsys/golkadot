package messages

import "math/big"

// BlockImport ...
type BlockImport struct {
	BestHash   []uint8
	BestNumber *big.Int
}

// NewBlockImport ...
func NewBlockImport(bestHash []uint8, bestNumber *big.Int) *BlockImport {
	return &BlockImport{
		BestHash:   bestHash,
		BestNumber: bestNumber,
	}
}

// ToJSON ...
func (b *BlockImport) ToJSON() {
	/*
	   return {
	     ...super.toJSON(),
	     // NOTE the endpoint expects non-prefixed values, as much as I hate doing it...
	     best: u8aToHex(this.bestHash).slice(2),
	     height: this.bestNumber.toNumber()
	   };
	*/
}
