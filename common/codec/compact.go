package codec

import (
	"math/big"
)

// EncodeToCompact ...
func EncodeToCompact(val *big.Int) (*Compact, error) {
	if val == nil {
		return nil, ErrNilInput
	}

	if val.Cmp(big.NewInt(int64(MAX_U8))) <= 0 {
		c := Compact([]byte{uint8(val.Int64()) << 2})
		return &c, nil
	}
	if val.Cmp(big.NewInt(int64(MAX_U16))) <= 0 {
		val = val.Lsh(val, 2)
		val = val.Add(val, big.NewInt(1))
		b, err := Encode(int16(val.Int64()))
		c := Compact(b)
		return &c, err
	}
	if val.Cmp(big.NewInt(int64(MAX_U32))) <= 0 {
		val = val.Lsh(val, 2)
		val = val.Add(val, big.NewInt(2))
		b, err := Encode(int32(val.Int64()))
		c := Compact(b)
		return &c, err
	}

	// note: bytes is big-endian
	b := val.Bytes()
	bigEToLittleE(b)
	l := len(b)

	if l < 4 {
		return nil, ErrInvalidLength
	}

	for ; b[l-1] == 0; l-- {
	}

	ret := []byte{uint8(((l - 4) << 2) + 3)}
	ret = append(ret, b[0:l]...)
	c := Compact(ret)
	return &c, nil
}

// note: https://stackoverflow.com/a/51123337/3512709
// @asked Shahul Hameed https://stackoverflow.com/users/4518319/shahul-hameed
// @answered Tim Cooper https://stackoverflow.com/users/142162/tim-cooper
// per https://stackoverflow.blog/2009/06/25/attribution-required/
func bigEToLittleE(b []byte) {
	l := len(b)
	for i := 0; i < l/2; i++ {
		b[i], b[l-i-1] = b[l-i-1], b[i]
	}
}
