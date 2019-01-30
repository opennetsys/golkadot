package handler

import (
	handlertypes "github.com/c3systems/go-substrate/client/p2p/handler/types"
	clienttypes "github.com/c3systems/go-substrate/client/types"
)

// InterfaceHandler describes the methods of the handler package
type InterfaceHandler interface {
	// Func handles the message
	Func(p clienttypes.InterfaceP2P, pr clienttypes.InterfacePeer, msg clienttypes.InterfaceMessage) error
	// Type returns the handler type
	Type() handlertypes.FuncEnum
}
