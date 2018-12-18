package system

import rpctypes "github.com/c3systems/go-substrate/rpc/types"

// ServiceInterface describes the methods performed by the system rpc service
type ServiceInterface interface {
	// Name returns the node's implementation name.
	Name(args rpctypes.NilArgs, response *string) error
	// Version returns the node's version. The result should be a semvar.
	Version(args rpctypes.NilArgs, response *string) error
	// Chain returns the node's chain type.
	Chain(args rpctypes.NilArgs, response *string) error
	// Properties returns the node's properties.
	Properties(args rpctypes.NilArgs, response *Properties) error
	// Health returns the health status of the node.
	//
	// Node is considered healthy if it is:
	// - connected to some peers (unless running in dev mode)
	// - not performing a major sync
	Health(args rpctypes.NilArgs, response *Health) error
}
