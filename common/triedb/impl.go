package triedb

import (
	"log"
	"math"

	"github.com/c3systems/go-substrate/common/db"
	"github.com/c3systems/go-substrate/common/triecodec"
	"github.com/c3systems/go-substrate/common/u8util"
)

// TODO: tests (not working yet)

// Impl ...
type Impl struct {
	checkpoint *Checkpoint
	db         db.TXDB
}

// TxDB ...
type TxDB struct {
	db.MemoryDB
}

// NewImpl ...
func NewImpl(db db.TXDB, rootHash []uint8) *Impl {
	checkpoint := NewCheckpoint(rootHash)
	return &Impl{
		checkpoint: checkpoint,
		db:         db,
	}
}

// Snapshot ...
func (i *Impl) Snapshot(dest Trie, fn db.ProgressCB, root []uint8, keys int, percent int, depth int) int {
	node := i.GetNode(root)

	if node == nil {
		return keys
	}

	keys++

	dest.impl.db.Put(root, EncodeNode(node))

	fn(&db.ProgressValue{
		IsCompleted: false,
		Keys:        keys,
		Percent:     percent,
	})

	vals, ok := node.([][]uint8)
	if !ok {
		log.Fatal("not ok")
	}
	for _, val := range vals {
		if val != nil && len(val) == 32 {
			keys = i.Snapshot(dest, fn, val, keys, percent, depth+1)
		}

		percent += int((float64(100) / float64(len(vals))) / math.Pow(float64(16), float64(depth)))
	}

	return keys
}

// GetNode ...
func (i *Impl) GetNode(hash []uint8) Node {
	if hash == nil || len(hash) == 0 || KeyEquals(hash, []uint8{}) {
		return nil
	} else if len(hash) < 32 {
		return DecodeNode(hash)
	}

	return DecodeNode(i.db.Get(hash))
}

// Del ...
func (i *Impl) Del(node Node, trieKey []uint8) Node {
	if IsEmptyNode(node) {
		return nil
	} else if IsBranchNode(node) {
		nodeBranch, ok := node.(NodeBranch)
		if !ok {
			log.Fatal("not ok")
		}
		return i.DelBranchNode(nodeBranch, trieKey)
	} else if IsKvNode(node) {
		nodekv, ok := node.(NodeKv)
		if !ok {
			log.Fatal("not ok")
		}
		return i.DelKvNode(nodekv, trieKey)
	}

	log.Fatal("Unreachable")
	return NewNode(nil)
}

// DelBranchNode ...
func (i *Impl) DelBranchNode(node NodeBranch, trieKey []uint8) Node {
	if len(trieKey) == 0 {
		node[len(node)-1] = nil

		return i.NormalizeBranchNode(node)
	}

	k, ok := node[trieKey[0]].([]uint8)
	if !ok {
		log.Fatal("not ok")
	}
	nodeToDelete := i.GetNode(k)
	subNode := i.Del(nodeToDelete, trieKey[1:])
	encodedSubNode := i.PersistNode(subNode)
	esn, ok := encodedSubNode.([]uint8)
	if !ok {
		log.Fatal("not ok")
	}
	if KeyEquals(esn, k) {
		return node
	}

	node[trieKey[0]] = encodedSubNode

	if IsNull(encodedSubNode) {
		return i.NormalizeBranchNode(node)
	}

	return node
}

// DelKvNode ...
func (i *Impl) DelKvNode(node Node, trieKey []uint8) Node {
	n, ok := node.([][]uint8)
	if !ok {
		log.Fatal("not ok")
	}
	currentKey := triecodec.ExtractNodeKey(n)
	nodeType := GetNodeType(node)

	if !KeyStartsWith(trieKey, currentKey) {
		return node
	} else if nodeType == NodeTypeLeaf {
		if KeyEquals(trieKey, currentKey) {
			return nil
		}

		return node
	}

	subKey := trieKey[len(currentKey):]
	subNode := i.GetNode(n[1])
	newSub := i.Del(subNode, subKey)
	encodedNewSub := i.PersistNode(newSub)

	ens, ok := encodedNewSub.([]uint8)
	if KeyEquals(ens, n[1]) {
		return node
	} else if IsNull(newSub) {
		return nil
	}

	if IsKvNode(newSub) {
		ns, ok := newSub.([][]uint8)
		if !ok {
			log.Fatal("not ok")
		}
		subNibbles := triecodec.DecodeNibbles(ns[0])
		newKey := u8util.Concat(currentKey, subNibbles)

		return NewNode([][]uint8{triecodec.EncodeNibbles(newKey), ns[1]})
	} else if IsBranchNode(newSub) {
		return NewNode([][]uint8{triecodec.EncodeNibbles(currentKey), ens})
	}

	log.Fatal("Unreachable")
	return NewNode(nil)
}

