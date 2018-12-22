package sync

import (
	"github.com/c3systems/go-substrate/block"
	"github.com/c3systems/go-substrate/p2p/peer"
)

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
