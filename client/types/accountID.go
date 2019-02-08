package clienttypes

import (
	"github.com/opennetsys/godot/common/keyring/address"
)

// Encode ...
func (a *AccountID) Encode() (string, error) {
	return address.Encode(a[:], nil)
}
