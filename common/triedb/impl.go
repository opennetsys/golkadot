package triedb

import "github.com/c3systems/go-substrate/db"

// Impl ...
type Impl struct {
	checkpoint *Checkpoint
	db         *TxDB
}

// TxDB ...
type TxDB interface {
}

// NewImpl ...
func NewImpl(db *TxDB, rootHash []uint8) *Impl {
	checkpoint := NewCheckpoint(rootHash)

	return &Impl{
		checkpoint: checkpoint,
	}
}

// Snapshot ...
func (i *Impl) Snapshot(dest TrieDB, fn db.ProgressCB, root []uint8, keys int, percent int, depth int) int {
	node := i.GetNode(root)

	if node == nil {
		return keys
	}

	keys++

	// TODO
	/*
	   dest.db.Put(root, encodeNode(node))
	   fn({ keys, percent })

	   node.forEach((u8a) => {
	     if (u8a != nil && len(u8a) == 32) {
	       keys = i._snapshot(dest, fn, u8a, keys, percent, depth + 1)
	     }

	     percent += (100 / len(node)) / mathutil.Pow(16, depth)
	   })

	   return keys
	*/
	return 0
}

// GetNode ...
func (i *Impl) GetNode(hash []uint8) *Node {
	// TODO
	/*
	    if hash == nil || len(hash) == 0 || keyEquals(hash, i.constants.EMPTY_HASH) {
	      return nil
	    } else if len(hash) < 32 {
	      return decodeNode(hash)
	    }

	    return decodeNode(i.db.Get(hash))
	  }
	*/
	return nil
}

// Del ...
func (i *Impl) Del(node *Node, trieKey []uint8) *Node {
	// TODO
	/*
	   if isEmptyNode(node) {
	     return nil
	   } else if isBranchNode(node) {
	     return i._delBranchNode(node, trieKey)
	   } else if isKvNode(node) {
	     return i._delKvNode(node, trieKey)
	   }

	   log.Fatal('Unreachable')
	*/
	return nil
}

// NodeBranch ...
type NodeBranch struct{}

// DelBranchNode ...
func (i *Impl) DelBranchNode(node *NodeBranch, trieKey []uint8) *Node {
	// TODO
	/*
		   if len(trieKey) == 0 {
		     node[len(node) - 1] = nil

		     return i._normaliseBranchNode(node)
		   }

			 nodeToDelete := i._getNode(node[trieKey[0]])
			 subNode := i._del(nodeToDelete, trieKey.subarray(1))
			 encodedSubNode := i._persistNode(subNode)

		   if keyEquals(encodedSubNode, node[trieKey[0]]) {
		     return node
		   }

		   node[trieKey[0]] = encodedSubNode

		   if isNull(encodedSubNode) {
		     return i._normaliseBranchNode(node)
			 }

		  return node
	*/
	return nil
}

// NodeNotEmpty ...
type NodeNotEmpty struct{}

// DelKvNode ...
func (i *Impl) DelKvNode(node *NodeNotEmpty, trieKey []uint8) *Node {
	// TODO
	/*
		   const currentKey = extractNodeKey(node)
		   const nodeType = getNodeType(node)

		   if !keyStartsWith(trieKey, currentKey) {
		     return node
		   } else if nodeType == NodeType.LEAF {
		     if keyEquals(trieKey, currentKey) {
					 return nil
				 } else {
					 reutrn node
				 }
		   }

			 subKey := trieKey.subarray(len(currentKey))
			 subNode := i._getNode(node[1])
			 newSub := i._del(subNode, subKey)
			 encodedNewSub := i._persistNode(newSub)

		   if keyEquals(encodedNewSub, node[1]) {
		     return node
		   } else if isNull(newSub) {
		     return nil
		   }

		   if isKvNode(newSub) {
		     const subNibbles = decodeNibbles(newSub[0])
		     const newKey = u8aConcat(currentKey, subNibbles)

		     return [encodeNibbles(newKey), newSub[1]]
		   } else if (isBranchNode(newSub)) {
		     return [encodeNibbles(currentKey), encodedNewSub]
		   }

		   log.Fatal('Unreachable')
	*/
	return nil
}

// NodeEncodedOrEmpty ...
type NodeEncodedOrEmpty struct{}

// Get ...
func (i *Impl) Get(node *Node, trieKey []uint8) *NodeEncodedOrEmpty {
	// TODO
	/*
	   if isEmptyNode(node) {
	     return nil
	   } else if isBranchNode(node) {
	     return i._getBranchNode(node, trieKey)
	   } else if isKvNode(node) {
	     return i._getKvNode(node, trieKey)
	   }

	   log.Fatal('Invalid NodeType')
	*/
	return nil
}

// GetBranchNode ...
func (i *Impl) GetBranchNode(node *NodeBranch, trieKey []uint8) *NodeEncodedOrEmpty {
	// TODO
	/*
		   if len(trieKey) == 0 {
		     return node[16]
		   }

			 subNode := i._getNode(node[trieKey[0]])

		   return i._get(subNode, trieKey.subarray(1))
	*/
	return nil
}

// NodeKv ...
type NodeKv struct{}

// GetKvNode ...
func (i *Impl) GetKvNode(node *NodeKv, trieKey []uint8) *NodeEncodedOrEmpty {
	// TODO
	/*
			currentKey := extractNodeKey(node)
			nodeType := getNodeType(node)

		   if nodeType == NodeType.LEAF {
		     if keyEquals(trieKey, currentKey) {
					 return node[1]
				 } else {
		       return null
				 }
		   } else if nodeType == NodeType.EXTENSION {
		     if keyStartsWith(trieKey, currentKey) {
					 subNode := i._getNode(node[1])

					 var start int
					 if currentKey {
						 start = len(currentKey)
					}

		       return i._get(subNode, trieKey.subarray(start))
		     }

		     return null
		   }

		   log.Fatal('Unreachable')
	*/
	return nil
}

