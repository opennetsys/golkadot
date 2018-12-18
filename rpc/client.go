package client

import (
	"errors"

	libp2p "github.com/libp2p/go-libp2p"
	protocol "github.com/libp2p/go-libp2p-protocol"
)

// New returns a new rpc client
func New(host *libp2p.BasicHost, id protocol.ID) (*Server, error) {
	if host == nil {
		return nil, errors.New("host is required")
	}

	rpcClient := gorpc.NewClient(client, protocolID)

	return &Client{
		RPCClient: rpcClient,
	}, nil
}
