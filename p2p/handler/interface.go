package handler

import (
	"github.com/c3systems/go-substrate/p2p"
	"github.com/c3systems/go-substrate/p2p/message"
	"github.com/c3systems/go-substrate/p2p/peer"
)

// Interface describes the methods of the handler package
type Interface interface {
	// Func handles the message
	Func(p p2p.Interface, pr peer.Interface, msg message.Interface) error
	// Type returns the handler type
	Type() FuncEnum
}
