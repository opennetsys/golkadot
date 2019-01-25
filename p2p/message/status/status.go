package status

import "math/big"

// Status ...
// TODO: this needs to implement the message interface
// https://github.com/polkadot-js/client/blob/master/packages/client-types/src/messages/Status.ts
type Status struct {
	Roles       []string
	BestNumber  *big.Int
	BestHash    []byte
	GenesisHash []byte
	Version     string
}
