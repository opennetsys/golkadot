package triecodec

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFromNibbles(t *testing.T) {
	for i, tt := range []struct {
		in  []uint8
		out []uint8
	}{
		{[]uint8{0x4, 0x1, 0x2, 0x0}, []uint8{0x41, 0x20}},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := FromNibbles(tt.in)
			if !reflect.DeepEqual(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}
