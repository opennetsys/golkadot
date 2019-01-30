package clienttypes

import (
	handlertypes "github.com/c3systems/go-substrate/client/p2p/handler/types"
)

// TODO...

// Kind ...
func (s *StateRequest) Kind() handlertypes.FuncEnum {
	return handlertypes.StateRequest
}

// Encode serializes the message into a bytes array
func (s *StateRequest) Encode() ([]byte, error) {
	return nil, nil
}

// Decode deserializes a bytes array into a message
func (s *StateRequest) Decode(bytes []byte) error {
	return nil
}

// Marshal returns json
func (s *StateRequest) Marshal() ([]byte, error) {
	return nil, nil
}

// Unmarshal converts json to a message
func (s *StateRequest) Unmarshal(bytes []byte) error {
	return nil
}

// Header ...
func (s *StateRequest) Header() *Header {
	return nil
}
