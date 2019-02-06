package clienttypes

import (
	"github.com/opennetsys/go-substrate/common/keyring/address"
)

// Encode ...
func (a *AccountID) Encode() (string, error) {
	return address.Encode(a[:], nil)
}
