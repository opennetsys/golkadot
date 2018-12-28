package fileflatdb

import (
	"log"

	"github.com/c3systems/go-substrate/common/triecodec"
	"github.com/golang/snappy"
)

// NibbleBuffer ...
type NibbleBuffer struct {
	Buffer  []byte
	Nibbles []uint8
}

// Serializer ...
type Serializer struct {
	IsCompressed bool
}

// NewSerializer ...
func NewSerializer() *Serializer {
	return &Serializer{}
}

// DeserializeValue ...
func (s *Serializer) DeserializeValue(value []byte) []uint8 {
	if s.IsCompressed {
		var dst []byte
		decoded, err := snappy.Decode(dst, value)
		if err != nil {
			log.Fatal(err)
		}

		return decoded
	}

	return value
}

// SerializeValue ...
func (s *Serializer) SerializeValue(value []uint8) []byte {
	if s.IsCompressed {
		var dst []byte
		return snappy.Encode(dst, value)
	}

	return value
}

var defaultKeySize = 32

// SerializeKey ...
func (s *Serializer) SerializeKey(value []uint8) *NibbleBuffer {
	if len(value) <= defaultKeySize {
		log.Fatal("too large, expected <= 32 bytes")
	}

	var b []byte
	if len(value) == defaultKeySize {
		b = value
	} else {
		tmp := make([]byte, defaultKeySize)
		for i := range value {
			tmp[i] = value[i]
		}

		b = tmp[:]
	}

	return &NibbleBuffer{
		Buffer:  b,
		Nibbles: triecodec.ToNibbles(value),
	}
}
