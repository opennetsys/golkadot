package clienttypes

import (
	handlertypes "github.com/c3systems/go-substrate/client/p2p/handler/types"
)

// TODO...

// Kind ...
func (t *Transactions) Kind() handlertypes.FuncEnum {
	return handlertypes.Transactions
}

// Encode serializes the message into a bytes array
func (t *Transactions) Encode() ([]byte, error) {
	return nil, nil
}

// Decode deserializes a bytes array into a message
func (t *Transactions) Decode(bytes []byte) error {
	return nil
}

// MarshalJSON returns json
func (t *Transactions) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON converts json to a message
func (t *Transactions) UnmarshalJSON(bytes []byte) error {
	return nil
}

// Header ...
func (t *Transactions) Header() *Header {
	return nil
}
