package p2ptypes

import (
	"context"

	"github.com/c3systems/go-substrate/chain"
	"github.com/c3systems/go-substrate/p2p/peers"
	p2psync "github.com/c3systems/go-substrate/p2p/sync"
	ic "github.com/libp2p/go-libp2p-crypto"
	libp2pHost "github.com/libp2p/go-libp2p-host"
)

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
	Addres string
	// Port is the port on which the node is communicating
	Port uint
}

// Config is used to configure a new node
type Config struct {
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
	Syncer p2psync.Interface
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
	Chain chain.Interface
	// Config is the current configuration of the node
	Config *Config
	// Host is the libp2p host
	Host libp2pHost.Host
	// Peers are the connected peers
	Peers peers.Interface
	// Sync is the current sync state
	SyncState *p2psync.State
}
