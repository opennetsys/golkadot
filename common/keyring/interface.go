package keyring

import (
	"github.com/c3systems/go-substrate/common/keyring/address"
	"github.com/c3systems/go-substrate/common/keyring/pair"
	keytypes "github.com/c3systems/go-substrate/common/keyring/types"
)

type InterfaceKeyRing interface {
	DecodeAddress(encoded []byte) ([]byte, error)
	EncodeAddress(key []byte) (string, error)
	SetAddressPrefix(prefix address.PrefixEnum) error
	AddPair(pair *pair.Pair) (*pair.Pair, error)
	AddFromAddress(address []byte, meta *keytypes.Meta, defaultEncoded []byte) (*pair.Pair, error)
	AddFromMnemonic(mnemonic string, meta *keytypes.Meta) (*pair.Pair, error)
	AddFromSeed(seed []byte, meta *keytypes.Meta) (*pair.Pair, error)
	AddFromJson(pairJSON []byte) (*pair.Pair, error)
	GetPair(address []byte) (*pair.Pair, error)
	GetPairs() []*pair.Pair
	GetPublicKeys() ([][]byte, error)
	RemovePair(address []byte) error
	ToJSON(address []byte, passphrase *string) ([]byte, error)
}
