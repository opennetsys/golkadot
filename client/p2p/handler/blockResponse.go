package handler

import (
	handlertypes "github.com/opennetsys/go-substrate/client/p2p/handler/types"
	clienttypes "github.com/opennetsys/go-substrate/client/types"
)

// note: ensure the struct implements the interface
var _ InterfaceHandler = (*BlockResponseHandler)(nil)

// BlockResponseHandler implements the block response handler
type BlockResponseHandler struct{}

// Func handles incoming block response messages
// TODO ...
func (b *BlockResponseHandler) Func(p clienttypes.InterfaceP2P, pr clienttypes.InterfacePeer, msg clienttypes.InterfaceMessage) error {
	//var msgStrBytes []byte
	//if err := msg.Unmarshal(msgBytes); err != nil {
	//logger.Errorf("[handler] err unmarshalling block response message\n%v", err)
	//return err
	//}

	//logger.Infof("%v BlockResponse: %v", pr.Cfg().ShortID, string(msgStrBytes))

	//s := p.Cfg().Syncer
	//if s == nil {
	//return errors.New("syncer is nil")
	//}

	//s.QueueBlocks(pr, msg)
	//s.RequestBlocks(pr)

	return nil
}

// Type returns the func enum
func (b *BlockResponseHandler) Type() handlertypes.FuncEnum {
	return handlertypes.BlockResponse
}
