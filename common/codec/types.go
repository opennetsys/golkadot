package codec

import (
	"errors"
	"math/big"
)

// note: https://github.com/polkadot-js/common/blob/89030f41b34e5bb815f8004861c0424bdd3337f4/packages/util/src/compact/toU8a.ts#L11
const (
	// MAX_U8 ...
	MAX_U8 uint = 63
	// MAX_U16 ...
	MAX_U16 uint = 16383
	// MAX_U32 ...
	MAX_U32 uint = 1073741823
	// DEFAULT_BITLENGTH ...
	DEFAULT_BITLENGTH int = 32
)

var (
	// ErrInvalidKind ...
	ErrInvalidKind = errors.New("invalid kind")
	// ErrNilKind ...
	ErrNilKind = errors.New("kind cannot be nil")
	// ErrNilInput ...
	ErrNilInput = errors.New("input cannot be nil")
	// ErrNilTarget ...
	ErrNilTarget = errors.New("target cannot be nil")
	// ErrNonTargetPointer ...
	ErrNonTargetPointer = errors.New("target must be pointer")
	// ErrInvalidLength ...
	ErrInvalidLength = errors.New("invalid length")
)

// Compact ...
type Compact []byte

// CompactMeta ...
type CompactMeta struct {
	Offset int
	Length *big.Int
}
