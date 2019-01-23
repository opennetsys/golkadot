package pair

import (
	ktypes "github.com/c3systems/go-substrate/common/keyring/types"
)

// InterfacePair ...
type InterfacePair interface {
	Address() (string, error)
	DecodePkcs8(passphrase *string, encoded []byte) error
	EncodePkcs8(passphrase *string) ([]byte, error)
	GetMeta() (*ktypes.Meta, error)
	IsLocked() bool
	Lock() error
	PublicKey() ([32]byte, error)
	SetMeta(meta *ktypes.Meta) error
	Sign(message []byte) ([]byte, error)
	ToJSON(passphrase *string) ([]byte, error)
	Verify(message, signature []byte) (bool, error)
}

// InterfacePairs ...
type InterfacePairs interface {
	Add(pair *Pair) (*Pair, error)
	All() []*Pair
	Get(address []byte) (*Pair, error)
	Remove(address []byte) error
}
