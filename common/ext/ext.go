package ext

var (
	// Assert ...
	Assert = -90009 // Js client
	// Unknown ...
	Unknown = -1
	// InvalidJSONRPC ...
	InvalidJSONRPC = -99998 // Js client
	// MethodNotFound ...
	MethodNotFound = -32601 // Rust client
)

// Error ...
type Error struct {
	Message string
	Code    int
}

// NewExtError ...
func NewExtError(message string, code int) *Error {
	return &Error{
		Message: message,
		Code:    code,
	}
}
