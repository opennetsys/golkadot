package peers

import (
	"github.com/c3systems/go-substrate/p2p/peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
)

// Interface defines the methods of the peers
type Interface interface {
	// Add adds a peer to peers
	Add(peerInfo pstore.PeerInfo) (peer.Interface, error)
	// Count returns the number of connected peers
	Count() uint
	// Get returns a peer
	Get(peerInfo pstore.PeerInfo) (*peer.KnownPeer, error)
	// Log TODO
	Log(event EventEnum, p peer.Interface) error
	// On handles peers events
	On(event EventEnum, cb EventCallback) (interface{}, error)
	// Peers returns the peers
	Peers() []peer.Interface
}
