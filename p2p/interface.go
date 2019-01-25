package p2p

// InterfaceP2P defines the methods of the p2p service
type InterfaceP2P interface {
	// IsStarted returns true if the p2p interface has started
	IsStarted() bool
	// GetNumPeers returns the number of connected peers
	GetNumPeers() (uint, error)
	// On handles messages
	On(event EventEnum, cb EventCallback) (interface{}, error)
	// Start starts the p2p service
	//Start(ctx context.Context, ch chan interface{}) error
	// Stop stops the p2p service
	Stop() error
	// Cfg returns the config
	Cfg() Config
}
