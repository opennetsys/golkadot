package types

// U8a ...
type U8a struct {
	value []uint8
}

// NewU8a ...
func NewU8a(input interface{}) *U8a {
	var value []uint8
	switch v := input.(type) {
	case []byte:
		value = v
	}

	return &U8a{
		value: value,
	}
}

// Value ...
func (u *U8a) Value() []uint8 {
	return u.value
}
