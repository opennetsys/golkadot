package triecodec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	for i, tt := range []struct {
		in  []uint8
		out interface{}
	}{
		{
			[]uint8{0x05, 0x48, 0x19, 0x04, 0xfe},
			[][]uint8{
				[]uint8{0x20, 0x48, 0x19},
				[]uint8{0xfe},
			},
		},
		{
			[]uint8{0xfe, 0x00, 0x0c, 0x48, 0x81, 0x0a, 0x3c, 0xff, 0x00, 0x0c, 0x04, 0xa0, 0x10, 0x02, 0x0a, 0x04,
				0xaa, 0x10, 0x02, 0x0b, 0x04, 0xab, 0x10, 0x02, 0x0b, 0x04, 0xb0},
			[]interface{}{
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				[][]uint8{
					[]uint8{0x1A},
					[]uint8{0xFF, 0x0, 0xC, 0x4, 0xA0, 0x10, 0x2, 0xA, 0x4, 0xAA, 0x10, 0x2, 0xB, 0x4, 0xAB},
				},
				[][]uint8{
					[]uint8{0x3B},
					[]uint8{0xB0},
				},
				nil,
				nil,
				nil,
				nil,
				nil,
			},
		},
		{
			hexToU8a("0xfe8c00682007b3872d47181b4a2dc15f0da43e702620e80300000000000080170d322ac49d8708f151346c68d9e58452d83a9d3b710e1ead35eb3269ab235368200935e46f94f24b82716c0142e2271de9200087000000000000"),
			[]interface{}{
				nil,
				nil,
				[][]uint8{
					hexToU8a("0x37b3872d47181b4a2dc15f0da43e7026"),
					hexToU8a("0xe803000000000000"),
				},
				[]uint8{0x17, 0xD, 0x32, 0x2A, 0xC4, 0x9D, 0x87, 0x8, 0xF1, 0x51, 0x34, 0x6C, 0x68, 0xD9, 0xE5, 0x84, 0x52, 0xD8, 0x3A, 0x9D, 0x3B, 0x71, 0xE, 0x1E, 0xAD, 0x35, 0xEB, 0x32, 0x69, 0xAB, 0x23, 0x53},
				nil,
				nil,
				nil,
				[][]uint8{
					hexToU8a("0x3935e46f94f24b82716c0142e2271de9"),
					hexToU8a("0x0087000000000000"),
				},
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
			},
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := Decode(tt.in)
			js, err := json.Marshal(result)
			if err != nil {
				t.Error(err)
			}
			jsout, err := json.Marshal(tt.out)
			if err != nil {
				t.Error(err)
			}
			if !bytes.Equal(js, jsout) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}