// NodeToDBMapping ...
func (i *Impl) NodeToDBMapping(node *Node) *NodeNotEmpty {
	// TODO
	/*
		   if isEmptyNode(node) {
		     return [null, null]
		   }

			 encoded := encodeNode(node)

		   if len(encoded) < 32 {
		     return [node as any, null]
			 } else {
		     return [i.codec.hashing(encoded), encoded]
			 }
	*/
	return nil
}

// NormalizeBranchNode ...
func (i *Impl) NormalizeBranchNode(node *NodeNotEmpty) *Node {
	// TODO
	/*
		mapped := node
		     .map((value, index) => ({ index, value }))
		     .filter(({ value }) => !!value && len(value) !== 0)

		   if len(mapped) >= 2 {
		     return node
		   } else if node[16] != nil {
		     return [computeLeafKey(new Uint8Array()), node[16]]
		   }

		   const [{ index, value }] = mapped
			 subNode := i._getNode(value)

		   if isBranchNode(subNode) {
		     return [encodeNibbles(new Uint8Array([index])), i._persistNode(subNode)]
		   } else if isKvNode(subNode) {
				 subNibbles := decodeNibbles(subNode[0])
				 newKey := u8aConcat(new Uint8Array([index]), subNibbles)

		     return [encodeNibbles(newKey), subNode[1]]
		   }

		   throw new Error('Unreachable')
	*/
	return nil
}

// PersistNode ...
func (i *Impl) PersistNode(node *Node) *NodeEncodedOrEmpty {
	// TODO
	/*
	   const [key, value] = i._nodeToDbMapping(node)

	   if value != nil {
	     i.db.Put(key as Uint8Array, value)
	   }

	   return key
	*/
	return nil
}

// Put ...
func (i *Impl) Put(node *Node, trieKey []uint8, value []uint8) *NodeNotEmpty {
	// TODO
	/*

	   if isEmptyNode(node) {
	     return [computeLeafKey(trieKey), value]
	   } else if isKvNode(node) {
	     return i._putKvNode(node, trieKey, value)
	   } else if isBranchNode(node) {
	     return i._putBranchNode(node, trieKey, value)
	   }

	   log.Fatal('Unreachable')
	*/
	return nil
}

// PutBranchNode ...
func (i *Impl) PutBranchNode(node *Node, trieKey []uint8, value []uint8) *NodeNotEmpty {
	// TODO
	/*
		   if (trieKey && len(trieKey)) {
				 subNode := i._getNode(node[trieKey[0]])
				 newNode := i._put(subNode, trieKey.subarray(1), value)

		     node[trieKey[0]] = i._persistNode(newNode)
		   } else {
		     node[len(node) - 1] = value
		   }

		   return node
	*/
	return nil
}

// PutKvNode ...
func (i *Impl) PutKvNode(node *NodeKv, trieKey []uint8, value []uint8) *NodeNotEmpty {
	// TODO
	/*
			currentKey := extractNodeKey(node)
		   var [commonPrefix, currentRemainder, trieRemainder] = consumeCommonPrefix(currentKey, trieKey)
			 isExtension := isExtensionNode(node)
			 isLeaf := isLeafNode(node)
		   var newNode NodeNotEmpty

		   if len(currentRemainder) == 0 && len(trieRemainder) == 0 {
		     if isLeaf {
		       return [node[0], value]
		     }

				 subNode := i._getNode(node[1])

		     newNode = i._put(subNode, trieRemainder, value)
		   } else if len(currentRemainder) == 0 {
		     if isExtension {
					 subNode := i._getNode(node[1])

		       newNode = i._put(subNode, trieRemainder, value)
		     } else {
					 subPosition := trieRemainder[0]
					 subKey := computeLeafKey(trieRemainder.subarray(1))
		       var subNode: NodeKv = [subKey, value]

		       newNode = BLANK_BRANCH.concat(node[1]) as NodeNotEmpty
		       newNode[subPosition] = i._persistNode(subNode)
		     }
		   } else {
		     newNode = BLANK_BRANCH.concat(null) as NodeNotEmpty

		     if len(currentRemainder) == 1 && isExtension {
		       newNode[currentRemainder[0]] = node[1]
		     } else {
					 var computedKey
					 if isExtension {
						 computedKey = computeExtensionKey(currentRemainder.subarray(1))
					 } else {
		         computedKey = computeLeafKey(currentRemainder.subarray(1))
					 }

		       newNode[currentRemainder[0]] = i._persistNode([computedKey, node[1]])
		     }

		     if len(trieRemainder) {
		       newNode[trieRemainder[0]] = i._persistNode([computeLeafKey(trieRemainder.subarray(1)), value])
		     } else {
		       newNode[16] = value
		     }
		   }

		   if len(commonPrefix) != 0 {
		     return [computeExtensionKey(commonPrefix), i._persistNode(newNode)]
		   }

		   return newNode
	*/
	return nil
}

// SetRootNode ...
func (i *Impl) SetRootNode(node *Node) {
	// TODO
	/*
		   if isEmptyNode(node) {
		     i.rootHash = i.constants.EMPTY_HASH
		   } else {
				 encoded := encodeNode(node)
				 rootHash := codec.hashing(encoded)

		     i.db.Put(rootHash, encoded)

		     i.rootHash = rootHash
		   }
	*/
}
