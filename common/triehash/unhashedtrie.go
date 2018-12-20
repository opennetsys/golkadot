package triehash

import (
	"encoding/hex"
	"fmt"

	"github.com/c3systems/go-substrate/common/triecodec"
	"github.com/c3systems/go-substrate/common/u8util"
)

// UnhashedTrie ...
func UnhashedTrie(input []*TriePair) []uint8 {
	result := make(map[string]*TriePair)

	for index := 0; index < len(input); index++ {
		result[u8util.ToHex(input[index].K, 32, true)] = input[index]
	}

	pairs := make([][][]uint8, len(result))
	// TODO: sort by key before appending
	i := 0
	for k := range result {
		pair := result[k]
		pairs[i] = make([][]uint8, 2)
		pairs[i][0] = triecodec.ToNibbles(pair.K)
		pairs[i][1] = pair.V
		fmt.Println(hex.EncodeToString(pair.K))
		i++
	}

	return BuildTrie(pairs, 0)
}
