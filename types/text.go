package types

import (
	"math/big"

	"github.com/c3systems/go-substrate/common/u8compact"
	"github.com/c3systems/go-substrate/common/u8util"
)

// Text is a string wrapper, along with the length.
type Text struct {
	value string
}

// NewText ...
func NewText(i interface{}) *Text {
	switch v := i.(type) {
	case string:
		return NewTextFromString(v)
	case []byte:
		return NewTextFromBytes(v)
	case *Text:
		return NewTextFromText(v)
	}

	return &Text{}
}

// NewTextFromString ...
func NewTextFromString(value string) *Text {
	return &Text{
		value: value,
	}
}

// NewTextFromText ...
func NewTextFromText(value *Text) *Text {
	return &Text{
		value: value.String(),
	}
}

// NewTextFromBytes ...
func NewTextFromBytes(value []byte) *Text {
	offset, length := u8compact.FromUint8Slice(value, 8)
	result := value[offset : uint64(offset)+length.Uint64()]

	return &Text{
		value: string(result),
	}
}

// DecodeText ...
func (t *Text) DecodeText(value string) {
	// TODO:
	// but not sure if needed considering text is immediately decoded
}

// Len ...
func (t *Text) Len() int {
	return len(t.value)
}

// EncodedLen ...
func (t *Text) EncodedLen() int {
	return t.Len() + len(u8compact.CompactToUint8Slice(big.NewInt(int64(t.Len())), -1))
}

// String ...
func (t *Text) String() string {
	return t.value
}

// Hex ...
func (t *Text) Hex() string {
	return u8util.ToHex(t.ToU8a(false), -1, true)
}

// Bytes ...
func (t *Text) Bytes() []byte {
	return t.ToU8a(false)
}

// ToU8a encodes the value as a []uint8 as per the parity-codec specifications. Set isBare to true when the value has none of the type-specific prefixes (internal)
func (t *Text) ToU8a(isBare bool) []uint8 {
	encoded := []byte(t.value)
	if isBare {
		return encoded
	}

	return u8compact.AddLength(encoded, -1)
}
