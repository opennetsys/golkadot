package triedb

import (
	"log"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/opennetsys/godot/common/crypto"
	"github.com/opennetsys/godot/common/triecodec"
)

// InterfaceCodec ....
type InterfaceCodec interface {
	Encode(value interface{}) ([]uint8, error)
	Decode(encoded []byte, decoded interface{}) error
}

// RLPCodec ...
type RLPCodec struct {
	Name string
}

// NewRLPCodec ...
func NewRLPCodec() *RLPCodec {
	return &RLPCodec{
		Name: "Ethereum",
	}
}

// Encode ...
func (r *RLPCodec) Encode(value interface{}) ([]uint8, error) {
	var i []interface{}

	// TODO: refactor this
	slice, ok := value.([]Node)
	if ok {
		for _, s := range slice {
			if s == nil {
				// NOTE: empty string is required for nil values
				i = append(i, "")
			} else {
				slice, ok := s.([]Node)
				if ok {
					var n []Node
					for _, s := range slice {
						if s == nil {
							// NOTE: empty string is required for nil values
							n = append(n, "")
						} else {
							n = append(n, s)
						}
					}
					i = append(i, n)
				} else {
					i = append(i, s)
				}
			}
		}
	} else {
		slice, ok := value.([][][]uint8)
		if ok {
			for _, s := range slice {
				if s == nil {
					// NOTE: empty string is required for nil values
					i = append(i, "")
				} else {
					i = append(i, s)
				}
			}
		} else {
			slice, ok := value.([][]uint8)
			if ok {
				for _, s := range slice {
					if s == nil {
						// NOTE: empty string is required for nil values
						i = append(i, "")
					} else {
						i = append(i, s)
					}
				}
			} else {
				i = append(i, slice)
			}
		}
	}

	debugLog("Codec, Encode: normalized decoded arg to codec encoder", i)

	return rlp.EncodeToBytes(&i)
}

// Decode ...
func (r *RLPCodec) Decode(encoded []byte, result interface{}) error {
	return rlp.DecodeBytes(encoded, result)
}

// TrieCodec ...
type TrieCodec struct {
	Name string
}

// NewTrieCodec ...
func NewTrieCodec() *TrieCodec {
	return &TrieCodec{
		Name: "Substrate",
	}
}

func enc(input interface{}) interface{} {
	var output []interface{}

	switch v := input.(type) {
	case nil:
		return nil
	case Node:
		switch u := v.(type) {
		case []interface{}:
			for _, x := range u {
				result := enc(x)
				output = append(output, result)
			}
		case EncodedPath:
			var a []uint8
			for _, x := range u {
				a = append(a, x)
			}
			return a
		case []uint8:
			return u
		case *crypto.Blake2b256Hash:
			return u[:]
		case *crypto.Blake2b512Hash:
			return u[:]
		case *crypto.Hash:
			return u[:]
		case []Node:
			var a []interface{}
			for _, x := range u {
				a = append(a, enc(x))
			}
			return a
		default:
			spew.Dump(u)
			log.Fatal("enc error; not found")
		}
	case []Node:
		for _, x := range v {
			result := enc(x)
			output = append(output, result)
		}
	case [][]uint8:
		for _, x := range v {
			output = append(output, x)
		}
	case []*crypto.Blake2b256Hash:
		for _, x := range v {
			output = append(output, x[:])
		}
	case []*crypto.Blake2b512Hash:
		for _, x := range v {
			output = append(output, x[:])
		}
	case []*crypto.Hash:
		for _, x := range v {
			output = append(output, x[:])
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
		for _, x := range v {
			result := enc(x)
			output = append(output, result)
		}
	case EncodedPath:
		var a []uint8
		for _, x := range v {
			a = append(a, x)
		}
		return a
	case *[]interface{}:
		for _, x := range *v {
			result := enc(x)
			output = append(output, result)
		}
	default:
		log.Fatal("HO", v)
	}

	return output
}

// Encode ...
func (r *TrieCodec) Encode(value interface{}) ([]uint8, error) {
	debugLog("triedb triecodec, Encode raw input", value)

	var input []interface{}

	switch v := value.(type) {
	case []Node:
		for _, x := range v {
			result := enc(x)
			input = append(input, result)
		}
	case [][]uint8:
		for _, x := range v {
			input = append(input, x)
		}
	case []*crypto.Blake2b256Hash:
		for _, x := range v {
			input = append(input, x)
		}
	case []*crypto.Blake2b512Hash:
		for _, x := range v {
			input = append(input, x)
		}
	case []*crypto.Hash:
		for _, x := range v {
			input = append(input, x)
		}
	case []interface{}:
		for _, x := range v {
			result := enc(x)
			input = append(input, result)
		}
	case *[]interface{}:
		for _, x := range *v {
			result := enc(x)
			input = append(input, result)
		}
	}

	debugLog("triedb triecodec, Encode, normalized input", input)

	result := triecodec.Encode(input)
	return result, nil
}

// Decode ...
func (r *TrieCodec) Decode(encoded []byte, result interface{}) error {
	decoded := triecodec.Decode(encoded)

	switch v := decoded.(type) {
	case []uint8:
		var a []interface{}
		for _, x := range v {
			a = append(a, x)
		}
		reflect.ValueOf(result).Elem().Set(reflect.ValueOf(a))
	case *crypto.Blake2b256Hash:
		var a []interface{}
		for _, x := range v {
			a = append(a, x)
		}
		reflect.ValueOf(result).Elem().Set(reflect.ValueOf(a))
	case *crypto.Blake2b512Hash:
		var a []interface{}
		for _, x := range v {
			a = append(a, x)
		}
		reflect.ValueOf(result).Elem().Set(reflect.ValueOf(a))
	case *crypto.Hash:
		var a []interface{}
		for _, x := range v {
			a = append(a, x)
		}
		reflect.ValueOf(result).Elem().Set(reflect.ValueOf(a))
	default:
		reflect.ValueOf(result).Elem().Set(reflect.ValueOf(v))
	}

	return nil
}
