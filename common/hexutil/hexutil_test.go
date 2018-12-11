package hexutil

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestHasPrefix(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			"12",
			false,
		},
		{
			"0x12",
			true,
		},
		{
			"0x",
			true,
		},
		{
			"0",
			false,
		},
		{
			"",
			false,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := HasPrefix(tt.in)
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestValidHex(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			"12",
			true,
		},
		{
			"0x12",
			true,
		},
		{
			"0x",
			true,
		},
		{
			"0",
			true,
		},
		{
			"",
			false,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := ValidHex(tt.in)
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestAddPrefix(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"12",
			"0x12",
		},
		{
			"123",
			"0x0123",
		},
		{
			"0x123",
			"0x123",
		},
		{
			"",
			"0x",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := AddPrefix(tt.in)
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestStripPrefix(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"0x123",
			"123",
		},
		{
			"123",
			"123",
		},
		{
			"0x",
			"",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := StripPrefix(tt.in)
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestHexFixLength(t *testing.T) {
	type input struct {
		hexStr      string
		bitLength   int
		withPadding bool
	}
	tests := []struct {
		in  input
		out string
	}{
		{
			input{
				"0x12",
				16,
				false,
			},
			"0x12",
		},
		{
			input{
				"0x12",
				16,
				true,
			},
			"0x0012",
		},
		{
			input{
				"0x0012",
				8,
				false,
			},
			"0x12",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := HexFixLength(tt.in.hexStr, tt.in.bitLength, tt.in.withPadding)
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestToBN(t *testing.T) {
	type input struct {
		hexStr       string
		littleEndian bool
	}

	tests := []struct {
		in  input
		out *big.Int
	}{
		{
			input{
				"0x14",
				false,
			},
			big.NewInt(20),
		},
		{
			input{
				"14",
				false,
			},
			big.NewInt(20),
		},
		{
			input{
				"0x14",
				true,
			},
			big.NewInt(20),
		},
		{
			input{
				"14",
				true,
			},
			big.NewInt(20),
		},
		{
			input{
				"81",
				false,
			},
			big.NewInt(129),
		},
		{
			input{
				"0x",
				true,
			},
			big.NewInt(0),
		},
		{
			input{
				"0x4500000000000000",
				true,
			},
			big.NewInt(69),
		},
		{
			input{
				"0x0000000000000100",
				false,
			},
			big.NewInt(256),
		},
		{
			input{
				"0x0001000000000000",
				true,
			},
			big.NewInt(256),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := ToBN(tt.in.hexStr, tt.in.littleEndian)
			if err != nil {
				t.Error(err)
			}
			if result.String() != tt.out.String() {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestToUint8Slice(t *testing.T) {
	type input struct {
		hexStr    string
		bitLength int
	}

	tests := []struct {
		in  input
		out []uint8
	}{
		{
			input{
				"0x80001f",
				-1,
			},
			[]uint8{0x80, 0x00, 0x1f},
		},
		{
			input{
				"0x80001f",
				32,
			},
			[]uint8{0x00, 0x80, 0x00, 0x1f},
		},
		{
			input{
				"0x80000a",
				-1,
			},
			[]uint8{128, 0, 10},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := ToUint8Slice(tt.in.hexStr, tt.in.bitLength)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}
