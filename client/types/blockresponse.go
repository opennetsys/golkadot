package clienttypes

// TODO...

func (b *BlockResponse) Kind() uint {
	return 0
}

// Encode serializes the message into a bytes array
func (b *BlockResponse) Encode() ([]byte, error) {
	return nil, nil
}

// Decode deserializes a bytes array into a message
func (b *BlockResponse) Decode(bytes []byte) error {
	return nil
}

// Marshal returns json
func (b *BlockResponse) Marshal() ([]byte, error) {
	return nil, nil
}

// Unmarshal converts json to a message
func (b *BlockResponse) Unmarshal(bytes []byte) error {
	return nil
}

// Header ...
func (b *BlockResponse) Header() *Header {
	return nil
}
