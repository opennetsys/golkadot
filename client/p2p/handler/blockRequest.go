package handler

import (
	"errors"

	handlertypes "github.com/opennetsys/golkadot/client/p2p/handler/types"
	clienttypes "github.com/opennetsys/golkadot/client/types"
	"github.com/opennetsys/golkadot/logger"
)

// note: ensure the struct implements the interface
var _ InterfaceHandler = (*BlockRequestHandler)(nil)

// BlockRequestHandler implements the block request handler
type BlockRequestHandler struct{}

// Func handles incoming block request messages
// TODO ...
func (b *BlockRequestHandler) Func(p clienttypes.InterfaceP2P, pr clienttypes.InterfacePeer, msg clienttypes.InterfaceMessage) error {
	if p == nil {
		return errors.New("nil p2p")
	}
	if pr == nil {
		return errors.New("nil peer")
	}
	if msg == nil {
		return errors.New("nil message")
	}

	byt, err := msg.MarshalJSON()
	if err != nil {
		logger.Errorf("[handler] err unmarshalling block response message\n%v", err)
		return err
	}
	logger.Infof("%v BlockRequest: %v", pr.GetShortID(), string(byt))

	s, err := p.GetSyncer()
	if err != nil {
		return err
	}
	if s == nil {
		return errors.New("syncer is nil")
	}

	br, ok := msg.(*clienttypes.BlockRequest)
	if !ok {
		logger.Errorf("[handler] expected pointer to block request, received %T", br)
		return errors.New("message is not a block request")
	}
	return s.ProvideBlocks(pr, br)
}

// Type returns the func enum
func (b *BlockRequestHandler) Type() handlertypes.FuncEnum {
	return handlertypes.BlockRequest
}
