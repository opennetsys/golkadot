package peers

import (
	"github.com/c3systems/go-substrate/p2p/peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
)

// InterfacePeers defines the methods of the peers
type InterfacePeers interface {
	// Add adds a peer to peers
	Add(pi pstore.PeerInfo) (*peer.KnownPeer, error)
	// Count returns the number of connected peers
	Count() (int, error)
	// Get returns a peer
	Get(pi pstore.PeerInfo) (*peer.KnownPeer, error)
	// Log TODO
	Log(event EventEnum, p *peer.KnownPeer) error
	// On handles peers events
	On(event EventEnum, cb EventCallback) (interface{}, error)
	// Peers returns the peers
	Peers() ([]*peer.KnownPeer, error)
}
