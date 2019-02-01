package clienttypes

import (
	"context"
	"math/big"
	"time"

	clientdbtypes "github.com/c3systems/go-substrate/client/db/types"
	synctypes "github.com/c3systems/go-substrate/client/p2p/sync/types"
	"github.com/c3systems/go-substrate/client/rpc/author"
	"github.com/c3systems/go-substrate/client/rpc/chain"
	"github.com/c3systems/go-substrate/client/rpc/state"
	"github.com/c3systems/go-substrate/client/rpc/system"
	pcrypto "github.com/c3systems/go-substrate/common/crypto"

	ic "github.com/libp2p/go-libp2p-crypto"
	libp2pHost "github.com/libp2p/go-libp2p-host"
	libpeer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	protocol "github.com/libp2p/go-libp2p-protocol"
)

// TODO ...

// ChainName ...
type ChainName struct{}

// DevConfig ...
type DevConfig struct {
	genBlocks bool
}

// RolesConfig ...
type RolesConfig struct{}

// TelemetryConfig ...
type TelemetryConfig struct{}

// WasmConfig ...
type WasmConfig struct{}

// ConfigClient ...
type ConfigClient struct {
	// TODO: types
	Chain     *ChainName
	DB        *clientdbtypes.InterfaceDBConfig
	Dev       *DevConfig
	P2P       *ConfigP2P
	RPC       *ConfigRPC
	Peer      *ConfigPeer
	Peers     *ConfigPeers
	Roles     []string
	Telemetry *TelemetryConfig
	Wasm      *WasmConfig
}

// Nodes is a list of p2p nodes
type Nodes []string

// RawMessage is a message struct that is sent between nodes
type RawMessage struct {
	// Message is a json object
	Message map[string]interface{}
	// Type defines the message type
	Type uint
}

// Node contains information about a peer node
type Node struct {
	// Address is the address of the node
	Address string
	// Port is the port on which the node is communicating
	Port uint
}

// ConfigP2P is used to configure a new node
type ConfigP2P struct {
	// Address is the address of the new node
	Address string
	// ClientID is the id of the node
	ClientID string
	// MaxPeers is the maximum number of connected peers to accept
	MaxPeers uint
	// Nodes are the nodes that are connected
	Nodes Nodes
	// NoBootNodes defines whether the node should connect to others on boot
	// TODO: re-write this
	NoBootNodes bool
	// Port is the port of the new node
	Port uint
	// Syncer is used to sync the node
	Syncer InterfaceSync
	// Priv ..
	Priv ic.PrivKey
	// Pub ...
	Pub ic.PubKey
	// Context ...
	Context context.Context
}

// EventCallback is called after an event has been received
type EventCallback func() (interface{}, error)

// State is the current node state
type State struct {
	// Chain TODO
	Chain InterfaceChains
	// Config is the current configuration of the node
	Config *ConfigClient
	// Host is the libp2p host
	Host libp2pHost.Host
	// Peers are the connected peers
	Peers InterfacePeers
	// Sync is the current sync state
	SyncState *SyncState
}

// Requests ...
type Requests []*StateRequest

// StateRequest TODO
type StateRequest struct {
	// ID ...
	// TODO: big.Int?
	ID uint64
	// Peer ..
	Peer InterfacePeer
	// Request ...
	Request *BlockRequest
	// Timeout ...
	Timeout int64
}

// StateBlockRequests ...
type StateBlockRequests map[string]*StateRequest

// StateBlock ...
type StateBlock struct {
	// Block ...
	Block *BlockData
	// Peer ...
	Peer InterfacePeer
}

// StateBlockQueue ...
type StateBlockQueue map[string]*StateBlock

// SyncState TODO
type SyncState struct {
	// BlockRequests ...
	BlockRequests StateBlockRequests
	// BlockQueue ...
	BlockQueue StateBlockQueue
	// Status ...
	Status synctypes.StatusEnum
}

// ConfigPeer is passed to New to create a new peer
type ConfigPeer struct {
	// BestHash TODO
	BestHash []byte
	// BestNumber TODO
	BestNumber *big.Int
	// ID is the peer id
	ID libpeer.ID
	// PeerInfo is the peer metadata
	PeerInfo *pstore.PeerInfo
	// ShortID TODO
	ShortID string
}

