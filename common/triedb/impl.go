package triedb

import (
	"fmt"
	"log"
	"math"

	"github.com/opennetsys/golkadot/common/db"
	"github.com/opennetsys/golkadot/common/triecodec"
	"github.com/opennetsys/golkadot/common/u8util"
)

// TODO: refactor

// Impl ...
type Impl struct {
	checkpoint *Checkpoint
	db         db.TXDB
	codec      InterfaceCodec
	Debug      bool
}

// TxDB ...
type TxDB struct {
	db.MemoryDB
}

// NewImpl ...
func NewImpl(db db.TXDB, rootHash []byte, codec InterfaceCodec) *Impl {
	checkpoint := NewCheckpoint(rootHash)
	return &Impl{
		checkpoint: checkpoint,
		db:         db,
		codec:      codec,
		Debug:      false,
	}
}

// SetDebug ...
func (i *Impl) SetDebug(enabled bool) {
	i.Debug = enabled
}

// DebugLog ...
func (i *Impl) DebugLog(ifcs ...interface{}) {
	if i.Debug {
		ifcs = append([]interface{}{"Debug: impl"}, ifcs...)
		fmt.Println(ifcs...)
	}
}

// Snapshot ...
func (i *Impl) Snapshot(dest *TrieDB, fn db.ProgressCB, root []byte, keys int, percent int, depth int) int {
	i.DebugLog("Snapshot, root", root)
	node := i.GetNode(root)
	i.DebugLog("Snapshot, GetNode result", node)

	if IsNull(node) {
		i.DebugLog("Snapshot, node is null", node)
		return keys
	}

	keys++

	i.DebugLog("Snapshot, call Put with root", root)
	encodedNode := EncodeNode(node, i.codec)
	i.DebugLog("Snapshot, call Put with encoded node", encodedNode)
	dest.impl.db.Put(root[:], encodedNode)

	if fn != nil {
		fn(&db.ProgressValue{
			IsCompleted: false,
			Keys:        keys,
			Percent:     percent,
		})
	}

	nodes := NewNodeListFromNode(node)
	if len(nodes) == 0 {
		log.Fatal("Snapshot: not ok")
	}

	for _, val := range nodes {
		v := NewUint8FromNode(val)
		if v != nil && len(v) == 32 {
			keys = i.Snapshot(dest, fn, v, keys, percent, depth+1)
		}

		percent += int((float64(100) / float64(len(nodes))) / math.Pow(float64(16), float64(depth)))
	}

	return keys
}

// GetNode ...
// NOTE: should usually be single dimension array
func (i *Impl) GetNode(nhash Node) Node {
	i.DebugLog("GetNode, hash", nhash)
	hash := NewNodeListFromNode(nhash)[:]
	var l int
	var shash []uint8
	if len(hash) == 1 {
		shash = NewUint8FromNode(hash[0])
		l = len(shash)
	} else {
		shash = NewUint8FromNode(nhash)
		l = len(hash)
	}
	i.DebugLog("GetNode, hash single", shash)
	i.DebugLog("GetNode, lengths", l, len(hash), len(shash))
	i.DebugLog("GetNode, hashes", hash, shash)
	var isEmpty bool
	if len(hash) == 0 {
		isEmpty = KeyEquals(shash, []uint8{})
		i.DebugLog("GetNode, isEmpty", isEmpty)
	}
	if shash == nil || l == 0 || isEmpty {
		i.DebugLog("GetNode, get node empty")
		return nil
	} else if l < 32 { // it's encoded key if len 32
		// is encodeed bel if less than 32?
		i.DebugLog("GetNode, less than 32")
		x := DecodeNode(nhash, i.codec)
		i.DebugLog("GetNode, less than 32, decoded", x)
		return x
	}

	i.DebugLog("GetNode, get hash", shash)
	x := i.db.Get(shash)
	i.DebugLog("GetNode, get hash result", x)
	y := DecodeNode(x, i.codec)
	i.DebugLog("GetNode, decode node result", y)
	return y
}

// Del ...
func (i *Impl) Del(node Node, trieKey []uint8) Node {
	i.DebugLog("Del, node", node)
	i.DebugLog("Del, trikey", trieKey)
	if IsEmptyNode(node) {
		i.DebugLog("Del, is empty node")
		return nil
	} else if IsBranchNode(node) {
		i.DebugLog("Del, is branch node", node)
		return i.DelBranchNode(node, trieKey)
	} else if IsKvNode(node) {
		i.DebugLog("Del, is kv node")
		i.DebugLog("Del, call DelKNode")
		return i.DelKvNode(node, trieKey)
	}

	log.Fatal("Del: Unreachable")
	return NewNode(nil)
}

