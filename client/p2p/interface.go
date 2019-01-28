package p2p

import (
	clienttypes "github.com/c3systems/go-substrate/client/types"
)

// InterfaceP2P defines the methods of the p2p service
type InterfaceP2P interface {
	// IsStarted returns true if the p2p interface has started
	IsStarted() bool
	// GetNumPeers returns the number of connected peers
	GetNumPeers() (uint, error)
	// On handles messages
	On(event EventEnum, cb clienttypes.EventCallback) (interface{}, error)
	// Start starts the p2p service
	//Start(ctx context.Context, ch chan interface{}) error
	// Stop stops the p2p service
	Stop() error
	// Cfg returns the config
	Cfg() clienttypes.ConfigP2P
}
