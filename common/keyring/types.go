package keyring

import "github.com/c3systems/go-substrate/common/keyring/pair"

var (
	// note: ensure the struct(s) implement the interface(s) at compile time
	_ InterfaceKeyRing = (*KeyRing)(nil)
)

// KeyRing ...
type KeyRing struct {
	Pairs pair.InterfacePairs
}
