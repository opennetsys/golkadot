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

// Header ...
type Header struct{}

// SignedBlock ...
type SignedBlock struct{}

// MemoryBuffer ...
type MemoryBuffer map[int64]int64 // offset -> size

// HeapMemory ...
type HeapMemory struct {
	Allocated   MemoryBuffer
	Deallocated MemoryBuffer
	IsResized   bool
	Offset      int64
	End         int64
	Size        int64
	Buffer      []uint8
}

// Pointer ...
type Pointer int64

// SizeUsed ...
type SizeUsed struct {
	Allocated   int64
	Deallocated int64
}

// PageSize ...
var PageSize = 64 * 1024

// WasmMemory ...
type WasmMemory struct {
	Buffer []byte
}

// InterfaceEnvHeap ...
type InterfaceEnvHeap interface {
	Allocate(size int64) Pointer
	Deallocate(ptr Pointer) (int64, error)
	Dup(ptr Pointer, length int64) []uint8
	Fill(ptr Pointer, value uint8, length int64) []uint8
	Get(ptr Pointer, length int64) []uint8
	GetU32(ptr Pointer) []uint8
	Set(ptr Pointer, data []uint8) Pointer
	SetU32(ptr Pointer, value []uint8) Pointer
	SetWASMMemory(wasmMemory *WasmMemory, pageOffset int64)
	Size() int64
	Used() SizeUsed
	WasResized() bool
}

// Stat ...
type Stat struct {
	Average int64
	Calls   int64
	Elapsed int64
}

// Stats ...
type Stats map[string]Stat

// Interface ...
type Interface interface {
	//Environment Env
	//Exports InterfaceExports
	//Instrument InterfaceInstrument
}

// InterfaceInstrument ...
type InterfaceInstrument interface {
	Start()
	Stop() Stats
}

// InterfaceIO ...
type InterfaceIO interface {
	PrintHex(ptr Pointer, length int64)
	PrintUTF8(ptr Pointer, length int64)
	PrintNumber(hi int64, lo int64)
}

// InterfaceChain ...
type InterfaceChain interface {
	ChainID() int
}

// InterfaceCrypto ...
type InterfaceCrypto interface {
	Blake2b256(data Pointer, length int64, out Pointer)
	ED25519Verify(msgPtr Pointer, msgLen int64, sigPtr Pointer, pubkeyPtr Pointer) int64
	Twox128(data Pointer, length int64, out Pointer)
	Twox256(data Pointer, length int64, out Pointer)
}

// InterfaceMemory ...
type InterfaceMemory interface {
	Free(ptr Pointer)
	Malloc(size int64) Pointer
	Memcmp(s1 Pointer, s2 Pointer, length int64) int64
	Memcpy(dst Pointer, src Pointer, num int64) Pointer
	Memmove(dst Pointer, src Pointer, num int64) Pointer
	Memset(dst Pointer, val uint8, num int64) Pointer
}

// InterfaceSandbox ...
type InterfaceSandbox interface {
	Instantiate(a, b, c, d, e, f int64) int64
	InstanceTeardown(instanceIndex int64)
	Invoke(instanceIndex int64, exportPtr Pointer, exportLength int64, argsPtr Pointer, argsLength int64, returnValPtr Pointer, returnValLength int64, state int64)
	MemoryGet(memoryIndex int64, offset int64, ptr Pointer, length int64) int64
	MemoryNew(initial int64, maximum int64) int64
	MemorySet(memoryIndex int64, offset int64, ptr Pointer, length int64) int64
	MemoryTeardown(memoryIndex int64)
}

// InterfaceStorageData ...
type InterfaceStorageData interface {
	ClearPrefix(prefixPtr Pointer, prefixLength int64)
	ClearStorage(keyPtr Pointer, keyLength int64)
	ExistsStorage(keyPtr Pointer, keyLength int64) int64
	GetAllocatedStorage(keyPtr Pointer, keyLength int64, writtenPtr Pointer) Pointer
	GetStorageInto(keyPtr Pointer, keyLength int64, dataPtr Pointer, dataLength int64, offset int64) int64
	SetStorage(keyPtr Pointer, keyLength int64, dataPtr Pointer, dataLength int64)
}

// InterfaceStorageTrie ...
type InterfaceStorageTrie interface {
	Blake2b256EnumeratedTrieRoot(valuesPtr Pointer, lensPtr Pointer, lensLength int64, resultPtr Pointer)
	StorageChangesRoot(parentHashData Pointer, parentHashLength, parentNumHi, parentNumLo int64, result Pointer)
	StorageRoot(resultPtr Pointer)
}
