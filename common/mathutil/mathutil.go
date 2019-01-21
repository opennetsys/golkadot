package mathutil

import (
	"math"
	"math/big"
)

// Pow ...
func Pow(i, e *big.Int) *big.Int {
	return new(big.Int).Exp(i, e, nil)
}

// FromTwos ...
func FromTwos(i *big.Int, width int) *big.Int {
	a := inotn(i, width)
	b := new(big.Int).Add(a, big.NewInt(1))
	return new(big.Int).Neg(b)
}

// inotn ...
func inotn(n *big.Int, width int) *big.Int {
	var bytesNeeded = int(math.Ceil(float64(width) / float64(26)))
	var bitsLeft = width % 26
	var length = n.BitLen()

	words := make([]big.Word, bytesNeeded)
	copy(words[:], n.Bits())

	// extend the buffer with leading zeroes
	for length < bytesNeeded {
		words[length] = 0
		length = length + 1
	}

	if bitsLeft > 0 {
		bytesNeeded--
	}

	var i int
	// handle complete words
	for i = 0; i < bytesNeeded; i++ {
		words[i] = big.Word(^int(words[i]) & 0x3ffffff)
	}

	// handle the residue
	if bitsLeft > 0 {
		words[i] = big.Word(^int(words[i]) & (0x3ffffff >> (26 - uint(bitsLeft))))
	}

	ret := new(big.Int)
	ret.SetBits(words)
	return ret
}