// DelBranchNode ...
// NOTE: node must be NodeBranch
// NOTE: node branch must have 17 items, either single dimension
// array or nil
func (i *Impl) DelBranchNode(inode Node, trieKey []uint8) Node {
	node := NewNodeListFromNode(inode)
	i.DebugLog("DelBranchNode, node", node)
	i.DebugLog("DelBranchNode, triekey", trieKey)
	if len(trieKey) == 0 {
		node[len(node)-1] = nil

		i.DebugLog("DelBranchNode, normalizebranch")
		return i.NormalizeBranchNode(node)

	}

	i.DebugLog("DelBranchNode, node to delete initial", node[trieKey[0]])

	var nodeToDelete Node
	// TODO: refactor
	if n := NewUint8ListFromNode(node[trieKey[0]]); len(n) == 2 {
		nodeToDelete = DecodeNode(n, i.codec)
	} else {
		nodeToDelete = i.GetNode(NewUint8FromNode(node[trieKey[0]]))
	}

	i.DebugLog("DelBranchNode, node to delete", nodeToDelete)

	subNode := i.Del(nodeToDelete, trieKey[1:])
	encodedSubNode := i.PersistNode(subNode)
	i.DebugLog("DelBranchNode, delete encoded sub node", encodedSubNode)
	if encodedSubNode == nil && node[trieKey[0]] == nil {
		return node
	}

	if encodedSubNode != nil {
		if KeyEquals(NewUint8FromNode(encodedSubNode), node[trieKey[0]].([]uint8)) {
			return node
		}

		node[trieKey[0]] = NewUint8FromNode(encodedSubNode)
	} else {
		node[trieKey[0]] = nil
	}

	if IsNull(encodedSubNode) {
		return i.NormalizeBranchNode(node)
	}

	return node
}

// DelKvNode ...
// NOTE: node should be two item kv slice
func (i *Impl) DelKvNode(node Node, trieKey []uint8) Node {
	nodekv := NewUint8ListFromNode(node)
	currentKey := triecodec.ExtractNodeKey(nodekv)
	nodeType := GetNodeType(node)

	i.DebugLog("DelKvNode, node", node)
	i.DebugLog("DelKvNode, triekey", trieKey)

	if !KeyStartsWith(trieKey, currentKey) {
		i.DebugLog("DelKvNode, no starts with")
		return node
	} else if nodeType == NodeTypeLeaf {
		i.DebugLog("DelKvNode, is node type leaf")
		if KeyEquals(trieKey, currentKey) {
			i.DebugLog("DelKvNode, key equals true")
			return nil
		}

		i.DebugLog("DelKvNode, key equals false")
		return node
	}

	subKey := trieKey[len(currentKey):]
	subNode := i.GetNode(nodekv[1])
	newSub := i.Del(subNode, subKey)
	encodedNewSub := i.PersistNode(newSub)

	i.DebugLog("DelKvNode, encoded new sub", encodedNewSub)

	ens := NewUint8FromNode(encodedNewSub)
	if len(ens) != 0 && KeyEquals(ens, nodekv[1]) {
		return node
	} else if IsNull(newSub) {
		return nil
	}

	if IsKvNode(newSub) {
		ns := NewUint8ListFromNode(newSub)
		if len(ns) == 0 {
			log.Fatal("DelKvNode: not ok")
		}

		subNibbles := triecodec.DecodeNibbles(ns[0])
		newKey := u8util.Concat(currentKey, subNibbles)

		return NewNode([]Node{triecodec.EncodeNibbles(newKey), ns[1]})
	} else if IsBranchNode(newSub) {
		return NewNode([]Node{triecodec.EncodeNibbles(currentKey), encodedNewSub})
	}

	log.Fatal("DelKvNode: Unreachable")
	return NewNode(nil)
}

// Get ...
func (i *Impl) Get(node Node, trieKey []uint8) Node {
	i.DebugLog("Get, node", node)
	i.DebugLog("Get, triekey", trieKey)
	if IsEmptyNode(node) {
		i.DebugLog("Get, is empty node")
		return nil
	} else if IsBranchNode(node) {
		i.DebugLog("Get, is branch node")
		return i.GetBranchNode(NewNodeListFromNode(node), trieKey)
	} else if IsKvNode(node) {
		i.DebugLog("Get, is kv node")
		i.DebugLog("Get, kv node", node)
		return i.GetKvNode(NewNodeListFromNode(node), trieKey)
	}

	log.Fatal("Get: Invalid NodeType")
	return NewNode(nil)
}

