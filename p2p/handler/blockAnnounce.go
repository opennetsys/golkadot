package handler

import (
	"errors"

	"github.com/c3systems/go-substrate/logger"
	"github.com/c3systems/go-substrate/p2p"
	"github.com/c3systems/go-substrate/p2p/message"
	"github.com/c3systems/go-substrate/p2p/peer"
)

// note: ensure the struct implements the interface
var _ InterfaceHandler = (*BlockAnnounceHandler)(*nil)

// BlockAnnounceHandler implements the block announce handler
type BlockAnnounceHandler struct{}

// Func handles incoming block announce messages
// TODO ...
func (b *BlockAnnounceHandler) Func(p p2p.Interface, pr peer.Interface, msg message.Interface) error {
	var msgStrBytes []byte
	if err := msg.Unmarshal(msgBytes); err != nil {
		logger.Errorf("[handler] err unmarshalling block response message\n%v", err)
		return err
	}
	logger.Infof("%v BlockAnnounce: %v", pr.Cfg().ShortID, string(msgStrBytes))

	header := msg.Header()

	if pr.Cfg().BestNumber.Cmp(header.BlockNumber) == -1 {
		pr.SetBest(header.BlockNumber, header.Hash)
	}

	s := p.Cfg().Syncer
	if s == nil {
		return errors.New("syncer is nil")
	}
	s.RequestBlocks(peer)

	return nil
}

// Type returns the func enum
func (b *BlockAnnounceHandler) Type() FuncEnum {
	return BlockAnnounce
}
