package types

import (
	"reflect"

	"github.com/opennetsys/go-substrate/common"
	"github.com/opennetsys/go-substrate/common/u8util"
)

// Option is an optional field. Basically the first byte indicates that there is is value to follow. If the byte is `1` there is an actual value. So the Option implements that - decodes, checks for optionality and wraps the required structure with a value if/as required/found.
type Option struct {
	kind  reflect.Type
	value interface{}
}

// NewOption ...
func NewOption(value interface{}) *Option {
	return &Option{
		kind:  reflect.TypeOf(value),
		value: value,
	}
}

// Value ...
func (o *Option) Value() interface{} {
	return o.value
}

// IsNone ...
func (o *Option) IsNone() bool {
	return common.TypeIsNil(o.value)
}

// IsSome ...
func (o *Option) IsSome() bool {
	return !o.IsNone()
}

// String ...
func (o *Option) String() string {
	return common.TypeToString(o.value)
}

// Hex ...
func (o *Option) Hex() string {
	return u8util.ToHex(o.ToU8a(false), -1, true)
}

// ToU8a ...
func (o *Option) ToU8a(isBare bool) []uint8 {
	if isBare {
		return common.TypeToU8a(o.value, true)
	}

	slice := make([]uint8, o.EncodedLen())

	if o.IsSome() {
		copy(slice[:], []uint8{1})
		copy(slice[1:], common.TypeToU8a(o.value, false))
	}

	return slice
}

// Len ...
func (o *Option) Len() int {
	return common.TypeLen(o.value)
}

// EncodedLen ...
func (o *Option) EncodedLen() int {
	// boolean byte (has value, doesn't have) along with wrapped length
	return 1 + common.TypeEncodedLen(o.value)
}

// Bytes ...
func (o *Option) Bytes() []byte {
	return common.TypeToBytes(o.value)
}

// Unwrap returns the value that the Option represents (if available)
func (o *Option) Unwrap() interface{} {
	if o.IsNone() {
		panic("Option: unwrapping a None value")
	}

	return o.value
}

// Equals ...
func (o *Option) Equals(other interface{}) bool {
	// TODO: all types
	switch v := other.(type) {
	case *Option:
		return o.String() == v.String()
	default:
		return common.TypeEquals(o.value, other)
	}
}
