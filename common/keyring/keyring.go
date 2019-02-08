package keyring

import (
	"errors"

	"github.com/opennetsys/godot/common/crypto"
	"github.com/opennetsys/godot/common/keyring/address"
	"github.com/opennetsys/godot/common/keyring/pair"
	keytypes "github.com/opennetsys/godot/common/keyring/types"
	"github.com/opennetsys/godot/common/mnemonic"
)

// New ...
func New() (*KeyRing, error) {
	p, err := pair.NewPairs()
	if err != nil {
		return nil, err
	}

	return &KeyRing{
		Pairs: p,
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
func (k *KeyRing) AddFromAddress(addr []byte, meta keytypes.Meta, defaultEncoded []byte) (*pair.Pair, error) {
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
func (k *KeyRing) AddFromMnemonic(mn, password string, meta keytypes.Meta) (*pair.Pair, error) {
	seed, err := mnemonic.ToSeed(mn, password)
	if err != nil {
		return nil, err
	}

	return k.AddFromSeed(seed, meta)
}

// AddFromSeed ...
func (k *KeyRing) AddFromSeed(seed []byte, meta keytypes.Meta) (*pair.Pair, error) {
	pub, priv, err := crypto.NewNaclKeyPairFromSeed(seed)
	if err != nil {
		return nil, err
	}

	// TODO: nil defaultEncoded?
	pair, err := pair.NewPair(pub, priv, meta, nil)
	if err != nil {
		return nil, err
	}

	return k.AddPair(pair)
}

// AddFromJSON ...
func (k *KeyRing) AddFromJSON(data []byte, password *string) (*pair.Pair, error) {
	pair, err := pair.NewPairFromJSON(data, password)
	if err != nil {
		return nil, err
	}

	return k.AddPair(pair)
}

// GetPair ...
func (k *KeyRing) GetPair(addr []byte) (*pair.Pair, error) {
	if k.Pairs == nil {
		return nil, errors.New("nil pairs")
	}

	return k.Pairs.Get(addr)
}

// GetPairs ...
func (k *KeyRing) GetPairs() ([]*pair.Pair, error) {
	if k.Pairs == nil {
		return nil, errors.New("nil pairs")
	}

	return k.Pairs.All()
}

// GetPublicKeys ...
func (k *KeyRing) GetPublicKeys() ([][32]byte, error) {
	pairs, err := k.GetPairs()
	if err != nil {
		return nil, err
	}

	var pks [][32]byte
	for idx := range pairs {
		pk, err := pairs[idx].PublicKey()
		if err != nil {
			return nil, err
		}

		pks = append(pks, pk)
	}

	return pks, nil
}

// RemovePair ...
func (k *KeyRing) RemovePair(addr []byte) error {
	if k.Pairs == nil {
		return errors.New("nil pairs")
	}

	return k.Pairs.Remove(addr)
}

// ToJSON ...
func (k *KeyRing) ToJSON(addr []byte, password *string) ([]byte, error) {
	pair, err := k.GetPair(addr)
	if err != nil {
		return nil, err
	}

	return pair.ToJSON(password)
}
