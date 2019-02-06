package pair

import (
	ktypes "github.com/opennetsys/go-substrate/common/keyring/types"
)

var (
	// note: ensure the struct(s) implement the interface(s) at compile time
	_ InterfacePair  = (*Pair)(nil)
	_ InterfacePairs = (*Pairs)(nil)
)

// State ...
type State struct {
	Meta      ktypes.Meta
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
	Meta     ktypes.Meta
}

type encoding struct {
	Content encodingContentEnum
	Type    encodingTypeEnum
	Version string
}

// Pairs ...
type Pairs struct {
	PairMap MapPair
}

// MapPair ...
type MapPair map[string]*Pair

// note: implement nobody and everybody?
// https://github.com/polkadot-js/common/blob/master/packages/keyring/src/pair/nobody.ts
