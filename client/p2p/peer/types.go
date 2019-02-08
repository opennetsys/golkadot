package peer

import (
	"errors"
	"math/big"

	peertypes "github.com/opennetsys/golkadot/client/p2p/peer/types"
	clienttypes "github.com/opennetsys/golkadot/client/types"

	inet "github.com/libp2p/go-libp2p-net"
	pstore "github.com/libp2p/go-libp2p-peerstore"
)

var (
	// ErrNoConfig ...
	ErrNoConfig = errors.New("no config")
	// ErrNoChain ...
	ErrNoChain = errors.New("no chain")
)

// Peer ...
type Peer struct {
	// TODO: map?!?
	BestHash   []byte
	BestNumber *big.Int
	chain      clienttypes.InterfaceChains
	config     *clienttypes.ConfigClient
	//Config      *clienttypes.ConfigPeer
	id          string
	connections map[uint]inet.Conn
	nextID      uint
	nextConnID  uint
	peerInfo    pstore.PeerInfo
	shortID     string
	handlers    map[peertypes.EventEnum]clienttypes.PeerEventCallback
}
