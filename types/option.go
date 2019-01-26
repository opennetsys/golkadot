package types

import (
	"reflect"

	"github.com/c3systems/go-substrate/common/u8util"
)

// Option is an optional field. Basically the first byte indicates that there is is value to follow. If the byte is `1` there is an actual value. So the Option implements that - decodes, checks for optionality and wraps the required structure with a value if/as required/found.
type Option struct {
	kind  reflect.Type
	value InterfaceType
}

// NewOption ...
func NewOption(value InterfaceType) *Option {
	switch v := value.(type) {
	case *Text:
		if v.Value() == nil {
			value = NewNull()
		}
	}

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
	switch o.value.(type) {
	case *Null:
		return true
	}
	return false
}

// IsSome ...
func (o *Option) IsSome() bool {
	return !o.IsNone()
}

// String ...
func (o *Option) String() string {
	return o.value.String()
}

// Hex ...
func (o *Option) Hex() string {
	return u8util.ToHex(o.ToU8a(false), -1, true)
}

// ToU8a ...
func (o *Option) ToU8a(isBare bool) []uint8 {
	if isBare {
		o.value.ToU8a(true)
	}

	slice := make([]uint8, o.EncodedLen())

	if o.IsSome() {
		copy(slice[:], []uint8{1})
		copy(slice[1:], o.value.ToU8a(false))
	}

	return slice
}

// Len ...
func (o *Option) Len() int {
	return o.value.Len()
}

// EncodedLen ...
func (o *Option) EncodedLen() int {
	// boolean byte (has value, doesn't have) along with wrapped length
	return 1 + o.value.EncodedLen()
}

// Bytes ...
func (o *Option) Bytes() []byte {
	return o.value.Bytes()
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
	case InterfaceType:
		return o.Equals(v)
	default:
		switch u := o.value.(type) {
		case *Text:
			return u.Equals(other)
		case *Int:
			return u.Equals(other)
		case *U8Fixed:
			return u.Equals(other)
		}
	}

	return false
}
