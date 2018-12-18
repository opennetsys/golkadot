package rpc

import (
	gorpc "github.com/libp2p/go-libp2p-gorpc"
)

// Server is an RPC server that will respond to calls
type Server struct {
	RPCHost *gorpc.Server
}

// Client is an RPC client that can make calls
type Client struct {
	RPCClient *gorpc.Client
}
