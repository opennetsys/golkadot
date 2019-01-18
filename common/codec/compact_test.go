package codec

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestEncodeToCompact(t *testing.T) {
	bi, _ := big.NewInt(0).SetString("18446744073709551615", 10)
	for i, tt := range []struct {
		in  *big.Int
		out Compact
	}{
		{
			big.NewInt(18),
			Compact([]byte{18 << 2}),
		},
		{
			big.NewInt(63),
			Compact([]byte{252}),
		},
		{
			big.NewInt(511),
			Compact([]byte{253, 7}),
		},
		{
			big.NewInt(111),
			Compact([]byte{189, 1}),
		},
		{
			big.NewInt(65535),
			Compact([]byte{254, 255, 3, 0}),
		},
		{
			big.NewInt(4294967289),
			Compact([]byte{3, 249, 255, 255, 255}),
		},
		{
			big.NewInt(100000000000000),
			Compact([]byte{11, 0, 64, 122, 16, 243, 90}),
		},
		// note: https://github.com/paritytech/parity-codec/blob/master/src/codec.rs
		{
			big.NewInt(0),
			Compact([]byte{0}),
		},
		{
			big.NewInt(63),
			Compact([]byte{252}),
		},
		{
			big.NewInt(64),
			Compact([]byte{1, 1}),
		},
		{
			big.NewInt(16383),
			Compact([]byte{253, 255}),
		},
		{
			big.NewInt(16384),
			Compact([]byte{2, 0, 1, 0}),
		},
		{
			big.NewInt(1073741823),
			Compact([]byte{254, 255, 255, 255}),
		},
		{
			big.NewInt(1073741824),
			Compact([]byte{3, 0, 0, 0, 64}),
		},
		{
			big.NewInt(4294967295),
			Compact([]byte{3, 255, 255, 255, 255}),
		},
		{
			big.NewInt(4294967296),
			Compact([]byte{7, 0, 0, 0, 0, 1}),
		},
		{
			big.NewInt(1099511627776),
			Compact([]byte{11, 0, 0, 0, 0, 0, 1}),
		},
		{
			big.NewInt(281474976710656),
			Compact([]byte{15, 0, 0, 0, 0, 0, 0, 1}),
		},
		{
			big.NewInt(72057594037927935),
			Compact([]byte{15, 255, 255, 255, 255, 255, 255, 255}),
		},
		{
			big.NewInt(72057594037927936),
			Compact([]byte{19, 0, 0, 0, 0, 0, 0, 0, 1}),
		},
		{
			bi,
			Compact([]byte{19, 255, 255, 255, 255, 255, 255, 255, 255}),
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			comp, err := EncodeToCompact(tt.in)
			if err != nil || comp == nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(*comp, tt.out) {
				t.Errorf("want %v; got %v", tt.out, comp)
			}
		})
	}
}

func TestBigEToLittleE(t *testing.T) {
	for i, tt := range []struct {
		in  []byte
		out []byte
	}{
		{
			nil,
			nil,
		},
		{
			[]byte{},
			[]byte{},
		},
		{
			[]byte{0},
			[]byte{0},
		},
		{
			[]byte{0, 1},
			[]byte{1, 0},
		},
		{
			[]byte{0, 1, 2},
			[]byte{2, 1, 0},
		},
		{
			[]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]byte{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			bigEToLittleE(tt.in)

			if !reflect.DeepEqual(tt.in, tt.out) {
				t.Errorf("want %v; got %v", tt.out, tt.in)
			}
		})
	}
}
