package types

import (
	"math/big"

	"github.com/opennetsys/go-substrate/common/u8compact"
	"github.com/opennetsys/go-substrate/common/u8util"
)

// Text is a string wrapper, along with the length.
type Text struct {
	value *string
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
		value: &value,
	}
}

// NewTextFromText ...
func NewTextFromText(value *Text) *Text {
	v := value.String()
	return &Text{
		value: &v,
	}
}

// NewTextFromBytes ...
func NewTextFromBytes(value []byte) *Text {
	offset, length := u8compact.FromUint8Slice(value, 8)

	end := int(offset) + int(length.Uint64())
	if end > len(value) {
		end = len(value)
	}

	result := value[offset:end]
	v := string(result)

	return &Text{
		value: &v,
	}
}

// DecodeText ...
func (t *Text) DecodeText(value string) {
	// TODO:
	// but not sure if needed considering text is immediately decoded
}

// Len ...
func (t *Text) Len() int {
	if t.value != nil {
		return len(*t.value)
	}

	return 0
}

// EncodedLen ...
func (t *Text) EncodedLen() int {
	return t.Len() + len(u8compact.CompactToUint8Slice(big.NewInt(int64(t.Len())), -1))
}

// String ...
func (t *Text) String() string {
	if t.value != nil {
		return *t.value
	}

	return ""
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
	var encoded []byte
	if t.value != nil {
		encoded = []byte(*t.value)
	}

	if isBare {
		return encoded
	}

	return u8compact.AddLength(encoded, -1)
}

// Equals ...
func (t *Text) Equals(other interface{}) bool {
	switch v := other.(type) {
	case *Text:
		return t.String() == v.String()
	case string:
		return t.String() == v
	case *string:
		return t.String() == *v
	}

	return false
}

// Value ...
func (t *Text) Value() *string {
	return t.value
}
