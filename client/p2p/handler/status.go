package handler

import (
	"errors"

	handlertypes "github.com/opennetsys/go-substrate/client/p2p/handler/types"
	clienttypes "github.com/opennetsys/go-substrate/client/types"
	"github.com/opennetsys/go-substrate/logger"
)

// note: ensure the struct implements the interface
var _ InterfaceHandler = (*StatusHandler)(nil)

// StatusHandler implements the status handler
type StatusHandler struct{}

// Func handles incoming status messages
// TODO ...
// TODO: We should check the genesisHash here and act appropriately
func (s *StatusHandler) Func(p clienttypes.InterfaceP2P, pr clienttypes.InterfacePeer, msg clienttypes.InterfaceMessage) error {
	if p == nil {
		return errors.New("nil p2p")
	}
	if pr == nil {
		return errors.New("nil peer")
	}
	if msg == nil {
		return errors.New("nil message")
	}

	b, err := msg.MarshalJSON()
	if err != nil {
		logger.Errorf("[handler] err unmarshalling block response message\n%v", err)
		return err
	}

	logger.Infof("%v Status: %v", pr.GetShortID(), string(b))

	st, ok := msg.(*clienttypes.Status)
	if !ok {
		logger.Errorf("[handler] expected pointer to status, received %T", st)
		return errors.New("message is not a status")
	}
	if st == nil || st.Message == nil {
		return errors.New("nil status message")
	}

	return pr.SetBest(st.Message.BestNumber, st.Message.BestHash[:])
}

// Type returns the func enum
func (s *StatusHandler) Type() handlertypes.FuncEnum {
	return handlertypes.Status
}
