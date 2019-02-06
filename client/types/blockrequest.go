package clienttypes

import (
	handlertypes "github.com/c3systems/go-substrate/client/p2p/handler/types"
)

// TODO...

// Kind ...
func (b *BlockRequest) Kind() handlertypes.FuncEnum {
	return handlertypes.BlockRequest
}

// Encode serializes the message into a bytes array
func (b *BlockRequest) Encode() ([]byte, error) {
	return nil, nil
}

// Decode deserializes a bytes array into a message
func (b *BlockRequest) Decode(bytes []byte) error {
	return nil
}

// MarshalJSON returns json
func (b *BlockRequest) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON converts json to a message
func (b *BlockRequest) UnmarshalJSON(bytes []byte) error {
	return nil
}

// Header ...
func (b *BlockRequest) Header() *Header {
	return nil
}
