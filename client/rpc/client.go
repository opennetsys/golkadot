package rpc

import (
	"errors"

	gorpc "github.com/libp2p/go-libp2p-gorpc"
	libp2pHost "github.com/libp2p/go-libp2p-host"
	protocol "github.com/libp2p/go-libp2p-protocol"
)

// New returns a new rpc client
func New(host libp2pHost.Host, id *protocol.ID) (*gorpc.Client, error) {
	if id == nil {
		return nil, errors.New("protocol id is required")
	}

	return gorpc.NewClient(host, *id), nil
}
