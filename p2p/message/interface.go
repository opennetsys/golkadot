package message

// Interface defines the methods of Message
type Interface interface {
	// Kind returns the message's kind
	Kind() uint
	// Encode serializes the message into a bytes array
	Encode() ([]byte, error)
	// Decode deserializes a bytes array into a message
	Decode(bytes []byte) error
	// Marshal returns json
	Marshal() ([]byte, error)
	// Unmarshal converts json to a message
	Unmarshal(bytes []byte) error
	// Header ...
	Header() Header
}
