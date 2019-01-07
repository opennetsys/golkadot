package triedb

import (
	"reflect"
	"testing"

	"github.com/c3systems/go-substrate/common/db"
	"github.com/c3systems/go-substrate/common/triehash"
)

func newTrie() *Trie {
	memdb := db.NewMemoryDB(&db.BaseOptions{})
	basedb := db.BaseDB(memdb)
	txdbt := db.NewTransactionDB(&basedb)
	txdb := db.TXDB(txdbt)
	trie := NewTrie(txdb, nil)
	return trie
}

func TestSnapshots(t *testing.T) {
	t.Run("creates a snapshot of the (relevant) trie data", func(t *testing.T) {
		trie := newTrie()
		back := newTrie()

		values := []*triehash.TriePair{
			&triehash.TriePair{K: []uint8("test"), V: []uint8("one")},
		}

		//root := triehash.TrieRoot(values)

		trie.Put(values[0].K, values[0].V)
		trie.Put(values[0].K, []uint8("two"))
		trie.Del(values[0].K)
		trie.Put(values[0].K, values[0].V)
		trie.Put([]uint8("doge"), []uint8("coin"))
		trie.Del([]uint8("doge"))

		trie.Snapshot(back, nil)

		//fmt.Println("back root", back.GetRoot())
		//fmt.Println("trie root", trie.GetRoot())
		//fmt.Println("triehash root", root)

		// TODO: fix
		/*
			if hex.EncodeToString(back.GetRoot()) != hex.EncodeToString(root) {
				t.Fail(
			}
		*/

		if !reflect.DeepEqual(trie.Get(values[0].K), values[0].V) {
			t.Fail()
		}
	})

	t.Run("creates a snapshot of the (relevant) data", func(t *testing.T) {
		trie := newTrie()
		back := newTrie()

		values := []*triehash.TriePair{
			&triehash.TriePair{K: []uint8("one"), V: []uint8("testing")},
			&triehash.TriePair{K: []uint8("two"), V: []uint8("testing with a much longer value here")},
			&triehash.TriePair{K: []uint8("twzei"), V: []uint8("und Deutch")},
			&triehash.TriePair{K: []uint8("do"), V: []uint8("do it")},
			&triehash.TriePair{K: []uint8("dog"), V: []uint8("doggie")},
			&triehash.TriePair{K: []uint8("dogge"), V: []uint8("bigger doge")},
			&triehash.TriePair{K: []uint8("dodge"), V: []uint8("coin")},
		}

		//root := triehash.TrieRoot(values)

		for _, value := range values {
			trie.Put(value.K, value.V)
		}

		trie.Snapshot(back, nil)

		// TODO: fix
		/*
			if hex.EncodeToString(back.GetRoot()) != hex.EncodeToString(root) {
				t.Fail()
			}
		*/

		for _, value := range values {
			if !reflect.DeepEqual(trie.Get(value.K), value.V) {
				t.Fail()
			}
		}
	})
}
