package sync

import (
	"context"
	"errors"
	"math/big"

	synctypes "github.com/opennetsys/golkadot/client/p2p/sync/types"
	clienttypes "github.com/opennetsys/golkadot/client/types"
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
	bestQueued    *big.Int
	blockRequests clienttypes.StateBlockRequests
	blockQueue    clienttypes.StateBlockQueue
	chain         clienttypes.InterfaceChains
	config        *clienttypes.ConfigClient
	ctx           context.Context
	handlers      map[synctypes.EventEnum]clienttypes.EventCallback
	BestSeen      *big.Int
	Status        synctypes.StatusEnum
}
