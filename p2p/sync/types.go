package sync

import (
	"errors"
	"math/big"

	"github.com/c3systems/go-substrate/block"
	"github.com/c3systems/go-substrate/chain"
	"github.com/c3systems/go-substrate/client"
	"github.com/c3systems/go-substrate/p2p/peer"
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

// Requests ...
type Requests []*StateRequest

// StateRequest TODO
type StateRequest struct {
	// Peer ..
	Peer peer.Interface
	// Request ...
	Request block.Request
	// Timeout ...
	Timeout uint64
}

// StateBlockRequests ...
type StateBlockRequests map[string]*StateRequest

// StateBlock ...
type StateBlock struct {
	// Block ...
	Block block.Data
	// Peer ...
	Peer peer.Interface
}

// StateBlockQueue ...
type StateBlockQueue map[string]*StateBlock

// State TODO
type State struct {
	// BlockRequests ...
	BlockRequests StateBlockRequests
	// BlockQueue ...
	BlockQueue StateBlockQueue
	// Status ...
	Status StatusEnum
}

// EventCallback ...
type EventCallback func() (interface{}, error)

// Sync ...
type Sync struct {
	Chain         chain.Interface
	BlockRequests StateBlockRequests
	BlockQueue    StateBlockQueue
	BestQueued    *big.Int
	BestSeen      *big.Int
	Status        StatusEnum
	Config        *client.Config
}
