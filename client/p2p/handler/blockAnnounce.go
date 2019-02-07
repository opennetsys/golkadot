package handler

import (
	"errors"

	handlertypes "github.com/opennetsys/go-substrate/client/p2p/handler/types"
	clienttypes "github.com/opennetsys/go-substrate/client/types"
	"github.com/opennetsys/go-substrate/logger"
)

// note: ensure the struct implements the interface
var _ InterfaceHandler = (*BlockAnnounceHandler)(nil)

// BlockAnnounceHandler implements the block announce handler
type BlockAnnounceHandler struct{}

// Func handles incoming block announce messages
// TODO ...
func (b *BlockAnnounceHandler) Func(p clienttypes.InterfaceP2P, pr clienttypes.InterfacePeer, msg clienttypes.InterfaceMessage) error {
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
	logger.Infof("%v BlockAnnounce: %v", pr.GetShortID(), string(byt))

	header := msg.GetHeader()

	bn := pr.GetBestNumber()
	if bn == nil {
		return errors.New("[handler] peer best number is nil")
	}
	if bn.Cmp(header.BlockNumber) == -1 {
		if err = pr.SetBest(header.BlockNumber, header.Hash()[:]); err != nil {
			return err
		}
	}

	s, err := p.GetSyncer()
	if err != nil {
		return err
	}
	if s == nil {
		return errors.New("syncer is nil")
	}

	return s.RequestBlocks(pr)
}

// Type returns the func enum
func (b *BlockAnnounceHandler) Type() handlertypes.FuncEnum {
	return handlertypes.BlockAnnounce
}
