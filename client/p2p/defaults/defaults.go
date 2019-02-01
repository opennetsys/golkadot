package defaults

// Defaults ...
// TODO: don't use struct
var Defaults = struct {
	// MaxRequestBlocks ...
	MaxRequestBlocks uint
	// ProtocolBase ...
	ProtocolBase string
	// ProtocolType ...
	ProtocolType string
	// ProtocolVersion ...
	ProtocolVersion string
	// Address ...
	Address string
	// ClientID ...
	ClientID string
	// MaxPeers ...
	MaxPeers uint
	// MaxQueuedBlocks ...
	MaxQueuedBlocks uint
	// MinIdleBlocks ...
	MinIdleBlocks uint
	// Port ...
	Port uint
	// Role ...
	Role string
	// ProtocolDot ...
	ProtocolDot string
	// ProtocolPing ...
	ProtocolPing string
	// DialBackoff ...
	DialBackoff uint
	// DialInterval ...
	DialInterval uint
	// RequestInterval ...
	RequestInterval uint
	// PingInterval ...
	PingInterval uint
	// PingTimeout ...
	PingTimeout uint
	// Name is the version name.
	Name string
}{
	MaxRequestBlocks: 64,
	// TODO: substrate?
	ProtocolBase: "/substrate",
	// TODO: type?
	ProtocolType:    "/sup",
	ProtocolVersion: "1.0.0",
	Address:         "127.0.0.1",
	// TODO: ClientID?
	ClientID:        "polkadot-go/0.0.0",
	MaxPeers:        25,
	MaxQueuedBlocks: 64 * 8, // MaxRequestBlocks * 8
	MinIdleBlocks:   16,
	Port:            31333,
	Role:            "full",
	// TODO: substrate?
	ProtocolDot:     "/substrate/sup/1.0.0", // {ProtocolBase}/{ProtocolType}/{ProtocolVersion}
	ProtocolPing:    "/ipfs/ping/1.0.0",
	DialBackoff:     5 * 60000,
	DialInterval:    15000,
	RequestInterval: 15000,
	PingInterval:    30000,
	PingTimeout:     5000,
	Name:            "dot",
}
