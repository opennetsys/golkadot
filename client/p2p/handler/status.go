package handler

import (
	"github.com/c3systems/go-substrate/client/p2p"
	clienttypes "github.com/c3systems/go-substrate/client/types"
)

// note: ensure the struct implements the interface
var _ InterfaceHandler = (*StatusHandler)(nil)

// StatusHandler implements the status handler
type StatusHandler struct{}

// Func handles incoming status messages
// TODO ...
// TODO: We should check the genesisHash here and act appropriately
func (s *StatusHandler) Func(p p2p.InterfaceP2P, pr clienttypes.InterfacePeer, msg clienttypes.InterfaceMessage) error {
	//var msgStrBytes []byte
	//if err := msg.Unmarshal(msgBytes); err != nil {
	//logger.Errorf("[handler] err unmarshalling status message\n%v", err)
	//return err
	//}

	//logger.Infof("%v Status: %v", pr.Cfg().ShortID, string(msgStrBytes))
	//statusMessage, ok := msg.(clienttypes.Status)
	//if !ok {
	//err := fmt.Errorf("expected Status, received %T", msg)
	//logger.Errorf("[handler] err casting message\n%v", err)
	//return err
	//}

	//return pr.SetBest(statusMessage.BestNumber, statusMessag.BestHash)
	return nil
}

// Type returns the func enum
func (s *StatusHandler) Type() FuncEnum {
	return Status
}
