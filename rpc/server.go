package rpc

import (
	"errors"
	"log"

	"github.com/c3systems/go-substrate/rpc/author"
	"github.com/c3systems/go-substrate/rpc/chain"
	"github.com/c3systems/go-substrate/rpc/state"
	"github.com/c3systems/go-substrate/rpc/system"

	libp2p "github.com/libp2p/go-libp2p"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	protocol "github.com/libp2p/go-libp2p-protocol"
)

// New returns a new rpc server
func New(host *libp2p.BasicHost, id protocol.ID) (*Server, error) {
	if host == nil {
		return nil, errors.New("host is required")
	}

	rpcHost := gorpc.NewServer(host, protocolID)

	systemSVC := system.Service{}
	if err := rpcHost.Register(&systemSVC); err != nil {
		log.Printf("[rpc] err registering the system service\n%v", err)
		return nil, err
	}

	stateSVC := state.Service{}
	if err := rpcHost.Register(&stateSVC); err != nil {
		log.Printf("[rpc] err registering the state service\n%v", err)
		return nil, err
	}

	chainSVC := chain.Service{}
	if err := rpcHost.Register(&chainSVC); err != nil {
		log.Printf("[rpc] err registering the chain service\n%v", err)
		return nil, err
	}

	authorSVC := author.Service{}
	if err := rpcHost.Register(&authorSVC); err != nil {
		log.Printf("[rpc] err registering the author service\n%v", err)
		return nil, err
	}

	return &Server{
		RPCHost: rpcHost,
	}
}
