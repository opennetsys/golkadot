package handler

import (
	"github.com/c3systems/go-substrate/p2p"
	"github.com/c3systems/go-substrate/p2p/message"
	"github.com/c3systems/go-substrate/p2p/peer"
)

// BFTHandler implements the bft handler
type BFTHandler struct{}

// Func handles incoming bft messages
// TODO ...
func (b *BFTHandler) Func(p p2p.Interface, pr peer.Interface, msg message.Interface) error {
	return nil
}

// Type returns the func enum
func (b *BFTHandler) Type() FuncEnum {
	return BFT
}
