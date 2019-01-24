package types

import (
	"encoding/hex"
	"math/big"

	"github.com/c3systems/go-substrate/common/bnutil"
	"github.com/c3systems/go-substrate/common/hexutil"
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
func (i *Int) ToU8a() []uint8 {
	return bnutil.ToUint8Slice(i.value, i.bitLen, true, false)
}

// Hex ...
func (i *Int) Hex() string {
	return hexutil.HexFixLength(hex.EncodeToString(i.value.Bytes()), i.bitLen, true)
}

// Equals ...
func (i *Int) Equals(value *Int) bool {
	return i.String() == value.String()
}
