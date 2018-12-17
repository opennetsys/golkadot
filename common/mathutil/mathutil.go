package mathutil

import "math/big"

// Pow ...
func Pow(i, e *big.Int) *big.Int {
	return new(big.Int).Exp(i, e, nil)
}
