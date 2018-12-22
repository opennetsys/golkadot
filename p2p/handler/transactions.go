package handler

import (
	"github.com/c3systems/go-substrate/p2p"
	"github.com/c3systems/go-substrate/p2p/message"
	"github.com/c3systems/go-substrate/p2p/peer"
)

// TransactionsHandler implements the transactions handler
type TransactionsHandler struct{}

// Func handles incoming transactions messages
// TODO ...
func (t *TransactionsHandler) Func(p p2p.Interface, pr peer.Interface, msg message.Interface) error {
	return nil
}

// Type returns the func enum
func (t *TransactionsHandler) Type() FuncEnum {
	return Transactions
}
