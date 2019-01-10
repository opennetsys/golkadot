package handler

import (
	"github.com/c3systems/go-substrate/logger"
	"github.com/c3systems/go-substrate/p2p"
	"github.com/c3systems/go-substrate/p2p/message"
	"github.com/c3systems/go-substrate/p2p/peer"
)

// BFTHandler implements the bft handler
type BFTHandler struct{}

// Func handles incoming bft messages
// TODO ...
// TODO Propagate
func (b *BFTHandler) Func(p p2p.Interface, pr peer.Interface, msg message.Interface) error {
	var msgStrBytes []byte
	if err := msg.Unmarshal(msgBytes); err != nil {
		logger.Errorf("[handler] err unmarshalling block response message\n%v", err)
		return err
	}
	logger.Infof("%v BFT: %v", pr.Cfg().ShortID, string(msgStrBytes))

	return nil
}

// Type returns the func enum
func (b *BFTHandler) Type() FuncEnum {
	return BFT
}
