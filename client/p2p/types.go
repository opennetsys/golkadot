package p2p

import (
	"context"
	"errors"

	p2ptypes "github.com/opennetsys/go-substrate/client/p2p/types"
	clienttypes "github.com/opennetsys/go-substrate/client/types"
)

const (
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
	state     *clienttypes.State
	cfg       *clienttypes.ConfigClient
	dialQueue map[string]*clienttypes.QueuedPeer
	sync      clienttypes.InterfaceSync
	ctx       context.Context
	ch        chan interface{}
	cancel    context.CancelFunc
	handlers  map[p2ptypes.EventEnum]clienttypes.EventCallback
}
