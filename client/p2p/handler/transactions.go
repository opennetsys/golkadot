package handler

import (
	"errors"

	handlertypes "github.com/c3systems/go-substrate/client/p2p/handler/types"
	clienttypes "github.com/c3systems/go-substrate/client/types"
	"github.com/c3systems/go-substrate/logger"
)

// note: ensure the struct implements the interface
var _ InterfaceHandler = (*TransactionsHandler)(nil)

// TransactionsHandler implements the transactions handler
type TransactionsHandler struct{}

// Func handles incoming transactions messages
// TODO ...
// TODO Propagate
func (t *TransactionsHandler) Func(p clienttypes.InterfaceP2P, pr clienttypes.InterfacePeer, msg clienttypes.InterfaceMessage) error {
	if p == nil {
		return errors.New("nil p2p")
	}
	if pr == nil {
		return errors.New("nil peer")
	}
	if msg == nil {
		return errors.New("nil message")
	}

	b, err := msg.Marshal()
	if err != nil {
		logger.Errorf("[handler] err unmarshalling block response message\n%v", err)
		return err
	}

	logger.Infof("%v Transaction: %v", pr.GetShortID(), string(b))

	return nil
}

// Type returns the func enum
func (t *TransactionsHandler) Type() handlertypes.FuncEnum {
	return handlertypes.Transactions
}
