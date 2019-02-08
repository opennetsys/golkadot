package handler

import (
	handlertypes "github.com/opennetsys/godot/client/p2p/handler/types"
	clienttypes "github.com/opennetsys/godot/client/types"
)

// InterfaceHandler describes the methods of the handler package
type InterfaceHandler interface {
	// Func handles the message
	Func(p clienttypes.InterfaceP2P, pr clienttypes.InterfacePeer, msg clienttypes.InterfaceMessage) error
	// Type returns the handler type
	Type() handlertypes.FuncEnum
}
