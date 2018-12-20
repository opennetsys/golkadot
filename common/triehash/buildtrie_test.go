package triehash

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuildTrie(t *testing.T) {
	type input struct {
		input  [][][]uint8
		cursor int
	}
	for i, tt := range []struct {
		in  input
		out []uint8
	}{
		{input{[][][]uint8{}, 0}, []uint8{0}},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := BuildTrie(tt.in.input, tt.in.cursor)
			if !reflect.DeepEqual(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}
