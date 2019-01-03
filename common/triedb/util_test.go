package triedb

import (
	"fmt"
	"testing"
)

func TestUtils(t *testing.T) {
	// TODO: table tests

	branchNode := NewNodeBranch([]Node{
		NewNodeEmpty(),
		NewNodeEmpty(),
		NewNodeEmpty(),
		NewNodeEmpty(),
		NewNodeEmpty(),
		NewNodeEmpty(),
		NewNodeEmpty(),
		NewNodeEmpty(),
		NewNodeEmpty(),
		NewNodeEmpty(),
		NewNodeEmpty(),
		NewNodeEmpty(),
		NewNodeEmpty(),
		NewNodeEmpty(),
		NewNodeEmpty(),
		NewNodeEmpty(),
		NewNodeEmpty(),
	})

	if IsBranchNode(branchNode) != true {
		t.Error("expected branch node")
	}

	branchNode2 := NewNodeBranch([]Node{
		NewNodeEmpty(),
	})

	if IsBranchNode(branchNode2) != false {
		t.Error("expected not a branch node")
	}
}

func TestKeyEquals(t *testing.T) {
	type input struct {
		key  []uint8
		test []uint8
	}
	for i, tt := range []struct {
		in  input
		out bool
	}{
		{input{[]byte{1}, []byte{1}}, true},
		{input{[]byte{1, 2}, []byte{1, 2}}, true},
		{input{[]byte{1, 2}, []byte{1, 2, 3}}, false},
		{input{[]byte{1, 2, 2}, []byte{1, 2, 3}}, false},
		{input{[]byte{0, 1}, []byte{1}}, false},
		{input{[]byte{0}, []byte{1}}, false},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := KeyEquals(tt.in.key, tt.in.test)
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestKeyStartsWith(t *testing.T) {
	type input struct {
		key     []uint8
		partial []uint8
	}
	for i, tt := range []struct {
		in  input
		out bool
	}{
		{input{[]byte{1}, []byte{1}}, true},
		{input{[]byte{1, 2}, []byte{1, 2}}, true},
		{input{[]byte{1, 2, 3}, []byte{1}}, true},
		{input{[]byte{1, 2, 3}, []byte{1, 2}}, true},
		{input{[]byte{1, 2}, []byte{1, 2, 3}}, false},
		{input{[]byte{1, 2, 2}, []byte{1, 2, 3}}, false},
		{input{[]byte{0, 1}, []byte{1}}, false},
		{input{[]byte{0}, []byte{1}}, false},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := KeyStartsWith(tt.in.key, tt.in.partial)
			if result != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}
