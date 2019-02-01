package clienttypes

import (
	handlertypes "github.com/c3systems/go-substrate/client/p2p/handler/types"
)

// TODO...

// Kind ...
func (b *BlockAnnounce) Kind() handlertypes.FuncEnum {
	return handlertypes.BlockAnnounce
}

// Encode serializes the message into a bytes array
func (b *BlockAnnounce) Encode() ([]byte, error) {
	return nil, nil
}

// Decode deserializes a bytes array into a message
func (b *BlockAnnounce) Decode(bytes []byte) error {
	return nil
}

// Marshal returns json
func (b *BlockAnnounce) Marshal() ([]byte, error) {
	return nil, nil
}

// Unmarshal converts json to a message
func (b *BlockAnnounce) Unmarshal(bytes []byte) error {
	return nil
}

// Header ...
func (b *BlockAnnounce) Header() *Header {
	return nil
}
