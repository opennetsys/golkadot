package bn

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestFromHex(t *testing.T) {
	tests := []struct {
		in  string
		out *big.Int
	}{
		{
			"0x14",
			big.NewInt(20),
		},
		{
			"81",
			big.NewInt(129),
		},
		{
			"0x",
			big.NewInt(0),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := FromHex(tt.in)
			if err != nil {
				t.Error(err)
			}
			if result.String() != tt.out.String() {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestToBN(t *testing.T) {
	tests := []struct {
		in  interface{}
		out *big.Int
	}{
		{
			"14",
			big.NewInt(14),
		},
		{
			81,
			big.NewInt(81),
		},
		{
			float64(21),
			big.NewInt(21),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := ToBN(tt.in)
			if result.String() != tt.out.String() {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestToHex(t *testing.T) {
	type input struct {
		i         *big.Int
		bitLength int
	}
	tests := []struct {
		in  input
		out string
	}{
		{
			input{
				big.NewInt(128),
				-1,
			},
			"0x80",
		},
		{
			input{
				big.NewInt(81),
				-1,
			},
			"0x51",
		},
		{
			input{
				big.NewInt(21),
				-1,
			},
			"0x15",
		},
		{
			input{
				big.NewInt(128),
				16,
			},
			"0x0080",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := ToHex(tt.in.i, tt.in.bitLength)
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestToUint8Slice(t *testing.T) {
	type input struct {
		i            *big.Int
		bitLength    int
		littleEndian bool
	}
	tests := []struct {
		in  input
		out []uint8
	}{
		{
			input{
				big.NewInt(1193046),
				-1,
				false,
			},
			[]uint8{0x12, 0x34, 0x56},
		},
		{
			input{
				big.NewInt(1193046),
				32,
				false,
			},
			[]uint8{0x00, 0x12, 0x34, 0x56},
		},
		{
			input{
				big.NewInt(1193046),
				32,
				true,
			},
			[]uint8{0x56, 0x34, 0x12, 0x00},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := ToUint8Slice(tt.in.i, tt.in.bitLength, tt.in.littleEndian)
			if !reflect.DeepEqual(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}