// GetBranchNode ...
func (i *Impl) GetBranchNode(node []Node, trieKey []uint8) Node {
	i.DebugLog("GetBranchNode, node", node)
	i.DebugLog("GetBranchNode, trieKey", trieKey)
	if len(trieKey) == 0 {
		i.DebugLog("GetBranchNode, triekey len is 0, returning", node[16])
		return node[16]
	}

	var subNode Node
	s := NewUint8ListFromNode(node[trieKey[0]])
	if len(s) >= 2 {
		subNode = s
	} else {
		i.DebugLog("GetBranchNode, call GetNode with trie key", trieKey[0])
		subNode = i.GetNode(NewUint8FromNode(node[trieKey[0]]))
	}
	i.DebugLog("GetBranchNode, call Get with sub node", subNode)
	i.DebugLog("GetBranchNode, call Get with triekey", trieKey[1:])

	return i.Get(subNode, trieKey[1:])
}

// GetKvNode ...
// NOTE: kv nodes only hold two items
func (i *Impl) GetKvNode(node []Node, trieKey []uint8) Node {
	i.DebugLog("GetKvNode, node", node)
	i.DebugLog("GetKvNode, trieKey", trieKey)

	currentKey := triecodec.ExtractNodeKey(NewUint8ListFromNode(node[:1]))
	nodeType := GetNodeType(node)

	if nodeType == NodeTypeLeaf {
		i.DebugLog("GetKvNode, nodetype is leaf")
		if KeyEquals(trieKey, currentKey) {
			i.DebugLog("GetKvNode, nodetype leaf key equals")
			return node[1]
		}

		i.DebugLog("GetKvNode, returning nil")
		return nil
	} else if nodeType == NodeTypeExtension {
		i.DebugLog("GetKvNode, node type is extension")
		if KeyStartsWith(trieKey, currentKey) {
			i.DebugLog("GetKvNode, key starts with true")

			// TODO: refactor because it's not all single dimension
			var subNode Node
			if s := NewUint8FromNode(node[1]); len(s) > 0 {
				subNode = i.GetNode(s)
			} else {
				subNode = node[1]
			}

			i.DebugLog("GetKvNode, sub node", subNode)

			var start int
			if currentKey != nil {
				start = len(currentKey)
			}

			i.DebugLog("GetKvNode, trie key with start", trieKey[start:])

			return i.Get(subNode, trieKey[start:])
		}

		return nil
	}

	log.Fatal("GetKvNode: Unreachable")
	return NewNode(nil)
}

// NodeToDBMapping ...
// NOTE: return can be single dimension or nil or same node back with no value
func (i *Impl) NodeToDBMapping(node Node) (interface{}, []uint8) {
	i.DebugLog("NodeToDBMapping, node", node)
	if IsEmptyNode(node) {
		i.DebugLog("NodeToDBMapping, is empty")
		return nil, nil
	}

	encoded := EncodeNode(node, i.codec)
	if len(encoded) < 32 {
		i.DebugLog("NodeToDBMapping, is less than 32, returning")
		return node, nil
	}

	return triecodec.Hashing(encoded), encoded
}

