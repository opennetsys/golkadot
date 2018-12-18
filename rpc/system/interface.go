package system

// ServiceInterface describes the methods performed by the system rpc service
type ServiceInterface interface {
	// Name returns the node's implementation name.
	Name() string
	// Version returns the node's version. The result should be a semvar.
	Version() string
	// Chain returns the node's chain type.
	Chain() string
	// Properties returns the node's properties.
	Properties() Properties
	// Health returns the health status of the node.
	//
	// Node is considered healthy if it is:
	// - connected to some peers (unless running in dev mode)
	// - not performing a major sync
	Health() Health
}
