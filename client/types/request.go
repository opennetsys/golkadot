package clienttypes

// TODO...

func (r *Request) Kind() uint {
	return 0
}

// Encode serializes the message into a bytes array
func (r *Request) Encode() ([]byte, error) {
	return nil, nil
}

// Decode deserializes a bytes array into a message
func (r *Request) Decode(bytes []byte) error {
	return nil
}

// Marshal returns json
func (r *Request) Marshal() ([]byte, error) {
	return nil, nil
}

// Unmarshal converts json to a message
func (r *Request) Unmarshal(bytes []byte) error {
	return nil
}

// Header ...
func (r *Request) Header() *Header {
	return nil
}
