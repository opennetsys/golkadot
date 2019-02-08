package types

import (
	"encoding/hex"
	"math/big"

	"github.com/opennetsys/godot/common/bnutil"
	"github.com/opennetsys/godot/common/hexutil"
)

// Int ...
type Int struct {
	value  *big.Int
	bitLen int
}

// NewInt ...
func NewInt(value interface{}, bitLen int) *Int {
	i := new(big.Int)
	switch v := value.(type) {
	case int:
		i = big.NewInt(int64(v))
	case uint:
		i = big.NewInt(int64(v))
	case int64:
		i = big.NewInt(v)
	case uint64:
		i = big.NewInt(int64(v))
	case string:
		return NewIntFromString(v, bitLen)
	}

	if bitLen == -1 {
		bitLen = DefaultUintBits
	}

	return &Int{
		value:  i,
		bitLen: bitLen,
	}
}

// NewIntFromHex ...
func NewIntFromHex(value string, bitLen int) *Int {
	i := new(big.Int)
	i.SetString(hexutil.StripPrefix(value), 16)

	return &Int{
		value:  i,
		bitLen: bitLen,
	}
}

// NewIntFromString ...
func NewIntFromString(value string, bitLen int) *Int {
	i := new(big.Int)
	i.SetString(value, 10)

	return &Int{
		value:  i,
		bitLen: bitLen,
	}
}

// BitLen ...
func (i *Int) BitLen() int {
	return i.bitLen
}

// Len ...
func (i *Int) Len() int {
	return i.BitLen()
}

// EncodedLen ...
func (i *Int) EncodedLen() int {
	return i.BitLen()
}

// BN ...
func (i *Int) BN() *big.Int {
	return i.value
}

// Int64 ...
func (i *Int) Int64() int64 {
	return i.value.Int64()
}

// Uint64 ...
func (i *Int) Uint64() uint64 {
	return i.value.Uint64()
}

// String ...
func (i *Int) String() string {
	return i.value.String()
}

// ToU8a ...
func (i *Int) ToU8a(isBare bool) []uint8 {
	return bnutil.ToUint8Slice(i.value, i.bitLen, true, false)
}

// Hex ...
func (i *Int) Hex() string {
	return hexutil.HexFixLength(hex.EncodeToString(i.value.Bytes()), i.bitLen, true)
}

// Equals ...
func (i *Int) Equals(other interface{}) bool {
	switch v := other.(type) {
	case *Int:
		return i.String() == v.String()
	case int:
		return i.Int64() == int64(v)
	case int64:
		return i.Int64() == v
	case uint:
		return i.Uint64() == uint64(v)
	case uint64:
		return i.Uint64() == v
	}

	return false
}

// IsZero ...
func (i *Int) IsZero() bool {
	if i == nil {
		return true
	}

	if i.value.Uint64() == 0 {
		return true
	}

	return false
}

// Bytes ...
func (i *Int) Bytes() []byte {
	return i.value.Bytes()
}
