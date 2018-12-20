package triehash

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/c3systems/go-substrate/common/chainspec"
	"github.com/c3systems/go-substrate/common/stringutil"
	"github.com/c3systems/go-substrate/common/u8util"
)

func TestTrieRoot(t *testing.T) {
	var pairs0 []*TriePair
	for k, v := range chainspec.BBQBirch.Genesis.Raw {
		pairs0 = append(pairs0, &TriePair{
			K: hexToU8a(k),
			V: hexToU8a(v),
		})
	}

	for i, tt := range []struct {
		in  []*TriePair
		out string
	}{
		{pairs0, chainspec.BBQBirch.GenesisRoot},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := TrieRoot(tt.in)
			if u8util.ToHex(result, -1, true) != tt.out {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}

func TestTrieRootOrdered(t *testing.T) {
	for i, tt := range []struct {
		in  [][]uint8
		out []uint8
	}{
		{
			[][]uint8{
				stringutil.ToUint8Slice("doe"),
				stringutil.ToUint8Slice("reindeer"),
			},
			u8util.FromHex("0xb9b1bb07e481f0393e15f32f34abd665f7a698786a7ec9feb31b2e8927ad5f86"),
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := TrieRootOrdered(tt.in)
			if !reflect.DeepEqual(result, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result)
			}
		})
	}
}
