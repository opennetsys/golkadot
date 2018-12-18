package runtime

// Version should not be thought of as classic Semver (major/minor/tiny).
// This triplet have different semantics and mis-interpretation could cause problems.
// In particular: bug fixes should result in an increment of `SpecVersion` and possibly `AuthoringVersion`,
// absolutely not `ImplVersion` since they change the semantics of the runtime.
type Version struct {
	// SpecName identifies the different Substrate runtimes. There'll be at least polkadot and node.
	// A different on-chain spec_name to that of the native runtime would normally result
	// in node not attempting to sync or author blocks.
	SpecName string
	// SpecVersion of the runtime specification. A full-node will not attempt to use its native
	// runtime in substitute for the on-chain Wasm runtime unless all of `SpecName`,
	// `SpecVersion` and `AuthoringVersion` are the same between Wasm and native.
	SpecVersion uint32
	// AuthoringVersion of the authorship interface. An authoring node
	// will not attempt to author blocks unless this is equal to its native runtime.
	AuthoringVersion uint32
	// ImplName is the name of the implementation of the spec. This is of little consequence for the node
	// and serves only to differentiate code of different implementation teams. For this
	// codebase, it will be parity-polkadot. If there were a non-Rust implementation of the
	// Polkadot runtime (e.g. C++), then it would identify itself with an accordingly different
	// `ImplName`.
	ImplName string
	// ImplVersion of the implementation of the specification. Nodes are free to ignore this; it
	// serves only as an indication that the code is different; as long as the other two versions
	// are the same then while the actual code may be different, it is nonetheless required to
	// do the same thing.
	// Non-consensus-breaking optimisations are about the only changes that could be made which
	// would result in only the `ImplVersion` changing.
	ImplVersion uint32
	// APIS is a list of supported API "features" along with their versions.
	APIS map[uint8]uint32
}

// Header...
type Header struct{}

// SignedBlock...
type SignedBlock struct{}
