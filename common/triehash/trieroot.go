package triehash

import (
	"math/big"

	"github.com/c3systems/go-substrate/common/crypto"
	"github.com/c3systems/go-substrate/common/triecodec"
	"github.com/c3systems/go-substrate/common/u8compact"
)

// TrieRoot creates a trie hash from the supplied pairs.
func TrieRoot(input []*TriePair) crypto.Blake2b256Hash {
	return triecodec.Hashing(
		UnhashedTrie(input),
	)
}

// TrieRootOrdered creates a trie hash from the supplied pairs.
func TrieRootOrdered(input [][]uint8) crypto.Blake2b256Hash {
	var values []*TriePair
	for index, value := range input {
		values = append(values, &TriePair{
			K: u8compact.CompactToUint8Slice(big.NewInt(int64(index)), 32),
			V: value,
		})

	}

	return TrieRoot(values)
}
