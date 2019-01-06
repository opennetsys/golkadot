package triedb

import (
	"testing"
)

// TODO

func TestSnapshots(t *testing.T) {
	/*
		trie := NewTrie()
		back := NewTrie()

		values := []*triehash.Trie{
			&triehash.Trie{
				K: stringutil.ToUint8Slice("test"),
				V: stringutil.ToUint8Slice("one"),
			},
		}

		root := triehash.TrieRoot(values)

		trie.Put(values[0].K, values[0].V)
		trie.Put(values[0].K, stringutil.ToUint8Slice("two"))
		trie.Del(values[0].K)
		trie.Put(values[0].K, values[0].V)
		trie.Put(stringutil.ToUint8Slice("doge"), stringutil.ToUint8Slice("coin"))
		trie.Del(stringutil.ToUint8Slice("doge"))

		trie.Snapshot(back, func() {})

		if hex.EncodeToString(back.GetRoot()) != hex.EncodeToString(root) {
			t.Fatal("not match")
		}

		if !reflect.DeepEqual(trie.Get(values[0].K), values[0].V) {
			t.Fatal("not match")
		}
	*/
}
