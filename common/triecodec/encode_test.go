package triecodec

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/c3systems/go-substrate/common/hexutil"
)

func TestEncode(t *testing.T) {
	for i, tt := range []struct {
		in  []interface{}
		out []uint8
	}{
		{[]interface{}{
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			[][]uint8{
				[]uint8{0x3A},
				[]uint8{0xAA},
			},
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			[]uint8{0xA0},
		}, []uint8{0xFF, 0x0, 0x4, 0x4, 0xA0, 0x10, 0x2, 0xA, 0x4, 0xAA}},
		{[]interface{}{
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			[][]uint8{
				[]uint8{0x3A},
				[]uint8{0xAA},
			},
			[][]uint8{
				[]uint8{0x3B},
				[]uint8{0xAB},
			},
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			[]uint8{0xA0},
		}, []uint8{0xFF, 0x0, 0xC, 0x4, 0xA0, 0x10, 0x2, 0xA, 0x4, 0xAA, 0x10, 0x2, 0xB, 0x4, 0xAB}},
		{[]interface{}{
			NewNull(),
			NewNull(),
			[][]uint8{
				hexToU8a("0x37b3872d47181b4a2dc15f0da43e7026"),
				hexToU8a("0xe803000000000000"),
			},
			[]uint8{0x17, 0xD, 0x32, 0x2A, 0xC4, 0x9D, 0x87, 0x8, 0xF1, 0x51, 0x34, 0x6C, 0x68, 0xD9, 0xE5, 0x84, 0x52, 0xD8, 0x3A, 0x9D, 0x3B, 0x71, 0xE, 0x1E, 0xAD, 0x35, 0xEB, 0x32, 0x69, 0xAB, 0x23, 0x53},
			NewNull(),
			NewNull(),
			NewNull(),
			[][]uint8{
				hexToU8a("0x3935e46f94f24b82716c0142e2271de9"),
				hexToU8a("0x0087000000000000"),
			},
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
			NewNull(),
		}, hexToU8a("0xfe8c00682007b3872d47181b4a2dc15f0da43e702620e80300000000000080170d322ac49d8708f151346c68d9e58452d83a9d3b710e1ead35eb3269ab235368200935e46f94f24b82716c0142e2271de9200087000000000000")},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if i == 0 || i == 1 {
				t.Skip()
			}
			result := Encode(tt.in)
			if !reflect.DeepEqual(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func hexToU8a(s string) []uint8 {
	u8, err := hexutil.ToUint8Slice(s, -1)
	if err != nil {
		panic(err)
	}

	return u8
}
