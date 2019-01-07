package triedb

import (
	"log"

	"github.com/c3systems/go-substrate/common/triecodec"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/rlp"
)

// InterfaceCodec ....
type InterfaceCodec interface {
	Encode(value interface{}) ([]uint8, error)
	Decode(encoded []byte, decoded interface{}) error
}

// RLPCodec ...
type RLPCodec struct{}

// NewRLPCodec ...
func NewRLPCodec() *RLPCodec {
	return &RLPCodec{}
}

// Encode ...
func (r *RLPCodec) Encode(value interface{}) ([]uint8, error) {
	return rlp.EncodeToBytes(&value)
}

// Decode ...
func (r *RLPCodec) Decode(encoded []byte, result interface{}) error {
	return rlp.DecodeBytes(encoded, result)
}

// TrieCodec ...
type TrieCodec struct{}

// NewTrieCodec ...
func NewTrieCodec() *TrieCodec {
	return &TrieCodec{}
}

// Encode ...
func (r *TrieCodec) Encode(value interface{}) ([]uint8, error) {
	// TODO: not working
	var input []interface{}
	switch v := value.(type) {
	case []interface{}:
		for _, x := range v {
			input = append(input, x)
		}
	case *[]interface{}:
		for _, x := range *v {
			input = append(input, x)
		}
	default:
		spew.Dump(v)
		log.Fatal("Codec: Encode; type not found")
	}
	spew.Dump(input)

	result := triecodec.Encode(input)
	return result, nil
}

// Decode ...
func (r *TrieCodec) Decode(encoded []byte, result interface{}) error {
	decoded := triecodec.Decode(encoded)
	result = &decoded
	return nil
}
