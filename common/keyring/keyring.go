package keyring

import (
	"errors"

	"github.com/c3systems/go-substrate/common/keyring/address"
	"github.com/c3systems/go-substrate/common/keyring/pair"
	keytypes "github.com/c3systems/go-substrate/common/keyring/types"
)

// New ...
func New() (*KeyRing, error) {
	return &KeyRing{
		Pairs: new(pair.Pairs),
	}, nil
}

// DecodeAddress ...
func (k *KeyRing) DecodeAddress(encoded []byte) ([]byte, error) {
	return address.Decode(string(encoded), nil)
}

// EncodeAddress ...
func (k *KeyRing) EncodeAddress(key []byte) (string, error) {
	return address.Encode(key, nil)
}

// SetAddressPrefix ...
func (k *KeyRing) SetAddressPrefix(prefix address.PrefixEnum) error {
	// note: is this really what we want? We won't be able to create multiple keyrings with different default prefixes...
	address.SetDefaultPrefix(prefix)
	return nil
}

// AddPair ...
func (k *KeyRing) AddPair(pair *pair.Pair) (*pair.Pair, error) {
	if k.Pairs == nil {
		return nil, errors.New("pairs is nil")
	}

	return k.Pairs.Add(pair)
}

// AddFromAddress ...
func (k *KeyRing) AddFromAddress(addr []byte, meta *keytypes.Meta, defaultEncoded []byte) (*pair.Pair, error) {
	tmp, err := address.Decode(string(addr), nil)
	if err != nil {
		return nil, err
	}
	var pub [32]byte
	copy(pub[:], tmp)

	// note: this pair will be locked bc no secret key ...
	pair, err := pair.NewPair(pub, [64]byte{}, meta, defaultEncoded)
	if err != nil {
		return nil, err
	}

	return k.AddPair(pair)
}

// AddFromMnemonic ...
func (k *KeyRing) AddFromMnemonic(mnemonic string, meta *keytypes.Meta) (*pair.Pair, error) {
	return nil, nil
}

// AddFromSeed ...
func (k *KeyRing) AddFromSeed(seed []byte, meta *keytypes.Meta) (*pair.Pair, error) {
	return nil, nil
}

// AddFromJSON ...
func (k *KeyRing) AddFromJSON(pairJSON []byte) (*pair.Pair, error) {
	return nil, nil
}

// GetPair ...
func (k *KeyRing) GetPair(address []byte) (*pair.Pair, error) {
	return nil, nil
}

// GetPairs ...
func (k *KeyRing) GetPairs() []*pair.Pair {
	return nil
}

// GetPublicKeys ...
func (k *KeyRing) GetPublicKeys() ([][]byte, error) {
	return nil, nil
}

// RemovePair ...
func (k *KeyRing) RemovePair(address []byte) error {
	return nil
}

// ToJSON ...
func (k *KeyRing) ToJSON(address []byte, passphrase *string) ([]byte, error) {
	return nil, nil
}
