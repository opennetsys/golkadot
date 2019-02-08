package triehash

import (
	"encoding/hex"
	"sort"

	"github.com/opennetsys/golkadot/common/triecodec"
)

// UnhashedTrie ...
func UnhashedTrie(input []*TriePair) []uint8 {
	result := make(map[string]*TriePair)

	for index := 0; index < len(input); index++ {
		result[hex.EncodeToString(input[index].K)] = input[index]
	}

	keylist := make([]string, 0)
	for k := range result {
		keylist = append(keylist, k)
	}

	sort.Strings(keylist)

	pairs := make([][][]uint8, len(result))
	i := 0
	for _, k := range keylist {
		pair := result[k]
		pairs[i] = make([][]uint8, 2)
		pairs[i][0] = triecodec.ToNibbles(pair.K)
		pairs[i][1] = pair.V
		i++
	}

	return BuildTrie(pairs, 0)
}
