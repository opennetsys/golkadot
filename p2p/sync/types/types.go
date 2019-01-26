package synctypes

import (
	"github.com/c3systems/go-substrate/block"
	peertypes "github.com/c3systems/go-substrate/p2p/peer/types"
)

// Requests ...
type Requests []*StateRequest

// StateRequest TODO
type StateRequest struct {
	// Peer ..
	Peer peertypes.InterfacePeer
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
	Peer peertypes.InterfacePeer
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
