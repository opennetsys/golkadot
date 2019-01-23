package keyring

import "github.com/c3systems/go-substrate/common/keyring/address"

type KeyRingInterface interface {
	DecodeAddress(encoded []byte) ([]byte, error)
	EncodeAddress(key []byte) ([]byte, error)
	SetAddressPrefix(prefix address.PrefixEnum) error
	AddPair(pair *pair.Pair) (*pair.Pair, error)
	AddFromAddress(address []byte, meta *pair.Meta) (*pair.Pair, error)
	AddFromMnemonic(mnemonic string, meta *pair.Meta) (*pair.Pair, error)
	AddFromSeed(seed []byte, meta *pair.Meta) (*pair.Pair, error)
	AddFromJson(pairJSON []byte) (*pair.Pair, error)
	GetPair(address []byte) (*pair.Pair, error)
	GetPairs() []*pair.Pair
	GetPublicKeys() ([][]byte, error)
	RemovePair(address []byte) error
	ToJSON(address []byte, passphrase *string) ([]byte, error)
}
