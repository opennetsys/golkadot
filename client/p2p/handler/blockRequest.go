package handler

import (
	"github.com/c3systems/go-substrate/client/p2p"
	clienttypes "github.com/c3systems/go-substrate/client/types"
)

// note: ensure the struct implements the interface
var _ InterfaceHandler = (*BlockRequestHandler)(nil)

// BlockRequestHandler implements the block request handler
type BlockRequestHandler struct{}

// Func handles incoming block request messages
// TODO ...
func (b *BlockRequestHandler) Func(p p2p.InterfaceP2P, pr clienttypes.InterfacePeer, msg clienttypes.InterfaceMessage) error {
	//var msgStrBytes []byte
	//if err := msg.Unmarshal(msgBytes); err != nil {
	//logger.Errorf("[handler] err unmarshalling block response message\n%v", err)
	//return err
	//}
	//logger.Infof("%v BlockRequest: %v", pr.Cfg().ShortID, string(msgStrBytes))

	//s := p.Cfg().Syncer
	//if s == nil {
	//return errors.New("syncer is nil")
	//}

	//s.ProvideBlocks(pr, msg)

	return nil
}

// Type returns the func enum
func (b *BlockRequestHandler) Type() FuncEnum {
	return BlockRequest
}
