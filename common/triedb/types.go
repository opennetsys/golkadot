package triedb

import (
	"github.com/opennetsys/go-substrate/common/crypto"
	"github.com/opennetsys/go-substrate/common/db"
	"github.com/davecgh/go-spew/spew"
)

// TODO: refactor

// NodeType ...
type NodeType int

// Node ..
type Node interface{}

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
	case []*crypto.Blake2b256Hash:
		for _, x := range v {
			ret = append(ret, x)
		}
	case []*crypto.Blake2b512Hash:
		for _, x := range v {
			ret = append(ret, x)
		}
	case []*crypto.Hash:
		for _, x := range v {
			ret = append(ret, x)
		}
	case []uint8:
		if v != nil || len(v) != 0 {
			ret = append(ret, v)
		}
	case *crypto.Blake2b256Hash:
		if v != nil || len(v) != 0 {
			ret = append(ret, v)
		}
	case *crypto.Blake2b512Hash:
		if v != nil || len(v) != 0 {
			ret = append(ret, v)
		}
	case *crypto.Hash:
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
		// note: check length?
		if len(v) > 0 && v[0] != nil {
			return NewUint8FromNode(v[0])
		}
	case Node:
		switch u := v.(type) {
		case []interface{}:
			if len(u) > 0 && u[0] != nil {
				return NewUint8FromNode(u[0])
				//return u[0].([]uint8)[:]
			}
		case []uint8:
			return u
		case *crypto.Blake2b256Hash:
			return u[:]
		case *crypto.Blake2b512Hash:
			return u[:]
		case *crypto.Hash:
			return u[:]
		}
	case []uint8:
		return v
	case *crypto.Blake2b256Hash:
		return v[:]
	case *crypto.Blake2b512Hash:
		return v[:]
	case *crypto.Hash:
		return v[:]
	case []interface{}:
		if len(v) > 0 && v[0] != nil {
			return NewUint8FromNode(v[0])
			//return v[0].(*crypto.Blake2b256Hash)[:]
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
	case []*crypto.Blake2b256Hash:
		var ret [][]uint8
		for idx := range v {
			ret = append(ret, v[idx][:])
		}
		return ret
	case []*crypto.Blake2b512Hash:
		var ret [][]uint8
		for idx := range v {
			ret = append(ret, v[idx][:])
		}
		return ret
	case []*crypto.Hash:
		var ret [][]uint8
		for idx := range v {
			ret = append(ret, v[idx][:])
		}
		return ret
	case []Node:
		for _, x := range v {
			ret = append(ret, NewUint8FromNode(x))
			//ret = append(ret, x.([]uint8))
		}
	case Node:
		switch u := v.(type) {
		case []Node:
			for _, x := range u {
				ret = append(ret, NewUint8FromNode(x))
				//ret = append(ret, x.([]uint8))
			}
		case []interface{}:
			spew.Dump(u)
			for _, x := range u {
				ret = append(ret, NewUint8FromNode(x))
				//ret = append(ret, x.([]uint8))
			}
		}
	case []interface{}:
		for _, x := range v {
			ret = append(ret, NewUint8FromNode(x))
			//ret = append(ret, x.([]uint8))
		}
	}

	return ret
}

// NewFirstUint8ListFromNode ...
func NewFirstUint8ListFromNode(node Node) []uint8 {
	// note: check array length, first?
	switch v := node.(type) {
	case []Node:
		return NewUint8FromNode(v[0])
		//return v[0].([]uint8)[:]
	case [][]uint8:
		return v[0]
	case []*crypto.Blake2b256Hash:
		return v[0][:]
	case []*crypto.Blake2b512Hash:
		return v[0][:]
	case []*crypto.Hash:
		return v[0][:]
	}

	return nil
}

// InterfaceTrieDB ....
type InterfaceTrieDB interface {
	GetRoot() []uint8
	SetRoot(rootHash []uint8)
	Snapshot(dest *TrieDB, fn db.ProgressCB) int64
}
