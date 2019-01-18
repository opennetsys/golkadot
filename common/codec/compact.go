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

	//var b []byte
	//w := val.Bits()
	//for idx := range w {
	//b = append(b, uint8(w[idx]))
	//}

	// note: bytes is big-endian
	b := val.Bytes()
	bigEToLittleE(b)
	l := len(b)

	for ; b[l-1] == 0; l-- {
	}

	// note: assert l > 4?
	ret := []byte{uint8(((l - 4) << 2) + 3)}
	ret = append(ret, b[0:l]...)
	c := Compact(ret)
	return &c, nil
}

// note: https://stackoverflow.com/a/51123337/3512709
func bigEToLittleE(b []byte) {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
}
