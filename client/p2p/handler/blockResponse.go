package handler

import (
	"errors"

	handlertypes "github.com/opennetsys/godot/client/p2p/handler/types"
	clienttypes "github.com/opennetsys/godot/client/types"
	"github.com/opennetsys/godot/logger"
)

// note: ensure the struct implements the interface
var _ InterfaceHandler = (*BlockResponseHandler)(nil)

// BlockResponseHandler implements the block response handler
type BlockResponseHandler struct{}

// Func handles incoming block response messages
// TODO ...
func (b *BlockResponseHandler) Func(p clienttypes.InterfaceP2P, pr clienttypes.InterfacePeer, msg clienttypes.InterfaceMessage) error {
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

	logger.Infof("%v BlockResponse: %v", pr.GetShortID(), string(byt))

	s, err := p.GetSyncer()
	if err != nil {
		return err
	}
	if s == nil {
		return errors.New("syncer is nil")
	}

	br, ok := msg.(*clienttypes.BlockResponse)
	if !ok {
		logger.Errorf("[handler] expected pointer to block response, received %T", br)
		return errors.New("message is not a block response")
	}
	if err = s.QueueBlocks(pr, br); err != nil {
		logger.Errorf("[handler] err queueing blocks\n%v", err)
		return err
	}

	return s.RequestBlocks(pr)
}

// Type returns the func enum
func (b *BlockResponseHandler) Type() handlertypes.FuncEnum {
	return handlertypes.BlockResponse
}
