package triehash

import (
	"fmt"
	"testing"

	"github.com/c3systems/go-substrate/common/chainspec"
	"github.com/c3systems/go-substrate/common/hexutil"
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

// helper
func hexToU8a(s string) []uint8 {
	u8, err := hexutil.ToUint8Slice(s, -1)
	if err != nil {
		panic(err)
	}

	return u8
}
