package peertypes

import (
	"math/big"

	"github.com/c3systems/go-substrate/p2p/message"
	libpeer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	transport "github.com/libp2p/go-libp2p-transport"
)

// Config is passed to New to create a new peer
type Config struct {
	// BestHash TODO
	BestHash []byte
	// BestNumber TODO
	BestNumber *big.Int
	// ID is the peer id
	ID *libpeer.ID
	// PeerInfo is the peer metadata
	PeerInfo *pstore.PeerInfo
	// ShortID TODO
	ShortID string
}

// EventCallback is a function that is called on a peer event
type EventCallback func(msg message.Interface) (interface{}, error)

// KnownPeer is a peer that has been discovered
type KnownPeer struct {
	// Peer is the known peer
	Peer InterfacePeer
	// IsConnected is true if the peer is connected
	IsConnected bool
}

// Connection ...
type Connection struct {
	Connection transport.Conn
	Pushable   chan<- interface{} // note: a write only channel
}
