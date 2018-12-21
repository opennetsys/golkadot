package handler

import (
	"github.com/c3systems/go-substrate/p2p"
	"github.com/c3systems/go-substrate/p2p/message"
	"github.com/c3systems/go-substrate/p2p/peer"
)

// BlockRequestHandler implements the block request handler
type BlockRequestHandler struct{}

// Func handles incoming block request messages
func (b *BlockRequestHandler) Func(p p2p.Interface, pr peer.Interface, msg message.Interface) error {
	return nil
}

// Type returns the func enum
func (b *BlockRequestHandler) Type() FuncEnum {
	return BlockRequest
}
