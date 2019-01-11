package triedb

import (
	"github.com/c3systems/go-substrate/common/db"
	"github.com/davecgh/go-spew/spew"
)

// TODO: refactor

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

// NodeEncoded ...
type NodeEncoded struct{}

// NewNodeEmpty ...
func NewNodeEmpty() Node {
	return Node(nil)
}

// EncodedPath ...
type EncodedPath []uint8

// NewBlankBranch ...
func NewBlankBranch() []EncodedPath {
	return []EncodedPath{
		NewEncodedPath(nil),
		NewEncodedPath(nil),
		NewEncodedPath(nil),
		NewEncodedPath(nil),
		NewEncodedPath(nil),
		NewEncodedPath(nil),
		NewEncodedPath(nil),
		NewEncodedPath(nil),
		NewEncodedPath(nil),
		NewEncodedPath(nil),
		NewEncodedPath(nil),
		NewEncodedPath(nil),
		NewEncodedPath(nil),
		NewEncodedPath(nil),
		NewEncodedPath(nil),
		NewEncodedPath(nil),
	}
}

// NewEncodedPath ...
func NewEncodedPath(value []uint8) EncodedPath {
	return EncodedPath(value)
}

// Node ..
type Node interface{}

// NewNode ...
func NewNode(value interface{}) Node {
	return Node(value)
}

// NewNodeListFromUint8 ...
func NewNodeListFromUint8(values [][]uint8) []Node {
	var ret []Node
	for _, x := range values {
		ret = append(ret, NewNode(x))
	}
	return ret
}

// NewNodeListFromNode ...
func NewNodeListFromNode(node Node) []Node {
	var ret []Node
	switch v := node.(type) {
	case []Node:
		for _, x := range v {
			ret = append(ret, NewNode(x))
		}
	case [][]uint8:
		for _, x := range v {
			ret = append(ret, x)
		}
	case []uint8:
		if v != nil || len(v) != 0 {
			ret = append(ret, v)
		}
	case []interface{}:
		for _, x := range v {
			ret = append(ret, NewNode(x))
		}
	}

	return ret
}

// NewUint8FromNode ...
func NewUint8FromNode(value interface{}) []uint8 {
	switch v := value.(type) {
	case []Node:
		return v[0].([]uint8)
	case Node:
		switch u := v.(type) {
		case []interface{}:
			if u[0] != nil {
				return u[0].([]uint8)
			}
		}
		if u, ok := v.([]uint8); ok {
			return u
		}
	case []uint8:
		return v
	case []interface{}:
		if len(v) > 0 {
			return v[0].([]uint8)
		}
	}

	return nil
}

// NewUint8ListFromNode ...
func NewUint8ListFromNode(value interface{}) [][]uint8 {
	spew.Dump(value)
	var ret [][]uint8
	switch v := value.(type) {
	case [][]uint8:
		return v
	case []Node:
		for _, x := range v {
			ret = append(ret, x.([]uint8))
		}
	case Node:
		switch u := v.(type) {
		case []Node:
			for _, x := range u {
				ret = append(ret, x.([]uint8))
			}
		case []interface{}:
			spew.Dump(u)
			for _, x := range u {
				ret = append(ret, x.([]uint8))
			}
		}
	case []interface{}:
		for _, x := range v {
			ret = append(ret, x.([]uint8))
		}
	}

	return ret
}

// NewFirstUint8ListFromNode ...
func NewFirstUint8ListFromNode(node Node) []uint8 {
	switch v := node.(type) {
	case []Node:
		return v[0].([]uint8)
	case [][]uint8:
		return v[0]
	}

	return nil
}

// InterfaceTrieDB ....
type InterfaceTrieDB interface {
	GetRoot() []uint8
	SetRoot(rootHash []uint8)
	Snapshot(dest *Trie, fn db.ProgressCB) int64
}