// Get ...
func (i *Impl) Get(node Node, trieKey []uint8) Node {
	if IsEmptyNode(node) {
		return nil
	} else if IsBranchNode(node) {
		nodeBranch, ok := node.(NodeBranch)
		if !ok {
			log.Fatal("not ok")
		}
		return i.GetBranchNode(nodeBranch, trieKey)
	} else if IsKvNode(node) {
		nodekv, ok := node.(NodeKv)
		if !ok {
			log.Fatal("not ok")
		}
		return i.GetKvNode(nodekv, trieKey)
	}

	log.Fatal("Invalid NodeType")
	return NewNode(nil)
}

// GetBranchNode ...
func (i *Impl) GetBranchNode(node NodeBranch, trieKey []uint8) Node {
	if len(trieKey) == 0 {
		return node[16]
	}

	n, ok := node[trieKey[0]].([]uint8)
	if !ok {
		log.Fatal("not ok")
	}
	subNode := i.GetNode(n)

	return i.Get(subNode, trieKey[1:])
}

// GetKvNode ...
func (i *Impl) GetKvNode(node NodeKv, trieKey []uint8) Node {
	currentKey := triecodec.ExtractNodeKey(node)
	nodeType := GetNodeType(node)

	if nodeType == NodeTypeLeaf {
		if KeyEquals(trieKey, currentKey) {
			return node[1]
		}

		return nil
	} else if nodeType == NodeTypeExtension {
		if KeyStartsWith(trieKey, currentKey) {
			subNode := i.GetNode(node[1])

			var start int
			if currentKey != nil {
				start = len(currentKey)
			}

			return i.Get(subNode, trieKey[start:])
		}

		return nil
	}

	log.Fatal("Unreachable")
	return NewNode(nil)
}

// NodeToDBMapping ...
func (i *Impl) NodeToDBMapping(node Node) Node {
	if IsEmptyNode(node) {
		return NewNode([][]uint8{nil, nil})
	}

	encoded := EncodeNode(node)
	if len(encoded) < 32 {
		return NewNode([][]uint8{nil, nil})
	}

	return NewNode([][]uint8{triecodec.Hashing(encoded), encoded})
}

// NormalizeBranchNode ...
func (i *Impl) NormalizeBranchNode(node Node) Node {
	n, ok := node.([][]uint8)
	if !ok {
		log.Fatal("not ok")
	}
	var indexed []struct {
		index int
		value []uint8
	}
	for i, val := range n {
		indexed[i] = struct {
			index int
			value []uint8
		}{
			index: i,
			value: val,
		}
	}

	var mapped []struct {
		index int
		value []uint8
	}

	for _, entry := range indexed {
		if entry.value != nil && len(entry.value) > 0 {
			mapped = append(mapped, entry)
		}
	}

	if len(mapped) >= 2 {
		return node
	} else if n[16] != nil {
		return NewNode([][]uint8{ComputeLeafKey([]byte{}), n[16]})
	}

	index := mapped[0].index
	value := mapped[0].value

	subNode := i.GetNode(value)

	if IsBranchNode(subNode) {
		pn := i.PersistNode(subNode)
		return [][]uint8{triecodec.EncodeNibbles([]uint8{uint8(index)}), pn.([]uint8)}
	} else if IsKvNode(subNode) {
		sn := subNode.([][]uint8)
		subNibbles := triecodec.DecodeNibbles(sn[0])
		newKey := u8util.Concat([]uint8{uint8(index)}, subNibbles)

		return [][]uint8{triecodec.EncodeNibbles(newKey), sn[1]}
	}

	log.Fatal("Unreachable")
	return NewNode(nil)
}

