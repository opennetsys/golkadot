package clienttypes

import (
	"github.com/c3systems/go-substrate/common/keyring/address"
)

// Encode ...
func (a *AccountID) Encode() (string, error) {
	return address.Encode(a[:], nil)
}
