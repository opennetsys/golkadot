package mathutil

import (
	"fmt"
	"math/big"
	"testing"
)

func TestPow(t *testing.T) {
	type input struct {
		i *big.Int
		e *big.Int
	}
	for i, tt := range []struct {
		in  input
		out *big.Int
	}{
		{input{big.NewInt(16), big.NewInt(2)}, big.NewInt(256)},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := Pow(tt.in.i, tt.in.e)
			if result.String() != tt.out.String() {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}
