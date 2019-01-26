package types

// InterfaceType ...
type InterfaceType interface {
	Len() int
	EncodedLen() int
	String() string
	Hex() string
	Bytes() []byte
	ToU8a(isBare bool) []uint8
	Equals(other interface{}) bool
}

// BlockDB ...
// TODO
type BlockDB struct{}

// Config ...
// TODO
type Config struct{}

// StateDB ...
// TODO
type StateDB struct{}

// BlockData ...
// TODO
type BlockData struct{}
