package types

import (
	"math/big"

	"github.com/opennetsys/go-substrate/common/hexutil"
)

// UintBitLength ...
var UintBitLength = 8 | 16 | 32 | 64 | 128 | 256

// DefaultUintBits ...
var DefaultUintBits = 64

// AbstractInt ...
type AbstractInt struct {
	bitLength int
	value     *big.Int
}

// NewAbstractInt ...
func NewAbstractInt() *AbstractInt {
	// TODO
	return &AbstractInt{}
}

// Len ...
func (a *AbstractInt) Len() int {
	return a.bitLength
}

// EncodedLen ...
func (a *AbstractInt) EncodedLen() int {
	return a.bitLength / 8
}

// BitLen ...
func (a *AbstractInt) BitLen() int {
	return a.bitLength
}

// ToBN ...
func (a *AbstractInt) ToBN() {
	//return this
}

// Hex ...
func (a *AbstractInt) Hex() {

}

// String ...
func (a *AbstractInt) String() {
	return
}

// Bytes ...
func (a *AbstractInt) Bytes() {
}

// ToU8a ...
func (a *AbstractInt) ToU8a(isBare bool) {
}

func decodeAbstracInt(value interface{}, bitLength int64, isNegative bool) *big.Int {

	switch v := value.(type) {
	case string:
		if hexutil.ValidHex(v) {
			i, err := hexutil.ToBN(v, false, isNegative)
			if err != nil {
				panic(err)
			}
			return i
		}

		i := new(big.Int)
		i.SetString(v, 10)
		return i
	case []uint8:
	case int:
		return big.NewInt(int64(v))
	case int64:
		return big.NewInt(v)
	case uint:
		return big.NewInt(int64(v))
	case uint64:
		return big.NewInt(int64(v))
	case *big.Int:
		return v
	}

	return nil
}