// NormalizeBranchNode ...
// NOTE: node can be either NodeKv | NodeBranch
func (i *Impl) NormalizeBranchNode(node []Node) Node {
	i.DebugLog("NormalizeBranchNode, node input", node)
	n := node
	indexed := make([]struct {
		index int
		value Node
	}, len(node))
	for i, val := range n {
		indexed[i] = struct {
			index int
			value Node
		}{
			index: i,
			value: val,
		}
	}

	var mapped []struct {
		index int
		value Node
	}

	for _, entry := range indexed {
		i.DebugLog("NormalizeBranchNode, map filter, value", entry.value)
		if n := NewUint8FromNode(entry.value); len(n) > 0 {
			if entry.value != nil {
				mapped = append(mapped, entry)
			}
		} else if n := NewUint8ListFromNode(entry.value); len(n) > 0 {
			if entry.value != nil {
				mapped = append(mapped, entry)
			}
		}
	}

	i.DebugLog("NormalizeBranchNode, mapped", mapped)

	if len(mapped) >= 2 {
		i.DebugLog("NormalizeBranchNode, mapped length larger than 2")
		return node
	} else if n[16] != nil && len(n[16].([]uint8)) > 0 {
		i.DebugLog("NormalizeBranchNode, mapped[16] is not nil", n[16])
		return NewNode([][]uint8{ComputeLeafKey([]byte{}), n[16].([]uint8)})
	}

	index := mapped[0].index
	value := mapped[0].value

	// TODO: refactor
	var subNode Node
	if n := NewUint8ListFromNode(value); len(n) == 2 {
		i.DebugLog("NormalizeBranchNode, setting sub node to", n)
		subNode = n
	} else {
		i.DebugLog("NormalizeBranchNode, call GetNode with value", value)
		subNode = i.GetNode(NewUint8FromNode(value))
	}

	i.DebugLog("NormalizeBranchNode, sub node", subNode)

	if IsBranchNode(subNode) {
		i.DebugLog("NormalizeBranchNode, is branch node")
		pn := i.PersistNode(subNode)
		i.DebugLog("NormalizeBranchNode, is branch PersistNode result", pn)
		nibs := triecodec.EncodeNibbles([]uint8{uint8(index)})
		i.DebugLog("NormalizeBranchNode, is branch encoded nibbles", pn)
		return NewNode([]Node{
			nibs,
			pn,
		})
	} else if IsKvNode(subNode) {
		i.DebugLog("NormalizeBranchNode, is kv node")
		subNibbles := triecodec.DecodeNibbles(NewUint8ListFromNode(subNode)[0])
		newKey := u8util.Concat([]uint8{uint8(index)}, subNibbles)
		i.DebugLog("NormalizeBranchNode, is kv node new key", newKey)

		return [][]uint8{triecodec.EncodeNibbles(newKey), NewUint8ListFromNode(subNode)[1]}
	}

	log.Fatal("NormalizeBranchNode: Unreachable")
	return NewNode(nil)
}

// PersistNode ...
// NOTE: Node should nil or single dimension array
func (i *Impl) PersistNode(node Node) Node {
	i.DebugLog("PersistNode, node", node)
	ikey, value := i.NodeToDBMapping(node)

	i.DebugLog("PersistNode, key", ikey)
	i.DebugLog("PersistNode, value", value)

	if value != nil {
		k := NewUint8FromNode(ikey)
		i.db.Put(k, value)
	}

	key := NewNode(ikey)
	return key
}

// Put ...
func (i *Impl) Put(node Node, trieKey []uint8, value []uint8) Node {
	i.DebugLog("Put, node", node)
	i.DebugLog("Put, triekey", trieKey)
	i.DebugLog("Put, value", value)
	if IsEmptyNode(node) {
		i.DebugLog("Put, is empty node")
		return NewNode([]Node{ComputeLeafKey(trieKey), value})
	} else if IsKvNode(node) {
		i.DebugLog("Put, is kv node", node)
		x := i.PutKvNode(NewNodeListFromNode(node), trieKey, value)
		i.DebugLog("Put, PutKvNode result", x)
		return x
	} else if IsBranchNode(node) {
		i.DebugLog("Put, is branch node")
		return i.PutBranchNode(node, trieKey, value)
	}

	log.Fatal("Put: Unreachable")
	return NewNode(nil)
}

// PutBranchNode ...
func (i *Impl) PutBranchNode(node Node, trieKey []uint8, value []uint8) Node {
	i.DebugLog("PutBranchNode, node", node)
	i.DebugLog("PutBranchNode, triekey", trieKey)
	i.DebugLog("PutBranchNode, value", value)
	n := NewNodeListFromNode(node)
	if len(n) == 0 {
		log.Fatal("PutBranchNode: not ok")
	}
	if trieKey != nil && len(trieKey) > 0 {
		v := n[trieKey[0]]
		i.DebugLog("PutBranchNode, trie key len is gt zero")
		i.DebugLog("PutBranchNode, call GetNode with value", v)
		subNode := i.GetNode(v)
		i.DebugLog("PutBranchNode, call Put with trie key", trieKey[1:])
		newNode := i.Put(subNode, trieKey[1:], value)

		pn := i.PersistNode(newNode)
		i.DebugLog("PutBranchNode, PersistNode resul", pn)
		n[trieKey[0]] = pn
	} else {
		n[len(n)-1] = value
	}

	return n
}

