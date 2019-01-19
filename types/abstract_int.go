package types

import "math/big"

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
