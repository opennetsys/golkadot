package clienttypes

// TODO...

func (b *BlockRequest) Kind() uint {
	return 0
}

// Encode serializes the message into a bytes array
func (b *BlockRequest) Encode() ([]byte, error) {
	return nil, nil
}

// Decode deserializes a bytes array into a message
func (b *BlockRequest) Decode(bytes []byte) error {
	return nil
}

// Marshal returns json
func (b *BlockRequest) Marshal() ([]byte, error) {
	return nil, nil
}

// Unmarshal converts json to a message
func (b *BlockRequest) Unmarshal(bytes []byte) error {
	return nil
}

// Header ...
func (b *BlockRequest) Header() *Header {
	return nil
}
