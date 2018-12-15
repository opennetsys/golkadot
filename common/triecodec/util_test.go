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
		{[]uint8{0x3, 0x1, 0x2, 0x3, 0x4, 0x5}, []uint8{0x31, 0x23, 0x45}},
		{[]uint8{0x3, 0x1, 0x2, 0x3, 0x4, 0x5, 0x7, 0x8}, []uint8{0x31, 0x23, 0x45, 0x78}},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := FromNibbles(tt.in)
			if !reflect.DeepEqual(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestToNibbles(t *testing.T) {
	for i, tt := range []struct {
		in  []uint8
		out []uint8
	}{
		{[]uint8{0x31, 0x23, 0x45}, []uint8{0x3, 0x1, 0x2, 0x3, 0x4, 0x5}},
		{[]uint8{0x41}, []uint8{0x4, 0x1}},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := ToNibbles(tt.in)
			if !reflect.DeepEqual(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestSharedPrefixLength(t *testing.T) {
	type input struct {
		first  []uint8
		second []uint8
	}
	for i, tt := range []struct {
		in  input
		out int
	}{
		{input{[]uint8{0x1, 0x2, 0x3, 0x4}, []uint8{0x1, 0x2, 0x3}}, 3},
		{input{[]uint8{0x1, 0x2}, []uint8{0x1, 0x2}}, 2},
		{input{[]uint8{}, []uint8{0x1, 0x2}}, 0},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := SharedPrefixLength(tt.in.first, tt.in.second)
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestFuseNibbles(t *testing.T) {
	type input struct {
		nibbles []uint8
		isLeaf  bool
	}
	for i, tt := range []struct {
		in  input
		out []uint8
	}{
		{input{[]uint8{0x1}, false}, []uint8{0x81, 0x1}},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := FuseNibbles(tt.in.nibbles, tt.in.isLeaf)
			if !reflect.DeepEqual(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}
