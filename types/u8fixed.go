package types

import (
	"reflect"

	"github.com/opennetsys/godot/common/u8util"
)

// DefaultBitLength ...
var DefaultBitLength = 256

// U8Fixed ...
type U8Fixed struct {
	value  []byte
	length int
}

// NewU8Fixed ...
func NewU8Fixed(value []byte, bitLength int) *U8Fixed {
	if bitLength == -1 {
		bitLength = len(value) * 8
	}
	byteLength := bitLength / 8

	end := byteLength
	if end > len(value) {
		end = len(value)
	}

	sub := value[0:end]

	if len(sub) == byteLength {
		return &U8Fixed{
			length: byteLength,
			value:  sub[:],
		}
	}

	v := make([]uint8, byteLength)
	copy(v[:], sub)

	return &U8Fixed{
		length: byteLength,
		value:  v[:],
	}
}

// NewU8FixedFromHex ...
func NewU8FixedFromHex(value string, bitLength int) *U8Fixed {
	return NewU8Fixed(u8util.FromHex(value), bitLength)
}

// BitLen returns the number of bits in the value
func (u *U8Fixed) BitLen() int {
	return u.length * 8
}

// EncodedLen ...
func (u *U8Fixed) EncodedLen() int {
	return u.length
}

// Len ...
func (u *U8Fixed) Len() int {
	return len(u.value)
}

// String ...
func (u *U8Fixed) String() string {
	return u.Hex()
}

// Hex ...
func (u *U8Fixed) Hex() string {
	return u8util.ToHex(u.value, -1, true)
}

// IsEmpty ...
func (u *U8Fixed) IsEmpty() bool {
	if u.length == 0 {
		return true
	}

	if reflect.DeepEqual(u.value, make([]uint8, len(u.value))) {
		return true
	}

	return false
}

// Sub returns a sub slice
func (u *U8Fixed) Sub(start, end int) []uint8 {
	return u.value[start:end]
}

// Bytes ...
func (u *U8Fixed) Bytes() []byte {
	return u.value[:]
}

// ToU8a ...
func (u *U8Fixed) ToU8a(isBare bool) []byte {
	return u.value[:]
}

// Equals ...
func (u *U8Fixed) Equals(other interface{}) bool {
	switch v := other.(type) {
	case *U8Fixed:
		return u.String() == v.String()
	case []uint8:
		return reflect.DeepEqual(u.value, v)
	}

	return false
}