// PutKvNode ...
// NOTE: KV node should be two item list
func (i *Impl) PutKvNode(node []Node, trieKey []uint8, value []uint8) Node {
	i.DebugLog("PutKvNode, input node", node)
	currentKey := triecodec.ExtractNodeKey(NewUint8ListFromNode(node[:1]))

	ccp := ConsumeCommonPrefix(currentKey, trieKey)
	commonPrefix := ccp[0]
	currentRemainder := ccp[1]
	trieRemainder := ccp[2]

	isExtension := IsExtensionNode(node)
	isLeaf := IsLeafNode(node)
	var newNode []Node
	i.DebugLog("PutKvNode, is extension?", isExtension)
	i.DebugLog("PutKvNode, is leaf?", isLeaf)

	if len(currentRemainder) == 0 && len(trieRemainder) == 0 {
		i.DebugLog("PutBranchNode, remainders is zero")
		if isLeaf {
			i.DebugLog("PutKvNode, is leaf value", value)
			return []Node{node[0], value}
		}

		// FIX
		subNode := i.GetNode(NewUint8FromNode(node[1]))

		nn := i.Put(subNode, trieRemainder, value)
		nodes := NewUint8ListFromNode(nn)
		newNode = []Node{}
		for _, n := range nodes {
			newNode = append(newNode, NewNode(n))
		}
	} else if len(currentRemainder) == 0 {
		i.DebugLog("PutKvNode, no remainders")
		if isExtension {
			i.DebugLog("PutKvNode, is extension node", node)
			i.DebugLog("PutKvNode, call GetNode with node", node[1])
			subNode := i.GetNode(node[1])

			nn := i.Put(subNode, trieRemainder, value)
			nodes := NewNodeListFromNode(nn)
			newNode = []Node{}
			for _, n := range nodes {
				newNode = append(newNode, NewNode(n))
			}
		} else {
			i.DebugLog("PutKvNode, is not extension value", value)
			subPosition := trieRemainder[0]
			subKey := ComputeLeafKey(trieRemainder[1:])
			subNode := []Node{subKey, value}

			blankBranch := make([]Node, 16)
			newNode = append(blankBranch, node[1])

			i.DebugLog("PutKvNode, call PersistNode with sub node", subNode)
			n := i.PersistNode(subNode)
			newNode[subPosition] = n
		}
	} else {
		i.DebugLog("PutKvNode, else block")
		newNode = make([]Node, 17)

		i.DebugLog("PutKvNode, new node", newNode)
		if len(currentRemainder) == 1 && isExtension {
			i.DebugLog("PutKvNode, remainder is one and is extension")
			newNode[currentRemainder[0]] = node[1]
		} else {
			i.DebugLog("PutKvNode, remainder is not one and is not extension")
			var computedKey EncodedPath
			if isExtension {
				i.DebugLog("PutKvNode, else, is extension")
				computedKey = ComputeExtensionKey(currentRemainder[1:])
			} else {
				i.DebugLog("PutKvNode, else, is not extension")
				computedKey = ComputeLeafKey(currentRemainder[1:])
			}

			i.DebugLog("PutKvNode, computed key", computedKey)
			i.DebugLog("PutKvNode, node[1]", node[1])
			n := i.PersistNode([]Node{computedKey, node[1].([]uint8)})
			i.DebugLog("PutKvNode, PersistNode result", n)

			newNode[currentRemainder[0]] = n
		}

		if len(trieRemainder) > 0 {
			i.DebugLog("PutKvNode, trie remainder is gt zero")
			i.DebugLog("PutKvNode, trie remainder is gt zero new node", newNode)
			n := i.PersistNode([]Node{ComputeLeafKey(trieRemainder[1:]), value})
			newNode[trieRemainder[0]] = n
		} else {
			i.DebugLog("PutKvNode, trie remainder is not gt zero")
			i.DebugLog("PutKvNode, trie remainder is not gt zero value", value)
			newNode[16] = value
		}
	}

	if len(commonPrefix) != 0 {
		i.DebugLog("PutKvNode, has common prefix")
		i.DebugLog("PutKvNode, has common prefix new node", newNode)
		n := i.PersistNode(newNode)
		i.DebugLog("PutKvNode, has commonprefix PersistNode result", n)
		return []Node{ComputeExtensionKey(commonPrefix), n}
	}

	i.DebugLog("PutKvNode, final new node", newNode)
	return newNode
}

// SetRootNode ...
func (i *Impl) SetRootNode(node Node) {
	i.DebugLog("SetRootNode, node", node)

	if IsEmptyNode(node) {
		i.DebugLog("SetRootNode, is empty")
		i.checkpoint.rootHash = []byte{}
	} else {
		i.DebugLog("SetRootNode, call EncodeNode")
		encoded := EncodeNode(node, i.codec)
		i.DebugLog("SetRootNode, encoded", encoded)
		rootHash := triecodec.Hashing(encoded)
		i.DebugLog("SetRootNode, root hash", rootHash)

		i.DebugLog("SetRootNode, call Put")
		i.db.Put(rootHash[:], encoded)
		i.DebugLog("SetRootNode, call set root hash", rootHash)

		i.checkpoint.rootHash = rootHash
	}
}
