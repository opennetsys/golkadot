package u8compact

import (
	"fmt"
	"math/big"
	"testing"
)

func TestFromUint8Slice(t *testing.T) {
	type input struct {
		value     []uint8
		bigLength int
	}

	type output struct {
		offset  int
		encoded *big.Int
	}

	for i, tt := range []struct {
		in  input
		out output
	}{
		{
			input{[]uint8{0xFC}, 32},
			output{1, big.NewInt(63)},
		},
		{
			input{[]uint8{0xFD, 0x7}, 32},
			output{2, big.NewInt(511)},
		},
		{
			input{[]uint8{0xFE, 0xFF, 0x3, 0x0}, 32},
			output{4, big.NewInt(65535)}, // 0xffff
		},
		{
			input{[]uint8{0x3, 0xF9, 0xFF, 0xFF, 0xFF}, 64},
			output{9, big.NewInt(4294967289)}, // 0xfffffff9
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			offset, encoded := FromUint8Slice(tt.in.value, tt.in.bigLength)
			if offset != tt.out.offset {
				t.Errorf("want %v; got %v", tt.out.offset, offset)
			}
			if encoded.String() != tt.out.encoded.String() {
				t.Errorf("want %s; got %s", tt.out.encoded, encoded)
			}
		})
	}
}
