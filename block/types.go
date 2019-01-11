package block

import (
	pcrypto "github.com/c3systems/go-substrate/common/crypto"
)

// AccountID ...
type AccountID []uint8

// AuthorityID ...
type AuthorityID AccountID

// AuthoritiesChangeObj ...
// note: obj suffix is required so as to not interfere with the enum
type AuthoritiesChangeObj []*AuthorityID

// ChangesTrieRootObj ...
// note: obj suffix is required so as to not interfere with the enum
type ChangesTrieRootObj pcrypto.Hash

// SealObj ...
// note: obj suffix is required so as to not interfere with the enum
type SealObj struct {
}

// OtherObj ...
// note: obj suffix is required so as to not interfere with the enum
type OtherObj []byte

// DigestItem ..
type DigestItem map[DigestEnum]interface{}

// Digest ...
type Digest struct {
	Logs []*DigestItem
}

// Header ...
type Header struct {
	ParentHash     pcrypto.Hash
	Number         uint64
	StateRoot      pcrypto.Hash
	ExtrinsicsRoot pcrypto.Hash
	Digest         *Digest
}

// Request TODO
type Request struct{}

// Data TODO
type Data struct {
	Hash pcrypto.Hash
}
