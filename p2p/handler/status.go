package handler

import (
	"fmt"

	"github.com/c3systems/go-substrate/logger"
	"github.com/c3systems/go-substrate/p2p"
	"github.com/c3systems/go-substrate/p2p/message"
	"github.com/c3systems/go-substrate/p2p/message/status"
	"github.com/c3systems/go-substrate/p2p/peer"
)

// note: ensure the struct implements the interface
var _ InterfaceHandler = (*StatusHandler)(*nil)

// StatusHandler implements the status handler
type StatusHandler struct{}

// Func handles incoming status messages
// TODO ...
// TODO: We should check the genesisHash here and act appropriately
func (s *StatusHandler) Func(p p2p.Interface, pr peer.Interface, msg message.Interface) error {
	var msgStrBytes []byte
	if err := msg.Unmarshal(msgBytes); err != nil {
		logger.Errorf("[handler] err unmarshalling status message\n%v", err)
		return err
	}

	logger.Infof("%v Status: %v", pr.Cfg().ShortID, string(msgStrBytes))
	statusMessage, ok := msg.(status.Status)
	if !ok {
		err := fmt.Errorf("expected status.Status, received %T", msg)
		logger.Errorf("[handler] err casting message\n%v", err)
		return err
	}

	return pr.SetBest(statusMessage.BestNumber, statusMessag.BestHash)
}

// Type returns the func enum
func (s *StatusHandler) Type() FuncEnum {
	return Status
}
