package peer

import (
	"errors"
	"math/big"

	clienttypes "github.com/c3systems/go-substrate/client/types"
	libpeer "github.com/libp2p/go-libp2p-peer"
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
	Map        map[libpeer.ID]*clienttypes.KnownPeer
	BestHash   []byte
	BestNumber *big.Int
	Chain      clienttypes.InterfaceChains
	Config     *clienttypes.ConfigClient
	//Config      *clienttypes.ConfigPeer
	ID          string
	Connections map[int]*clienttypes.Connection
	NextID      uint
	NextConnID  uint
	PeerInfo    pstore.PeerInfo
	ShortID     string
}
