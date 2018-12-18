package rpc

import (
	"github.com/c3systems/go-substrate/rpc/author"
	"github.com/c3systems/go-substrate/rpc/chain"
	"github.com/c3systems/go-substrate/rpc/state"
	"github.com/c3systems/go-substrate/rpc/system"

	libp2pHost "github.com/libp2p/go-libp2p-host"
	protocol "github.com/libp2p/go-libp2p-protocol"
)

// ServerConfig is passed to NewServer
type ServerConfig struct {
	Host          libp2pHost.Host
	SystemService system.ServiceInterface
	StateService  state.ServiceInterface
	ChainService  chain.ServiceInterface
	AuthorService author.ServiceInterface
	ID            *protocol.ID
}
