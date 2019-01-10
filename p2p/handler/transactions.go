package handler

import (
	"github.com/c3systems/go-substrate/logger"
	"github.com/c3systems/go-substrate/p2p"
	"github.com/c3systems/go-substrate/p2p/message"
	"github.com/c3systems/go-substrate/p2p/peer"
)

// TransactionsHandler implements the transactions handler
type TransactionsHandler struct{}

// Func handles incoming transactions messages
// TODO ...
// TODO Propagate
func (t *TransactionsHandler) Func(p p2p.Interface, pr peer.Interface, msg message.Interface) error {
	var msgStrBytes []byte
	if err := msg.Unmarshal(msgBytes); err != nil {
		logger.Errorf("[handler] err unmarshalling transaction message\n%v", err)
		return err
	}

	logger.Infof("%v Transaction: %v", pr.Cfg().ShortID, string(msgStrBytes))

	return nil
}

// Type returns the func enum
func (t *TransactionsHandler) Type() FuncEnum {
	return Transactions
}
