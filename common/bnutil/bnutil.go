package bnutil

import (
	"encoding/hex"
	"math"
	"math/big"

	"github.com/c3systems/go-substrate/common/hexutil"
	"github.com/c3systems/go-substrate/common/mathutil"
	"github.com/c3systems/go-substrate/common/u8util"
)

// FromHex creates a math/big big number from a hex string.
func FromHex(hexStr string) (*big.Int, error) {
	return hexutil.ToBN(hexStr, false, false)
}

// ToBN creates a BN value from a number input
func ToBN(ivalue interface{}, isLittleEndian bool) *big.Int {
	switch v := ivalue.(type) {
	case *big.Int:
		return v
	case string:
		i := new(big.Int)
		i.SetString(v, 10)
		return i
	case float32:
		return big.NewInt(int64(v))
	case float64:
		return big.NewInt(int64(v))
	case int64:
		return big.NewInt(int64(v))
	case int32:
		return big.NewInt(int64(v))
	case int16:
		return big.NewInt(int64(v))
	case int8:
		return big.NewInt(int64(v))
	case int:
		return big.NewInt(int64(v))
	case uint64:
		return big.NewInt(int64(v))
	case uint32:
		return big.NewInt(int64(v))
	case uint16:
		return big.NewInt(int64(v))
	case uint8:
		return big.NewInt(int64(v))
	case uint:
		return big.NewInt(int64(v))
	case []uint8:
		hx := hex.EncodeToString(v)
		n, err := hexutil.ToBN(hx, isLittleEndian, false)
		if err != nil {
			panic(err)
		}
		return n
	}

	return new(big.Int)
}

// ToHex creates a hex value from a math/big big number. 0 inputs returns a `0x` result, BN values return the actual value as a `0x` prefixed hex value. With `bitLength` set, it fixes the number to the specified length.
func ToHex(value *big.Int, bitLength int, isNegative bool) string {
	slice := ToUint8Slice(value, bitLength, false, false)
	return u8util.ToHex(slice, bitLength, true)
}

// ToUint8Slice creates a uint8 array from a big number. Empty input returns an empty uint8 array result. Convert using little-endian format if `isLittleEndian` is set.
func ToUint8Slice(value *big.Int, bitLength int, isLittleEndian bool, isNegative bool) []uint8 {
	var byteLength int
	if bitLength == -1 {
		byteLength = int(math.Ceil(float64(mathutil.BitLen(value)) / float64(8)))
	} else {
		byteLength = int(math.Ceil(float64(bitLength) / float64(8)))
	}

	if value == nil {
		if bitLength == -1 {
			return []uint8{}
		}
		return make([]uint8, byteLength)
	}

	if isNegative {
		value = mathutil.ToTwos(value, byteLength*8)
	}

	return mathutil.ToUint8Slice(value, isLittleEndian, byteLength)
}

// Uint8Slice  ...
// Example: sort.Sort(sort.Reverse(Uint8Slice(output)))
type Uint8Slice []uint8

// Len ...
func (s Uint8Slice) Len() int {
	return len(s)
}

// Less ...
func (s Uint8Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap ...
func (s Uint8Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
