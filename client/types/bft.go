package clienttypes

import (
	handlertypes "github.com/c3systems/go-substrate/client/p2p/handler/types"
)

// TODO...

// Kind ...
func (b *BFT) Kind() handlertypes.FuncEnum {
	return handlertypes.BFT
}

// Encode serializes the message into a bytes array
func (b *BFT) Encode() ([]byte, error) {
	return nil, nil
}

// Decode deserializes a bytes array into a message
func (b *BFT) Decode(bytes []byte) error {
	return nil
}

// Marshal returns json
func (b *BFT) Marshal() ([]byte, error) {
	return nil, nil
}

// Unmarshal converts json to a message
func (b *BFT) Unmarshal(bytes []byte) error {
	return nil
}

// Header ...
func (b *BFT) Header() *Header {
	return nil
}
