package triehash

import "github.com/c3systems/go-substrate/common/triecodec"

// TrieRoot creates a trie hash from the supplied pairs.
func TrieRoot(input []*TriePair) []uint8 {
	return triecodec.Hashing(
		UnhashedTrie(input),
	)
}
