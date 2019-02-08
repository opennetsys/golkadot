package rpc

import (
	"errors"

	clienttypes "github.com/opennetsys/golkadot/client/types"
	"github.com/opennetsys/golkadot/logger"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
)

// NewServer returns a new rpc server
func NewServer(config clienttypes.ConfigRPC) (*gorpc.Server, error) {
	if config.ID == nil {
		return nil, errors.New("protocol id is required")
	}

	rpcHost := gorpc.NewServer(config.Host, *config.ID)

	if err := rpcHost.Register(config.SystemService); err != nil {
		logger.Errorf("[rpc] err registering the system service\n%v", err)
		return nil, err
	}

	if err := rpcHost.Register(config.StateService); err != nil {
		logger.Errorf("[rpc] err registering the state service\n%v", err)
		return nil, err
	}

	if err := rpcHost.Register(config.ChainService); err != nil {
		logger.Errorf("[rpc] err registering the chain service\n%v", err)
		return nil, err
	}

	if err := rpcHost.Register(config.AuthorService); err != nil {
		logger.Errorf("[rpc] err registering the author service\n%v", err)
		return nil, err
	}

	return rpcHost, nil
}
