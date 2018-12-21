package handler

import (
	"github.com/c3systems/go-substrate/p2p"
	"github.com/c3systems/go-substrate/p2p/message"
	"github.com/c3systems/go-substrate/p2p/peer"
)

// StatusHandler implements the status handler
type StatusHandler struct{}

// Func handles incoming status messages
func (s *StatusHandler) Func(p p2p.Interface, pr peer.Interface, msg message.Interface) error {
	return nil
}

// Type returns the func enum
func (s *StatusHandler) Type() FuncEnum {
	return Status
}
