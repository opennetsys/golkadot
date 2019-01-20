package types

import "math/big"

// Int ...
type Int struct {
	value *big.Int
}

// NewInt ...
func NewInt(value interface{}) *Int {
	var i *big.Int
	switch v := value.(type) {
	case string:
		i.SetString(v, 10)
	case int:
		i = big.NewInt(int64(v))
	case int64:
		i = big.NewInt(v)
	}
	return &Int{
		value: i,
	}
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
