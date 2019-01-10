package handler

import (
	"errors"

	"github.com/c3systems/go-substrate/logger"
	"github.com/c3systems/go-substrate/p2p"
	"github.com/c3systems/go-substrate/p2p/message"
	"github.com/c3systems/go-substrate/p2p/peer"
)

// BlockRequestHandler implements the block request handler
type BlockRequestHandler struct{}

// Func handles incoming block request messages
// TODO ...
func (b *BlockRequestHandler) Func(p p2p.Interface, pr peer.Interface, msg message.Interface) error {
	var msgStrBytes []byte
	if err := msg.Unmarshal(msgBytes); err != nil {
		logger.Errorf("[handler] err unmarshalling block response message\n%v", err)
		return err
	}
	logger.Infof("%v BlockRequest: %v", pr.Cfg().ShortID, string(msgStrBytes))

	s := p.Cfg().Syncer
	if s == nil {
		return errors.New("syncer is nil")
	}

	s.ProvideBlocks(pr, msg)

	return nil
}

// Type returns the func enum
func (b *BlockRequestHandler) Type() FuncEnum {
	return BlockRequest
}
