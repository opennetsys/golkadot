package messages

import (
	"math/big"

	synctypes "github.com/opennetsys/golkadot/client/p2p/sync/types"
)

// Interval ...
type Interval struct {
	BestHash   []uint8
	BestNumber *big.Int
	Peers      int
	Status     synctypes.StatusEnum
}

// NewInterval ...
func NewInterval(bestHash []uint8, bestNumber *big.Int, peers int, status synctypes.StatusEnum) *Interval {
	return &Interval{
		BestHash:   bestHash,
		BestNumber: bestNumber,
		Peers:      peers,
		Status:     status,
	}
}
