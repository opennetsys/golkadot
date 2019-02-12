package storagetypes

import (
	"github.com/opennetsys/golkadot/common/u8compact"
)

// TODO ...

// SubstrateType ...
type SubstrateType struct{}

// Substrate ....
var Substrate *SubstrateType

// Code ...
func (s *SubstrateType) Code(i interface{}) []uint8 {
	key := []byte(":code")

	// StorageKey is a Bytes, so is length-prefixed
	return u8compact.AddLength(
		key,
		-1,
	)
}

func init() {
	Substrate = &SubstrateType{}
}
