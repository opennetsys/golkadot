package types

// Null ...
type Null struct {
}

// NewNull ...
func NewNull() *Null {
	return &Null{}
}

// Len ...
func (n *Null) Len() int {
	return 0
}

// EncodedLen ...
func (n *Null) EncodedLen() int {
	return 0
}

// String ...
func (n *Null) String() string {
	return ""
}

// Hex ...
func (n *Null) Hex() string {
	return "0x"
}

// Bytes ...
func (n *Null) Bytes() []byte {
	return nil
}

// ToU8a ...
func (n *Null) ToU8a(isBare bool) []byte {
	return nil
}

// Equals ...
func (n *Null) Equals(other interface{}) bool {
	switch other.(type) {
	case *Null:
		return true
	}

	return false
}
