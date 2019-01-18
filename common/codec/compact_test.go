package codec

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestEncodeToCompact(t *testing.T) {
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
