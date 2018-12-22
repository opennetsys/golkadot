package p2p

import "context"

// Interface defines the methods of the p2p service
type Interface interface {
	// IsStarted returns true if the p2p interface has started
	IsStarted() bool
	// On handles messages
	On(event EventEnum, cb EventCallback) (interface{}, error)
	// Start starts the p2p service
	Start(ctx context.Context, ch chan interface{}) error
	// Stop stops the p2p service
	Stop(ch chan interface{}) error
}
