package triedb

import "github.com/c3systems/go-substrate/common/db"

// NodeType ...
type NodeType int

// NodeTypeEmpty ...
var NodeTypeEmpty = 0

// NodeTypeLeaf ...
var NodeTypeLeaf = 1

// NodeTypeExtension ...
var NodeTypeExtension = 2

// NodeTypeBranch ...
var NodeTypeBranch = 3

// NodeEmpty ...
type NodeEmpty struct{}

// null

// NodeEncoded ...
type NodeEncoded struct{}

// Uint8Array

// NodeEncodedOrEmpty ...
type NodeEncodedOrEmpty struct{}

// NodeEncoded | NodeEmpty

// NodeBranch ...
type NodeBranch []Node

/*
type NodeBranch = [
  NodeEncodedOrEmpty, NodeEncodedOrEmpty, NodeEncodedOrEmpty, NodeEncodedOrEmpty,
  NodeEncodedOrEmpty, NodeEncodedOrEmpty, NodeEncodedOrEmpty, NodeEncodedOrEmpty,
  NodeEncodedOrEmpty, NodeEncodedOrEmpty, NodeEncodedOrEmpty, NodeEncodedOrEmpty,
  NodeEncodedOrEmpty, NodeEncodedOrEmpty, NodeEncodedOrEmpty, NodeEncodedOrEmpty,
  NodeEncodedOrEmpty
];
*/

// NewNodeEmpty ...
func NewNodeEmpty() Node {
	return Node(NodeEmpty{})
}

// NewNodeBranch ...
func NewNodeBranch(nodes []Node) NodeBranch {
	return NodeBranch(nodes)
}

// EncodedPath ...
type EncodedPath []uint8

// NewEncodedPath ...
func NewEncodedPath(value []uint8) EncodedPath {
	return EncodedPath(value)
}

// Uint8Array

// NodeKv ...
type NodeKv struct{}

// [EncodedPath, NodeEncodedOrEmpty];

// NodeNotEmpty ...
type NodeNotEmpty struct{}

// NodeKv | NodeBranch;

// Node ..
type Node interface{}

// NewNode ...
func NewNode(value interface{}) Node {
	return Node(value)
}

// NodeEmpty | NodeNotEmpty;

// InterfaceTrieDB ....
type InterfaceTrieDB interface {
	GetRoot() []uint8
	SetRoot(rootHash []uint8)
	Snapshot(dest TrieDB, fn db.ProgressCB) int64
}

//type TrieDB interface TrieDb extends TxDb {
