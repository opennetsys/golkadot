package triedb

import (
	"reflect"
	"testing"

	"github.com/c3systems/go-substrate/common/db"
)

func newTrie(codec InterfaceCodec) *Trie {
	memdb := db.NewMemoryDB(&db.BaseOptions{})
	basedb := db.BaseDB(memdb)
	txdbt := db.NewTransactionDB(&basedb)
	txdb := db.TXDB(txdbt)
	trie := NewTrie(txdb, nil, codec)
	return trie
}

func TestTrieDB(t *testing.T) {
	codec := NewRLPCodec()

	t.Run("test 1: simple save and retrieve", func(t *testing.T) {
		trie := newTrie(codec)

		t.Run("starts with a valid root", func(t *testing.T) {
			root := trie.GetRoot()
			expectedRoot := []uint8{3, 23, 10, 46, 117, 151, 183, 183, 227, 216, 76, 5, 57, 29, 19, 154, 98, 177, 87, 231, 135, 134, 216, 192, 130, 242, 157, 207, 76, 17, 19, 20}
			if !reflect.DeepEqual(root, expectedRoot) {
				t.Fail()
			}
		})

		t.Run("save a value", func(t *testing.T) {
			trie.Put([]uint8("test"), []uint8("one"))

			root := trie.GetRoot()
			expectedRoot := []uint8{205, 153, 117, 166, 91, 174, 139, 99, 166, 27, 129, 69, 109, 204, 32, 65, 229, 200, 16, 4, 231, 103, 24, 40, 229, 140, 75, 86, 49, 96, 92, 23}
			if !reflect.DeepEqual(root, expectedRoot) {
				t.Fail()
			}
		})

		t.Run("should get a value", func(t *testing.T) {
			if !reflect.DeepEqual(trie.Get([]uint8("test")), []uint8("one")) {
				t.Fail()
			}
		})

		t.Run("should update a value", func(t *testing.T) {
			trie.Put([]uint8("test"), []uint8("two"))
			if !reflect.DeepEqual(trie.Get([]uint8("test")), []uint8("two")) {
				t.Fail()
			}

			root := trie.GetRoot()
			expectedRoot := []uint8{242, 194, 105, 66, 186, 10, 101, 45, 41, 26, 169, 239, 28, 230, 66, 122, 84, 130, 77, 233, 242, 119, 172, 196, 48, 76, 47, 184, 208, 173, 140, 83}
			if !reflect.DeepEqual(root, expectedRoot) {
				t.Fail()
			}
		})

		t.Run("should delete a value", func(t *testing.T) {
			trie.Del([]uint8("test"))
			root := trie.GetRoot()
			expectedRoot := []uint8{3, 23, 10, 46, 117, 151, 183, 183, 227, 216, 76, 5, 57, 29, 19, 154, 98, 177, 87, 231, 135, 134, 216, 192, 130, 242, 157, 207, 76, 17, 19, 20}
			if !reflect.DeepEqual(root, expectedRoot) {
				t.Fail()
			}
		})

		t.Run("should recreate a value", func(t *testing.T) {
			trie.Put([]uint8("test"), []uint8("one"))

			root := trie.GetRoot()
			expectedRoot := []uint8{205, 153, 117, 166, 91, 174, 139, 99, 166, 27, 129, 69, 109, 204, 32, 65, 229, 200, 16, 4, 231, 103, 24, 40, 229, 140, 75, 86, 49, 96, 92, 23}
			if !reflect.DeepEqual(root, expectedRoot) {
				t.Fail()
			}
		})

		t.Run("should get updated a value", func(t *testing.T) {
			if !reflect.DeepEqual(trie.Get([]uint8("test")), []uint8("one")) {
				t.Fail()
			}
		})

		t.Run("should create a branch here", func(t *testing.T) {
			trie.Put([]uint8("doge"), []uint8("coin"))
			root := trie.GetRoot()
			expectedRoot := []uint8{183, 144, 27, 71, 56, 21, 57, 163, 157, 65, 172, 123, 91, 199, 80, 175, 117, 39, 202, 240, 188, 37, 153, 92, 132, 180, 224, 112, 180, 14, 106, 18}
			if !reflect.DeepEqual(root, expectedRoot) {
				t.Fail()
			}
		})

		t.Run("should get a value that is in a branch", func(t *testing.T) {
			if !reflect.DeepEqual(trie.Get([]uint8("doge")), []uint8("coin")) {
				t.Fail()
			}
		})

		t.Run("should delete from a branch", func(t *testing.T) {
			trie.Del([]uint8("doge"))
			if trie.Get([]uint8("doge")) != nil {
				t.Fail()
			}
			root := trie.GetRoot()
			expectedRoot := []uint8{205, 153, 117, 166, 91, 174, 139, 99, 166, 27, 129, 69, 109, 204, 32, 65, 229, 200, 16, 4, 231, 103, 24, 40, 229, 140, 75, 86, 49, 96, 92, 23}
			if !reflect.DeepEqual(root, expectedRoot) {
				t.Fail()
			}
		})
	})

	t.Run("test 2: storing longer values", func(t *testing.T) {
		trie := newTrie(codec)

		longString := "this will be a really really really long value"
		longStringRoot := []uint8{23, 204, 252, 0, 51, 163, 54, 163, 91, 210, 76, 17, 64, 20, 221, 47, 231, 80, 223, 210, 146, 205, 224, 233, 94, 124, 55, 100, 172, 218, 10, 231}

		t.Run("should store a longer string", func(t *testing.T) {
			trie.Put([]uint8("done"), []uint8(longString))

			root := trie.GetRoot()
			expectedRoot := []uint8{148, 127, 147, 249, 249, 250, 169, 115, 16, 185, 79, 81, 241, 124, 81, 180, 253, 119, 188, 217, 101, 135, 135, 112, 81, 98, 213, 176, 126, 136, 90, 210}
			if !reflect.DeepEqual(root, expectedRoot) {
				t.Fail()
			}
		})

		t.Run("should retrieve a longer value (first pass)", func(t *testing.T) {
			if !reflect.DeepEqual(trie.Get([]uint8("done")), []uint8(longString)) {
				t.Fail()
			}
		})

		t.Run("should store subsequent values", func(t *testing.T) {
			trie.Put([]uint8("doge"), []uint8("coin"))

			root := trie.GetRoot()
			if !reflect.DeepEqual(root, longStringRoot) {
				t.Fail()
			}
		})

		t.Run("should retrieve a longer value (second pass)", func(t *testing.T) {
			if !reflect.DeepEqual(trie.Get([]uint8("done")), []uint8(longString)) {
				t.Fail()
			}
		})

		t.Run("should retrieve subsequent values", func(t *testing.T) {
			if !reflect.DeepEqual(trie.Get([]uint8("doge")), []uint8("coin")) {
				t.Fail()
			}
		})

		t.Run("should when being modified delete the old value", func(t *testing.T) {
			trie.Put([]uint8("done"), []uint8("test"))
			root := trie.GetRoot()
			expectedRoot := []uint8{122, 108, 225, 206, 60, 150, 134, 74, 53, 137, 106, 42, 243, 75, 45, 208, 46, 105, 74, 189, 67, 167, 13, 149, 141, 126, 151, 245, 229, 200, 119, 220}
			if !reflect.DeepEqual(root, expectedRoot) {
				t.Fail()
			}
		})

	})

	t.Run("test 3: testing Extentions and branches", func(t *testing.T) {
		trie := newTrie(codec)

		t.Run("should store a value", func(t *testing.T) {
			trie.Put([]uint8("doge"), []uint8("coin"))
		})
		t.Run("should create extention to store this value", func(t *testing.T) {
			trie.Put([]uint8("do"), []uint8("verb"))
			if !reflect.DeepEqual(trie.Get([]uint8("do")), []uint8("verb")) {
				t.Fail()
			}

			root := trie.GetRoot()
			expectedRoot := []uint8{64, 150, 109, 76, 24, 123, 197, 7, 75, 83, 149, 223, 204, 7, 19, 117, 211, 36, 195, 240, 236, 214, 197, 81, 230, 7, 166, 75, 213, 246, 179, 19}
			if !reflect.DeepEqual(root, expectedRoot) {
				t.Fail()
			}
		})
		t.Run("should store this value under the extention", func(t *testing.T) {
			trie.Put([]uint8("done"), []uint8("finished"))
			root := trie.GetRoot()
			expectedRoot := []uint8{192, 214, 81, 170, 221, 82, 60, 25, 190, 123, 112, 7, 138, 253, 63, 178, 198, 192, 194, 173, 133, 193, 240, 169, 194, 176, 141, 106, 11, 13, 117, 97}
			if !reflect.DeepEqual(root, expectedRoot) {
				t.Fail()
			}
		})
	})

	t.Run("test 4: testing Extentions and branches - reverse", func(t *testing.T) {
		trie := newTrie(codec)

		t.Run("should create extention to store this value", func(t *testing.T) {
			trie.Put([]uint8("do"), []uint8("verb"))
			if !reflect.DeepEqual(trie.Get([]uint8("do")), []uint8("verb")) {
				t.Fail()
			}
		})

		t.Run("should store a value", func(t *testing.T) {
			trie.Put([]uint8("doge"), []uint8("coin"))
		})

		t.Run("should store this value under the extention", func(t *testing.T) {
			trie.Put([]uint8("done"), []uint8("finished"))
			root := trie.GetRoot()
			expectedRoot := []uint8{192, 214, 81, 170, 221, 82, 60, 25, 190, 123, 112, 7, 138, 253, 63, 178, 198, 192, 194, 173, 133, 193, 240, 169, 194, 176, 141, 106, 11, 13, 117, 97}
			if !reflect.DeepEqual(root, expectedRoot) {
				t.Fail()
			}
		})
	})

	t.Run("test 5: testing deletions cases", func(t *testing.T) {
		trie := newTrie(codec)

		t.Run("should delete from a branch->branch-branch", func(t *testing.T) {
			trie.Put([]uint8{11, 11, 11}, []uint8("first"))
			trie.Put([]uint8{12, 22, 22}, []uint8("create the first branch"))

			trie.Put([]uint8{12, 34, 44}, []uint8("create the last branch"))
			trie.Del([]uint8{12, 22, 22})

			if trie.Get([]uint8{12, 22, 22}) != nil {
				t.Fail()
			}
		})

		t.Run("should delete from a branch->branch-extension", func(t *testing.T) {
			trie.Put([]uint8{11, 11, 11}, []uint8("first"))
			trie.Put([]uint8{12, 22, 22}, []uint8("create the first branch"))
			trie.Put([]uint8{12, 33, 33}, []uint8("create the middle branch"))
			trie.Put([]uint8{12, 33, 44}, []uint8("create the last branch"))
			trie.Del([]uint8{12, 22, 22})
			if trie.Get([]uint8{12, 22, 22}) != nil {
				t.Fail()
			}
		})

		t.Run("should delete from a extension->branch-extension", func(t *testing.T) {
			trie.Put([]uint8{11, 11, 11}, []uint8("first"))
			trie.Put([]uint8{12, 22, 22}, []uint8("create the first branch"))
			trie.Put([]uint8{12, 33, 33}, []uint8("create the middle branch"))
			trie.Put([]uint8{12, 33, 44}, []uint8("create the last branch"))
			trie.Del([]uint8{11, 11, 11})
			if trie.Get([]uint8{11, 11, 11}) != nil {
				t.Fail()
			}
		})

		t.Run("should delete from a extension->branch-branch", func(t *testing.T) {
			trie.Put([]uint8{11, 11, 11}, []uint8("first"))
			trie.Put([]uint8{12, 22, 22}, []uint8("create the first branch"))
			trie.Put([]uint8{12, 33, 33}, []uint8("create the middle branch"))
			trie.Put([]uint8{12, 34, 44}, []uint8("create the last branch"))
			trie.Del([]uint8{11, 11, 11})
			if trie.Get([]uint8{11, 11, 11}) != nil {
				t.Fail()
			}
		})
	})
}
