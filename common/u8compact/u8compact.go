package u8compact

import (
	"math/big"

	"github.com/c3systems/go-substrate/common/bn"
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
