package handler

import (
	"github.com/c3systems/go-substrate/client/p2p"
	clienttypes "github.com/c3systems/go-substrate/client/types"
)

// InterfaceHandler describes the methods of the handler package
type InterfaceHandler interface {
	// Func handles the message
	Func(p p2p.InterfaceP2P, pr clienttypes.InterfacePeer, msg clienttypes.InterfaceMessage) error
	// Type returns the handler type
	Type() FuncEnum
}
