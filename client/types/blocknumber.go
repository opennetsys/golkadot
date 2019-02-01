package clienttypes

import (
	"math/big"
	"strconv"
)

// NewBlockNumber ...
func NewBlockNumber(value interface{}) *BlockNumber {
	var i uint64
	switch v := value.(type) {
	case int:
		i = uint64(v)
	case uint:
		i = uint64(v)
	case int64:
		i = uint64(v)
	case uint64:
		i = v
	case string:
		u, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
		i = uint64(u)
	case *big.Int:
		i = uint64(v.Int64())
	}
	return &BlockNumber{
		value: i,
	}
}

// Value ...
func (b *BlockNumber) Value() uint64 {
	return b.value
}

// Uint64 ...
func (b *BlockNumber) Uint64() uint64 {
	return b.value
}

// BN ...
func (b *BlockNumber) BN() *big.Int {
	return big.NewInt(int64(b.value))
}

// String ...
func (b *BlockNumber) String() string {
	return strconv.FormatUint(b.value, 10)
}
