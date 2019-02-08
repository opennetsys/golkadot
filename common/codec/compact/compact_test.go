package codec

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	codectypes "github.com/opennetsys/golkadot/common/codec/types"
	"github.com/opennetsys/golkadot/common/u8util"
)

func TestBNToCompact(t *testing.T) {
	bi, _ := big.NewInt(0).SetString("18446744073709551615", 10)
	for i, tt := range []struct {
		in  *big.Int
		out codectypes.Compact
	}{
		{
			big.NewInt(18),
			codectypes.Compact([]byte{18 << 2}),
		},
		{
			big.NewInt(63),
			codectypes.Compact([]byte{252}),
		},
		{
			big.NewInt(511),
			codectypes.Compact([]byte{253, 7}),
		},
		{
			big.NewInt(111),
			codectypes.Compact([]byte{189, 1}),
		},
		{
			big.NewInt(65535),
			codectypes.Compact([]byte{254, 255, 3, 0}),
		},
		{
			big.NewInt(4294967289),
			codectypes.Compact([]byte{3, 249, 255, 255, 255}),
		},
		{
			big.NewInt(100000000000000),
			codectypes.Compact([]byte{11, 0, 64, 122, 16, 243, 90}),
		},
		// note: https://github.com/paritytech/parity-codec/blob/master/src/codec.rs
		{
			big.NewInt(0),
			codectypes.Compact([]byte{0}),
		},
		{
			big.NewInt(63),
			codectypes.Compact([]byte{252}),
		},
		{
			big.NewInt(64),
			codectypes.Compact([]byte{1, 1}),
		},
		{
			big.NewInt(16383),
			codectypes.Compact([]byte{253, 255}),
		},
		{
			big.NewInt(16384),
			codectypes.Compact([]byte{2, 0, 1, 0}),
		},
		{
			big.NewInt(1073741823),
			codectypes.Compact([]byte{254, 255, 255, 255}),
		},
		{
			big.NewInt(1073741824),
			codectypes.Compact([]byte{3, 0, 0, 0, 64}),
		},
		{
			big.NewInt(4294967295),
			codectypes.Compact([]byte{3, 255, 255, 255, 255}),
		},
		{
			big.NewInt(4294967296),
			codectypes.Compact([]byte{7, 0, 0, 0, 0, 1}),
		},
		{
			big.NewInt(1099511627776),
			codectypes.Compact([]byte{11, 0, 0, 0, 0, 0, 1}),
		},
		{
			big.NewInt(281474976710656),
			codectypes.Compact([]byte{15, 0, 0, 0, 0, 0, 0, 1}),
		},
		{
			big.NewInt(72057594037927935),
			codectypes.Compact([]byte{15, 255, 255, 255, 255, 255, 255, 255}),
		},
		{
			big.NewInt(72057594037927936),
			codectypes.Compact([]byte{19, 0, 0, 0, 0, 0, 0, 0, 1}),
		},
		{
			bi,
			codectypes.Compact([]byte{19, 255, 255, 255, 255, 255, 255, 255, 255}),
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			comp, err := BNToCompact(tt.in)
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

func TestCompactMetaFromBytes(t *testing.T) {
	type out struct {
		Offset int
		Length string
	}

	for i, tt := range []struct {
		in  []byte
		out out
	}{
		{
			[]byte{252},
			out{
				Offset: 1,
				Length: "63",
			},
		},
		{
			[]byte{253, 7},
			out{
				Offset: 2,
				Length: "511",
			},
		},
		{
			[]byte{254, 255, 3, 0},
			out{
				Offset: 4,
				Length: "65535",
			},
		},
		{
			[]byte{3, 249, 255, 255, 255},
			out{
				Offset: 5,
				Length: "4294967289",
			},
		},
		// note: fails...
		// see #36 https://github.com/opennetsys/golkadot/issues/36
		{
			u8util.FromHex("0x0b00407a10f35a"),
			out{
				Offset: 7,
				Length: "100000000000000",
			},
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			cm, err := CompactMetaFromBytes(tt.in)
			if err != nil || cm == nil {
				t.Error(err)
			}

			if tt.out.Offset != cm.Offset || tt.out.Length != cm.Length.String() {
				t.Errorf("want offset %d length %s; got offset %d length %v", tt.out.Offset, tt.out.Length, cm.Offset, cm.Length.String())
			}
		})
	}
}
