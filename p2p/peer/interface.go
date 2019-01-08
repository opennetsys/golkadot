package peer

import (
	"math/big"

	"github.com/c3systems/go-substrate/p2p/message"
	transport "github.com/libp2p/go-libp2p-transport"
)

// Interface defines the methods of peer
type Interface interface {
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
	On(event EventEnum, cb EventCallback) (interface{}, error)
	// Send is used to send the peer a message
	Send(msg message.Interface) (bool, error)
	// SetBest sets a new block
	SetBest(blockNumber *big.Int, hash []byte) error
}
