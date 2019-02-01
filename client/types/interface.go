package clienttypes

import (
	"math/big"

	handlertypes "github.com/c3systems/go-substrate/client/p2p/handler/types"
	peertypes "github.com/c3systems/go-substrate/client/p2p/peer/types"
	peerstypes "github.com/c3systems/go-substrate/client/p2p/peers/types"
	synctypes "github.com/c3systems/go-substrate/client/p2p/sync/types"
	p2ptypes "github.com/c3systems/go-substrate/client/p2p/types"
	"github.com/c3systems/go-substrate/common/crypto"

	inet "github.com/libp2p/go-libp2p-net"
	libpeer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
)

// InterfaceSync defines the methods of the sync service
type InterfaceSync interface {
	// On handles events
	On(event synctypes.EventEnum, cb EventCallback)
	// QueueBlocks ...
	// TODO ...
	QueueBlocks(pr InterfacePeer, response *BlockResponse) error
	// RequestBlocks ...
	// TODO ...
	RequestBlocks(pr InterfacePeer) error
	// ProvideBlocks ...
	ProvideBlocks(pr InterfacePeer, request *BlockRequest) error
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
	AddConnection(conn inet.Conn, isWritable bool) (uint, error)
	// Disconnect disconnects from the peer
	Disconnect() error
	// IsActive returns whether the peer is active or not
	IsActive() (bool, error)
	// IsWritable returns whether the peer is writable or not
	IsWritable() (bool, error)
	// On defines the event handlers
	On(event peertypes.EventEnum, cb PeerEventCallback)
	// Send is used to send the peer a message
	Send(msg InterfaceMessage) (bool, error)
	// SetBest sets a new block
	SetBest(blockNumber *big.Int, hash []byte) error
	// Cfg returns the peer config
	Cfg() ConfigClient
	// GetChain ...
	GetChain() (InterfaceChains, error)
	// GetID ...
	GetID() string
	// GetNextID ...
	GetNextID() uint
	// GetPeerInfo ...
	GetPeerInfo() pstore.PeerInfo
	// GetShortID ...
	GetShortID() string
	// Receive ...
	Receive(stream inet.Stream) error
	// GetBestNumber ...
	GetBestNumber() *big.Int
}

// InterfacePeers defines the methods of the peers
type InterfacePeers interface {
	// Add adds a peer to peers
	Add(pi pstore.PeerInfo) (*KnownPeer, error)
	// CountAll returns the number of known peers
	CountAll() (uint, error)
	// Count returns the number of connected peers
	Count() (uint, error)
	// Get returns a peer
	Get(pi pstore.PeerInfo) (*KnownPeer, error)
	// GetFromID returns a peer
	GetFromID(id libpeer.ID) (*KnownPeer, error)
	// Log TODO
	Log(event peerstypes.EventEnum, iface interface{}) error
	// On handles peers events
	On(event peerstypes.EventEnum, cb PeersEventCallback)
	// KnownPeers returns the peers
	KnownPeers() ([]*KnownPeer, error)
}

// InterfaceMessage defines the methods of Message
type InterfaceMessage interface {
	// Kind returns the message's kind
	Kind() handlertypes.FuncEnum
	// Encode serializes the message into a bytes array
	Encode() ([]byte, error)
	// Decode deserializes a bytes array into a message
	Decode(bytes []byte) error
	// Marshal returns json
	// TODO: change to MarshalJSON
	Marshal() ([]byte, error)
	// Unmarshal converts json to a message
	// TODO: change to UnmarshalJSON
	Unmarshal(bytes []byte) error
	// Header ...
	Header() *Header
}

// InterfaceTelemetry ...
type InterfaceTelemetry interface{}

// InterfaceRPC ...
type InterfaceRPC interface{}

// InterfaceP2P defines the methods of the p2p service
type InterfaceP2P interface {
	// IsStarted returns true if the p2p interface has started
	IsStarted() bool
	// GetNumPeers returns the number of connected peers
	GetNumPeers() (uint, error)
	// On handles messages
	On(event p2ptypes.EventEnum, cb EventCallback)
	// Start starts the p2p service
	Start() error
	// Stop stops the p2p service
	Stop() error
	// Cfg returns the config
	Cfg() ConfigClient
	// GetSyncer ...
	GetSyncer() (InterfaceSync, error)
}
