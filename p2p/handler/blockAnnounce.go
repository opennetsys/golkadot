package handler

import (
	"github.com/c3systems/go-substrate/p2p"
	"github.com/c3systems/go-substrate/p2p/message"
	"github.com/c3systems/go-substrate/p2p/peer"
)

// BlockAnnounceHandler implements the block announce handler
type BlockAnnounceHandler struct{}

// Func handles incoming block announce messages
// TODO ...
func (b *BlockAnnounceHandler) Func(p p2p.Interface, pr peer.Interface, msg message.Interface) error {
	return nil
}

// Type returns the func enum
func (b *BlockAnnounceHandler) Type() FuncEnum {
	return BlockAnnounce
}
