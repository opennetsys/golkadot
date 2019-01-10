package message

import "math/big"

// TODO: implement https://github.com/polkadot-js/api/blob/master/packages/types/src/Header.ts
type Header struct {
	BlockNumber *big.Int
	Hash        []byte
}
