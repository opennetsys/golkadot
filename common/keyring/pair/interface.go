package pair

import (
	ktypes "github.com/opennetsys/godot/common/keyring/types"
)

// InterfacePair ...
type InterfacePair interface {
	Address() (string, error)
	DecodePkcs8(password *string, encoded []byte) error
	EncodePkcs8(password *string) ([]byte, error)
	GetMeta() (ktypes.Meta, error)
	IsLocked() bool
	Lock() error
	PublicKey() ([32]byte, error)
	SetMeta(meta ktypes.Meta) error
	Sign(message []byte) ([]byte, error)
	// note: change to Marshal?
	ToJSON(password *string) ([]byte, error)
	Verify(message, signature []byte) (bool, error)
}

// InterfacePairs ...
type InterfacePairs interface {
	Add(pair *Pair) (*Pair, error)
	All() ([]*Pair, error)
	Get(addr []byte) (*Pair, error)
	Remove(addr []byte) error
}
