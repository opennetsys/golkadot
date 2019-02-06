package codec

import (
	"log"
	"math/big"

	"github.com/opennetsys/go-substrate/common/codec"
	codectypes "github.com/opennetsys/go-substrate/common/codec/types"
	"github.com/opennetsys/go-substrate/common/u8util"
)

// BNToCompact ...
func BNToCompact(val *big.Int) (*codectypes.Compact, error) {
	if val == nil {
		return nil, codectypes.ErrNilInput
	}

	if val.Cmp(big.NewInt(int64(codectypes.MAX_U8))) <= 0 {
		c := codectypes.Compact([]byte{uint8(val.Int64()) << 2})
		return &c, nil
	}
	if val.Cmp(big.NewInt(int64(codectypes.MAX_U16))) <= 0 {
		val = val.Lsh(val, 2)
		val = val.Add(val, big.NewInt(1))
		b, err := codec.Encode(int16(val.Int64()))
		c := codectypes.Compact(b)
		return &c, err
	}
	if val.Cmp(big.NewInt(int64(codectypes.MAX_U32))) <= 0 {
		val = val.Lsh(val, 2)
		val = val.Add(val, big.NewInt(2))
		b, err := codec.Encode(int32(val.Int64()))
		c := codectypes.Compact(b)
		return &c, err
	}

	// note: bytes is big-endian
	b := val.Bytes()
	bigEToLittleE(b)
	l := len(b)

	if l < 4 {
		return nil, codectypes.ErrInvalidLength
	}

	for ; b[l-1] == 0; l-- {
	}

	ret := []byte{uint8(((l - 4) << 2) + 3)}
	ret = append(ret, b[0:l]...)
	c := codectypes.Compact(ret)
	return &c, nil
}

// CompactMetaFromBytes retrieves the offset and encoded length from a compact-prefixed value
func CompactMetaFromBytes(input []byte) (*codectypes.CompactMeta, error) {
	if input == nil || len(input) == 0 {
		return nil, codectypes.ErrNilInput
	}

	switch input[0] & 3 {
	case 0:
		{
			l := u8util.ToBN([]byte{input[0]}, true)
			l = l.Rsh(l, 2)
			return &codectypes.CompactMeta{
				Offset: 1,
				Length: l,
			}, nil
		}

	case 1:
		{
			l := u8util.ToBN(input[0:2], true)
			l = l.Rsh(l, 2)
			return &codectypes.CompactMeta{
				Offset: 2,
				Length: l,
			}, nil
		}

	case 2:
		{
			l := u8util.ToBN(input[0:4], true)
			l = l.Rsh(l, 2)
			return &codectypes.CompactMeta{
				Offset: 4,
				Length: l,
			}, nil
		}

	default:
		{
			// note: clear flag and add 4 for base length
			l := u8util.ToBN([]byte{input[0]}, true)
			l = l.Rsh(l, 2)
			l = l.Add(l, big.NewInt(4))
			offset := 1 + int(l.Int64())

			bn := u8util.ToBN(input[1:offset], true)
			log.Println(bn.String())
			return &codectypes.CompactMeta{
				Offset: offset,
				Length: bn,
			}, nil
		}
	}
}

// note:
// @from: 		https://stackoverflow.com/a/51123337/3512709
// @asked: 		Shahul Hameed https://stackoverflow.com/users/4518319/shahul-hameed
// @answered: 	Tim Cooper https://stackoverflow.com/users/142162/tim-cooper
// @per: 		https://stackoverflow.blog/2009/06/25/attribution-required/
func bigEToLittleE(b []byte) {
	l := len(b)
	for i := 0; i < l/2; i++ {
		b[i], b[l-i-1] = b[l-i-1], b[i]
	}
}
