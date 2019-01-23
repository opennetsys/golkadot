package pair

import (
	ktypes "github.com/c3systems/go-substrate/common/keyring/types"
)

// State ...
type State struct {
	Meta      *ktypes.Meta
	PublicKey [32]byte
}

// Pair ...
type Pair struct {
	State          *State
	defaultEncoded []byte
	secretKey      [64]byte
}

type forJSON struct {
	Address  string
	Encoded  string
	Encoding encoding
	Meta     *ktypes.Meta
}

type encoding struct {
	Content EncodingContentEnum
	Type    EncodingTypeEnum
	Version string
}
