package clienttypes

import (
	"github.com/opennetsys/golkadot/common/keyring/address"
)

// Encode ...
func (a *AccountID) Encode() (string, error) {
	return address.Encode(a[:], nil)
}
