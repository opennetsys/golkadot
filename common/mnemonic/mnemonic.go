package mnemonic

import (
	"errors"

	"github.com/tyler-smith/go-bip39"
)

// Generate ...
func Generate(entropy *int) (string, error) {
	if entropy == nil {
		entropy = new(int)
		*entropy = DefaultEntropy
	}

	// note: already checked for nil pointer
	ent, err := bip39.NewEntropy(*entropy)
	if err != nil {
		return "", err
	}

	return bip39.NewMnemonic(ent)
}

// Validate ...
func Validate(mnemonic string) bool {
	return bip39.IsMnemonicValid(mnemonic)
}

// ToSecret ...
func ToSecret(mnemonic, password string) ([]byte, error) {
	// TODO: is this correct? rather 'bip32.NewMasterKey(seed)'?
	return bip39.NewSeed(mnemonic, password), nil
}

// ToSeed ...
func ToSeed(mnemonic, password string) ([]byte, error) {
	tmp, err := ToSecret(mnemonic, password)
	if err != nil {
		return nil, err
	}

	if len(tmp) < 32 {
		return nil, errors.New("invalid secret generated")
	}

	// pair.DEFAULT_KEY_LENGTH?
	return tmp[:32], nil
}
