package system

const (
	// Error is the rpc system error code.
	Error int64 = 2000
)

// Service implements the system service interface
type Service struct{}

// Properties is the struct returned by the properties api.
type Properties map[string]interface{}

// Health is the data returned by the Health api.
type Health struct {
	// Peers is the number of connected peers to the node.
	Peers uint32
	// IsSyncing returns true if the node is currently syncing.
	IsSyncing bool
}

// Info is the node's static details
type Info struct {
	// Name is the implementation name.
	Name string
	// Version is the implementation version.
	Version string
	// ChainName is the name of the chain.
	ChainName string
	// Properties are the custom set of properties defined in the chain spec.
	Properties Properties
}
