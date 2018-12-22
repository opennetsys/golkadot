package handler

import (
	"github.com/c3systems/go-substrate/p2p"
	"github.com/c3systems/go-substrate/p2p/message"
	"github.com/c3systems/go-substrate/p2p/peer"
)

// BlockResponseHandler implements the block response handler
type BlockResponseHandler struct{}

// Func handles incoming block response messages
// TODO ...
func (b *BlockResponseHandler) Func(p p2p.Interface, pr peer.Interface, msg message.Interface) error {
	return nil
}

// Type returns the func enum
func (b *BlockResponseHandler) Type() FuncEnum {
	return BlockResponse
}
