package sync

import (
	"errors"
	"math/big"

	"github.com/c3systems/go-substrate/chain"
	"github.com/c3systems/go-substrate/client"
)

const (
	// REPORT_COUNT ...
	REPORT_COUNT int = 5
	// REQUEST_TIMEOUT ...
	// note: ms?
	REQUEST_TIMEOUT int = 60000
)

var (
	// ErrNilConfig ...
	ErrNilConfig = errors.New("nil config")
	// ErrNilChain ...
	ErrNilChain = errors.New("nil chain")
)

// Sync ...
type Sync struct {
	Chain         chain.InterfaceChain
	BlockRequests StateBlockRequests
	BlockQueue    StateBlockQueue
	BestQueued    *big.Int
	BestSeen      *big.Int
	Status        StatusEnum
	Config        *client.Config
}