// PeerEventCallback is a function that is called on a peer event
type PeerEventCallback func(iface interface{}) (interface{}, error)

// KnownPeer is a peer that has been discovered
type KnownPeer struct {
	// Peer is the known peer
	Peer InterfacePeer
	// IsConnected is true if the peer is connected
	IsConnected bool
}

//// Connection ...
//type Connection struct {
//Connection inet.Conn
//Pushable   chan<- []byte // note: a write only channel
//}

// BlockNumber ...
// note: required by sync.provideBlocks
type BlockNumber struct {
	value uint64
}

// BlockRequest ...
// note: required by sync.provideBlocks
// see: https://github.com/polkadot-js/client/blob/master/packages/client-types/src/messages/BlockRequest.ts
type BlockRequest struct {
	ID        uint64
	FromValue *big.Int // TODO: or BlockNumber???
	Max       int
	From      int
	Direction string // note: create enums?
}

// BlockResponse ...
// note: required by sync.provideBlocks
// see: https://github.com/polkadot-js/client/blob/master/packages/client-types/src/messages/BlockResponse.ts
type BlockResponse struct {
	Blocks []*StateBlock
	// TODO: big.Int?
	ID uint64
}

// BFT ...
type BFT struct{}

// BlockAnnounce ...
type BlockAnnounce struct{}

// Transactions ...
type Transactions struct{}

// ConfigPeers ...
type ConfigPeers struct {
	// Priv ..
	Priv ic.PrivKey
	// Pub ...
	Pub ic.PubKey
	// ID ...
	ID libpeer.ID
}

// ConfigRPC is passed to NewServer
type ConfigRPC struct {
	Host          libp2pHost.Host
	SystemService system.ServiceInterface
	StateService  state.ServiceInterface
	// TODO: What's the diff between rpc chain and p2p chain?
	ChainService  chain.ServiceInterface
	AuthorService author.ServiceInterface
	ID            *protocol.ID
}

// AccountID ...
type AccountID [32]uint8

// AuthorityID ...
type AuthorityID AccountID

// AuthoritiesChangeObj ...
// note: obj suffix is required so as to not interfere with the enum
type AuthoritiesChangeObj []*AuthorityID

// ChangesTrieRootObj ...
// note: obj suffix is required so as to not interfere with the enum
type ChangesTrieRootObj pcrypto.Hash

// SealObj ...
// note: obj suffix is required so as to not interfere with the enum
type SealObj struct {
}

// OtherObj ...
// note: obj suffix is required so as to not interfere with the enum
type OtherObj []byte

// DigestItem ..
// TODO: map[DigestEnum]Digest ???
type DigestItem map[DigestEnum]interface{}

// Digest ...
type Digest struct {
	Logs []*DigestItem
}

// Header ...
type Header struct {
	BlockNumber    *big.Int
	Hash           *pcrypto.Blake2b256Hash
	ParentHash     *pcrypto.Blake2b256Hash
	Number         *big.Int
	StateRoot      *pcrypto.Blake2b256Hash
	ExtrinsicsRoot *pcrypto.Blake2b256Hash
	Digest         *Digest
}

//// Request TODO
//type Request struct {
//From *big.Int
//ID   uint
//Max  uint64
//}

// BlockData TODO
// TODO: https://github.com/polkadot-js/client/blob/master/packages/client-types/src/BlockData.ts
type BlockData struct {
	Hash          *pcrypto.Blake2b256Hash
	Header        *Header
	Body          []byte
	Receipt       []byte
	MessageQueue  []byte
	Justification []byte
}

// Status ...
// TODO: this needs to implement the message interface
// https://github.com/polkadot-js/client/blob/master/packages/client-types/src/messages/Status.ts
type Status struct {
	Roles       []string
	BestNumber  *big.Int
	BestHash    *pcrypto.Blake2b256Hash
	GenesisHash *pcrypto.Blake2b256Hash
	Version     string
}

// QueuedPeer ...
type QueuedPeer struct {
	Peer     InterfacePeer
	NextDial time.Time
}

// OnMessage ...
type OnMessage struct {
	Peer    InterfacePeer
	Message InterfaceMessage
}

// PeersEventCallback ...
type PeersEventCallback func(iface interface{}) (interface{}, error)
