package triehash

import (
	"fmt"
	"testing"

	"github.com/c3systems/go-substrate/common/chainspec"
	"github.com/c3systems/go-substrate/common/hexutil"
	"github.com/c3systems/go-substrate/common/u8util"
)

func TestTrieRoot(t *testing.T) {
	var pairs []*TriePair
	for k, v := range chainspec.BBQBirch.Genesis.Raw {
		pairs = append(pairs, &TriePair{
			K: hexToU8a(k),
			V: hexToU8a(v),
		})
	}

	// TODO
	hx := u8util.ToHex(TrieRoot(pairs), -1, true)
	fmt.Println("H", hx)
	fmt.Println("R", chainspec.BBQBirch.GenesisRoot)

	/*
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
	*/
}

// helper
func hexToU8a(s string) []uint8 {
	u8, err := hexutil.ToUint8Slice(s, -1)
	if err != nil {
		panic(err)
	}

	return u8
}
