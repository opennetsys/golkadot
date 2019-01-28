package sync

import (
	"errors"
	"math/big"

	synctypes "github.com/c3systems/go-substrate/client/p2p/sync/types"
	clienttypes "github.com/c3systems/go-substrate/client/types"
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
	Chain         clienttypes.InterfaceChains
	BlockRequests clienttypes.StateBlockRequests
	BlockQueue    clienttypes.StateBlockQueue
	BestQueued    *big.Int
	BestSeen      *big.Int
	Status        synctypes.StatusEnum
	Config        *clienttypes.ConfigClient
}
