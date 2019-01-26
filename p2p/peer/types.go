package peer

import (
	"github.com/c3systems/go-substrate/chain"
	peertypes "github.com/c3systems/go-substrate/p2p/peer/types"
	pstore "github.com/libp2p/go-libp2p-peerstore"
)

// Peer ...
type Peer struct {
	Map        map[pstore.PeerInfo]*KnownPeer
	BestHash   []byte
	BestNumber *math.Big
	Chain      chain.Interface
	Config     *peertypes.Config
	//Config      *client.Config
	ID          string
	Connections map[int]*Connection
	NextID      int
	NextConnID  int
	PeerInfo    pstore.PeerInfo
	ShortID     string
}