// PersistNode ...
func (i *Impl) PersistNode(node Node) Node {
	n1 := i.NodeToDBMapping(node)
	n, ok := n1.([][]uint8)
	if !ok {
		log.Fatal("not ok")
	}
	key := n[0]
	value := n[1]

	if value != nil {
		i.db.Put(key, value)
	}

	return key
}

// Put ...
func (i *Impl) Put(node Node, trieKey []uint8, value []uint8) Node {
	if IsEmptyNode(node) {
		return NewNode([][]uint8{ComputeLeafKey(trieKey), value})
	} else if IsKvNode(node) {
		nodekv, ok := node.(NodeKv)
		if !ok {
			log.Fatal("not ok")
		}
		return i.PutKvNode(nodekv, trieKey, value)
	} else if IsBranchNode(node) {
		return i.PutBranchNode(node, trieKey, value)
	}

	log.Fatal("Unreachable")
	return NewNode(nil)
}

// PutBranchNode ...
func (i *Impl) PutBranchNode(node Node, trieKey []uint8, value []uint8) Node {
	n, ok := node.([][]uint8)
	if !ok {
		log.Fatal("not ok")
	}
	if trieKey != nil && len(trieKey) > 0 {
		subNode := i.GetNode(n[trieKey[0]])
		newNode := i.Put(subNode, trieKey[1:], value)

		pn := i.PersistNode(newNode)
		n[trieKey[0]] = pn.([]uint8)
	} else {
		n[len(n)-1] = value
	}

	return n
}

// PutKvNode ...
func (i *Impl) PutKvNode(node NodeKv, trieKey []uint8, value []uint8) Node {
	currentKey := triecodec.ExtractNodeKey(node)

	ccp := ConsumeCommonPrefix(currentKey, trieKey)
	commonPrefix := ccp[0]
	currentRemainder := ccp[1]
	trieRemainder := ccp[2]

	isExtension := IsExtensionNode(node)
	isLeaf := IsLeafNode(node)
	var newNode [][]uint8

	if len(currentRemainder) == 0 && len(trieRemainder) == 0 {
		if isLeaf {
			return [][]uint8{node[0], value}
		}

		subNode := i.GetNode(node[1])

		nn := i.Put(subNode, trieRemainder, value)
		newNode = nn.([][]uint8)
	} else if len(currentRemainder) == 0 {
		if isExtension {
			subNode := i.GetNode(node[1])

			nn := i.Put(subNode, trieRemainder, value)
			newNode = nn.([][]uint8)
		} else {
			subPosition := trieRemainder[0]
			subKey := ComputeLeafKey(trieRemainder[1:])
			subNode := NodeKv([][]uint8{subKey, value})

			blankBranch := make([][]uint8, 16)
			newNode = append(blankBranch, node[1])
			n := i.PersistNode(subNode)
			newNode[subPosition] = n.([]uint8)
		}
	} else {
		newNode = make([][]uint8, 17)

		if len(currentRemainder) == 1 && isExtension {
			newNode[currentRemainder[0]] = node[1]
		} else {
			var computedKey EncodedPath
			if isExtension {
				computedKey = ComputeExtensionKey(currentRemainder[1:])
			} else {
				computedKey = ComputeLeafKey(currentRemainder[1:])
			}

			n := i.PersistNode([][]uint8{computedKey, node[1]})
			newNode[currentRemainder[0]] = n.([]uint8)
		}

		if len(trieRemainder) > 0 {
			n := i.PersistNode([][]uint8{ComputeLeafKey(trieRemainder[1:]), value})
			newNode[trieRemainder[0]] = n.([]uint8)
		} else {
			newNode[16] = value
		}
	}

	if len(commonPrefix) != 0 {
		n := i.PersistNode(newNode)
		return [][]uint8{ComputeExtensionKey(commonPrefix), n.([]uint8)}
	}

	return newNode
}

// SetRootNode ...
func (i *Impl) SetRootNode(node Node) {
	if IsEmptyNode(node) {
		i.checkpoint.rootHash = []byte{}
	} else {
		encoded := EncodeNode(node)
		rootHash := triecodec.Hashing(encoded)

		i.db.Put(rootHash, encoded)

		i.checkpoint.rootHash = rootHash
	}
}
