package status

// Status ...
// TODO: this needs to implement the message interface
// https://github.com/polkadot-js/client/blob/master/packages/client-types/src/messages/Status.ts
type Status struct {
	Roles       []string
	BestNumber  *math.Big
	BestHash    []byte
	GenesisHash []byte
	Version     string
}
