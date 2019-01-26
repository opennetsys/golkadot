package p2p

import (
	"errors"

	p2ptypes "github.com/c3systems/go-substrate/p2p/types"
)

const (
	// Name is the version name.
	Name string = "dot"
	// Version is the semvar version of the build.
	Version string = "0.0.0"
)

var (
	// ErrNoPeersService ...
	ErrNoPeersService = errors.New("the p2p service has no peers service")
	// ErrUninitializedService ...
	ErrUninitializedService = errors.New("the p2p service has not been initialized")
	// ErrNoConfig ...
	ErrNoConfig = errors.New("a config is required")
	// ErrNoChainService ...
	ErrNoChainService = errors.New("a chain service is required")
	// ErrNoHost ...
	ErrNoHost = errors.New("the p2p service has no host")
)

// P2P implements the p2p interface
type P2P struct {
	state  *p2ptypes.State
	Config *p2ptypes.Config
}
