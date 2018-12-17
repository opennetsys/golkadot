package u8compact

import (
	"math/big"

	"github.com/c3systems/go-substrate/common/bn"
	"github.com/c3systems/go-substrate/common/u8util"
)

// DefaultBitLength ...
const DefaultBitLength = 32

// FromUint8Slice retrieves the offset and encoded length from a compact-prefixed value
func FromUint8Slice(input []uint8, bitLength int) (int, *big.Int) {
	flag := input[0] & 0x3

	if flag == 0x0 {
		x := big.NewInt(int64(input[0]))
		return 1, new(big.Int).Rsh(x, 2)
	} else if flag == 0x1 {
		x := bn.ToBN(input[0:2], true)
		y := new(big.Int).Rsh(x, 2)
		return 2, y
	} else if flag == 0x2 {
		x := bn.ToBN(input[0:4], true)
		return 4, new(big.Int).Rsh(x, 2)
	}

	offset := 1 + (bitLength / 8)
	end := offset
	if end > len(input) {
		end = len(input)
	}
	return offset, bn.ToBN(input[1:end], true)

}

// CompactToUint8Slice encodes a number into a compact representation.
func CompactToUint8Slice(value *big.Int, bitLength int) []uint8 {
	maxU8 := new(big.Int).Sub(pow(big.NewInt(2), big.NewInt(8-2)), big.NewInt(1))
	maxU16 := new(big.Int).Sub(pow(big.NewInt(2), big.NewInt(16-2)), big.NewInt(1))
	maxU32 := new(big.Int).Sub(pow(big.NewInt(2), big.NewInt(32-2)), big.NewInt(1))

	if value.Cmp(maxU8) <= 0 {
		return []uint8{uint8(value.Int64() << 2)}
	} else if value.Cmp(maxU16) <= 0 {
		i := new(big.Int).Add(new(big.Int).Lsh(value, 2), big.NewInt(1))
		return bn.ToUint8Slice(i, 16, true)
	} else if value.Cmp(maxU32) <= 0 {
		i := new(big.Int).Add(new(big.Int).Lsh(value, 2), big.NewInt(2))
		return bn.ToUint8Slice(i, 32, true)
	}

	return u8util.Concat(
		[]uint8{0x3},
		bn.ToUint8Slice(value, bitLength, true),
	)
}

// AddLength adds a length prefix to the input value.
func AddLength(input []uint8, bitLength int) []uint8 {
	return u8util.Concat(
		CompactToUint8Slice(big.NewInt(int64(len(input))), bitLength),
		input,
	)
}

func pow(i, e *big.Int) *big.Int {
	return new(big.Int).Exp(i, e, nil)
}
