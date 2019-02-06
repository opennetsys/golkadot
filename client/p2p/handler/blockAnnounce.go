package handler

import (
	handlertypes "github.com/opennetsys/go-substrate/client/p2p/handler/types"
	clienttypes "github.com/opennetsys/go-substrate/client/types"
)

// note: ensure the struct implements the interface
var _ InterfaceHandler = (*BlockAnnounceHandler)(nil)

// BlockAnnounceHandler implements the block announce handler
type BlockAnnounceHandler struct{}

// Func handles incoming block announce messages
// TODO ...
func (b *BlockAnnounceHandler) Func(p clienttypes.InterfaceP2P, pr clienttypes.InterfacePeer, msg clienttypes.InterfaceMessage) error {
	//var msgStrBytes []byte
	//if err := msg.Unmarshal(msgBytes); err != nil {
	//logger.Errorf("[handler] err unmarshalling block response message\n%v", err)
	//return err
	//}
	//logger.Infof("%v BlockAnnounce: %v", pr.Cfg().ShortID, string(msgStrBytes))

	//header := msg.Header()

	//if pr.Cfg().BestNumber.Cmp(header.BlockNumber) == -1 {
	//pr.SetBest(header.BlockNumber, header.Hash)
	//}

	//s := p.Cfg().Syncer
	//if s == nil {
	//return errors.New("syncer is nil")
	//}
	//s.RequestBlocks(peer)

	return nil
}

// Type returns the func enum
func (b *BlockAnnounceHandler) Type() handlertypes.FuncEnum {
	return handlertypes.BlockAnnounce
}
