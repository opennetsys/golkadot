package peers

import (
	"errors"

	peertypes "github.com/c3systems/go-substrate/p2p/peer/types"
	ic "github.com/libp2p/go-libp2p-crypto"
	libpeer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
)

var (
	// ErrNoStore ...
	// TODO: change to nil store
	ErrNoStore = errors.New("no peer store")
	// ErrNoConfig ...
	// TODO: change to nil config
	ErrNoConfig = errors.New("no config")
	// ErrNoPeerMap ...
	// TODO: change to nil peer map
	ErrNoPeerMap = errors.New("no peer map")
	// ErrNoSuchPeer ...
	ErrNoSuchPeer = errors.New("no such peer")
	// ErrNilEvent ...
	ErrNilEvent = errors.New("nil event")
	// ErrNoChain ...
	ErrNilChain = errors.New("nil chain")
)

// EventCallback is a function that is called on a peer event
type EventCallback func(p interface{}) (interface{}, error)

// Peers ...
type Peers struct {
	Store         pstore.Peerstore
	KnownPeersMap map[pstore.PeerInfo]*peertypes.KnownPeer
}

// Config ...
type Config struct {
	// Priv ..
	Priv ic.PrivKey
	// Pub ...
	Pub ic.PubKey
	// ID ...
	ID libpeer.ID
}
