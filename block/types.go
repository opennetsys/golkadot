package block

import (
	"math/big"

	pcrypto "github.com/c3systems/go-substrate/common/crypto"
)

// AccountID ...
type AccountID [32]uint8

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
	ParentHash     *pcrypto.Blake2b256Hash
	Number         *big.Int
	StateRoot      *pcrypto.Blake2b256Hash
	ExtrinsicsRoot *pcrypto.Blake2b256Hash
	Digest         *Digest
}

// Request TODO
type Request struct {
	From *big.Int
	ID   uint
	Max  uint64
}

// Data TODO
type Data struct {
	Hash          *pcrypto.Blake2b256Hash
	Header        *Header
	Body          []byte
	Receipt       []byte
	MessageQueue  []byte
	Justification []byte
}
