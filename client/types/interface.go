package clienttypes

import (
	"math/big"

	peertypes "github.com/c3systems/go-substrate/client/p2p/peer/types"
	peerstypes "github.com/c3systems/go-substrate/client/p2p/peers/types"
	synctypes "github.com/c3systems/go-substrate/client/p2p/sync/types"
	"github.com/c3systems/go-substrate/common/crypto"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	transport "github.com/libp2p/go-libp2p-transport"
)

// InterfaceSync defines the methods of the sync service
type InterfaceSync interface {
	// On handles events
	On(event synctypes.EventEnum, cb EventCallback) (interface{}, error)
	// QueueBlocks ...
	// TODO ...
	QueueBlocks(pr InterfacePeer, response *BlockResponse) error
	// RequestBlocks ...
	// TODO ...
	RequestBlocks(pr InterfacePeer) error
	// ProvideBlocks ...
	// TODO ...
	ProvideBlocks()
	// PeerRequests ...
	PeerRequests(pr InterfacePeer) (Requests, error)
}

// InterfaceChains describes the methods of the chains service
type InterfaceChains interface {
	// note: required from p2p.peer.AddConnection
	GetBestBlocksNumber() (*big.Int, error)
	GetBestBlocksHash() (*crypto.Blake2b256Hash, error)
	GetGenesisHash() (*crypto.Blake2b256Hash, error)
	// note: required by sync.processBlock
	ImportBlock(block *StateBlock) (bool, error)
	// note required by sync.QueuBlocks
	GetBlockDataByHash(hash *crypto.Blake2b256Hash) (*StateBlock, error)
}

// InterfacePeer defines the methods of peer
type InterfacePeer interface {
	// AddConnection is used to add a connection
	AddConnection(conn transport.Conn, isWritable bool) (uint, error)
	// Disconnect disconnects from the peer
	Disconnect() error
	// IsActive returns whether the peer is active or not
	IsActive() (bool, error)
	// IsWritable returns whether the peer is writable or not
	IsWritable() (bool, error)
	// GetNextID TODO
	GetNextID() (uint, error)
	// On defines the event handlers
	On(event peertypes.EventEnum, cb PeerEventCallback) (interface{}, error)
	// Send is used to send the peer a message
	Send(msg InterfaceMessage) (bool, error)
	// SetBest sets a new block
	SetBest(blockNumber *big.Int, hash []byte) error
	// Cfg returns the peer config
	Cfg() ConfigClient
	// GetID ...
	GetID() string
}

// InterfacePeers defines the methods of the peers
type InterfacePeers interface {
	// Add adds a peer to peers
	Add(pi pstore.PeerInfo) (*KnownPeer, error)
	// Count returns the number of connected peers
	Count() (uint, error)
	// Get returns a peer
	Get(pi pstore.PeerInfo) (*KnownPeer, error)
	// Log TODO
	Log(event peerstypes.EventEnum, kp *KnownPeer) error
	// On handles peers events
	On(event peerstypes.EventEnum, cb peerstypes.EventCallback) (interface{}, error)
	// KnownPeers returns the peers
	KnownPeers() ([]*KnownPeer, error)
}

// InterfaceMessage defines the methods of Message
type InterfaceMessage interface {
	// Kind returns the message's kind
	Kind() uint
	// Encode serializes the message into a bytes array
	Encode() ([]byte, error)
	// Decode deserializes a bytes array into a message
	Decode(bytes []byte) error
	// Marshal returns json
	Marshal() ([]byte, error)
	// Unmarshal converts json to a message
	Unmarshal(bytes []byte) error
	// Header ...
	Header() *Header
}

// InterfaceTelemetry ...
type InterfaceTelemetry interface{}

// InterfaceRPC ...
type InterfaceRPC interface{}
