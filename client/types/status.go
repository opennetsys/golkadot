package clienttypes

// Kind ...
func (s *Status) Kind() uint {
	return 0
}

// Encode serializes the message into a bytes array
func (s *Status) Encode() ([]byte, error) {
	return nil, nil
}

// Decode deserializes a bytes array into a message
func (s *Status) Decode(bytes []byte) error {
	return nil
}

// Marshal returns json
func (s *Status) Marshal() ([]byte, error) {
	return nil, nil
}

// Unmarshal converts json to a message
func (s *Status) Unmarshal(bytes []byte) error {
	return nil
}

// Header ...
func (s *Status) Header() *Header {
	return nil
}
