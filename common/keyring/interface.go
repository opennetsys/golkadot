package keyring

import (
	"github.com/opennetsys/go-substrate/common/keyring/address"
	"github.com/opennetsys/go-substrate/common/keyring/pair"
	keytypes "github.com/opennetsys/go-substrate/common/keyring/types"
)

// InterfaceKeyRing ...
type InterfaceKeyRing interface {
	DecodeAddress(encoded []byte) ([]byte, error)
	EncodeAddress(key []byte) (string, error)
	SetAddressPrefix(prefix address.PrefixEnum) error
	AddPair(pair *pair.Pair) (*pair.Pair, error)
	AddFromAddress(addr []byte, meta keytypes.Meta, defaultEncoded []byte) (*pair.Pair, error)
	AddFromMnemonic(mn, password string, meta keytypes.Meta) (*pair.Pair, error)
	AddFromSeed(seed []byte, meta keytypes.Meta) (*pair.Pair, error)
	AddFromJSON(data []byte, password *string) (*pair.Pair, error)
	GetPair(addr []byte) (*pair.Pair, error)
	GetPairs() ([]*pair.Pair, error)
	GetPublicKeys() ([][32]byte, error)
	RemovePair(addr []byte) error
	// note: change to Marshal? Add Unmarshal?
	ToJSON(addr []byte, password *string) ([]byte, error)
}
