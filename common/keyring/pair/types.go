package pair

import (
	ktypes "github.com/c3systems/go-substrate/common/keyring/types"
)

var (
	// note: ensure the pair struct implements the interface at compile time
	_ InterfacePair = (*Pair)(nil)
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

// note: implement nobody and everybody?
// https://github.com/polkadot-js/common/blob/master/packages/keyring/src/pair/nobody.ts
