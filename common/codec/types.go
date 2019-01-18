package codec

// note: https://github.com/polkadot-js/common/blob/89030f41b34e5bb815f8004861c0424bdd3337f4/packages/util/src/compact/toU8a.ts#L11
const (
	// MAX_U8 ...
	MAX_U8 uint = 63
	// MAX_U16 ...
	MAX_U16 uint = 16383
	// MAX_U32 ...
	MAX_U32 uint = 1073741823
)

// Compact ...
type Compact []byte
