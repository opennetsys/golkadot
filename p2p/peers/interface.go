package peers

import (
	peertypes "github.com/c3systems/go-substrate/p2p/peer/types"
	pstore "github.com/libp2p/go-libp2p-peerstore"
)

// InterfacePeers defines the methods of the peers
type InterfacePeers interface {
	// Add adds a peer to peers
	Add(pi pstore.PeerInfo) (*peertypes.KnownPeer, error)
	// Count returns the number of connected peers
	Count() (int, error)
	// Get returns a peer
	Get(pi pstore.PeerInfo) (*peertypes.KnownPeer, error)
	// Log TODO
	Log(event EventEnum, kp *peertypes.KnownPeer) error
	// On handles peers events
	On(event EventEnum, cb EventCallback) (interface{}, error)
	// Peers returns the peers
	KnownPeers() ([]*peertypes.KnownPeer, error)
}
