package bnutil

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestFromHex(t *testing.T) {
	for i, tt := range []struct {
		in  string
		out *big.Int
	}{
		{"0x14", big.NewInt(20)},
		{"81", big.NewInt(129)},
		{"0x", big.NewInt(0)},
	} {
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
	for i, tt := range []struct {
		in  interface{}
		le  bool
		out *big.Int
	}{
		{"14", false, big.NewInt(14)},
		{81, false, big.NewInt(81)},
		{float64(21), false, big.NewInt(21)},
		{[]byte{0, 64, 122, 16, 243, 90}, true, ToBN("100000000000000", false)},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := ToBN(tt.in, tt.le)
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
	for i, tt := range []struct {
		in  input
		out string
	}{
		{input{big.NewInt(128), -1}, "0x80"},
		{input{big.NewInt(81), -1}, "0x51"},
		{input{big.NewInt(21), -1}, "0x15"},
		{input{big.NewInt(128), 16}, "0x0080"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := ToHex(tt.in.i, tt.in.bitLength, false)
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
		isNegative   bool
	}

	for i, tt := range []struct {
		in  input
		out []uint8
	}{
		{
			input{big.NewInt(2045), 16, true, false},
			[]uint8{0xFD, 0x7},
		},
		{
			input{big.NewInt(1193046), -1, false, false}, // 0x123456
			[]uint8{0x12, 0x34, 0x56},
		},
		{
			input{big.NewInt(1193046), 32, false, false},
			[]uint8{0x00, 0x12, 0x34, 0x56},
		},
		{
			input{big.NewInt(1193046), 32, true, false},
			[]uint8{0x56, 0x34, 0x12, 0x00},
		},
		{
			input{big.NewInt(1234), 32, false, false},
			[]uint8{0, 0, 4, 210},
		},
		{
			input{big.NewInt(-1234), -1, true, true},
			[]uint8{46, 251},
		},
		{
			input{big.NewInt(-1234), -1, false, true},
			[]uint8{251, 46},
		},
		{
			input{big.NewInt(-1234), 32, true, true},
			[]uint8{46, 251, 255, 255},
		},
		{
			input{big.NewInt(-1234), 32, false, true},
			[]uint8{255, 255, 251, 46},
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := ToUint8Slice(tt.in.i, tt.in.bitLength, tt.in.littleEndian, tt.in.isNegative)
			if !reflect.DeepEqual(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}
