package peers

// EventCallback is a function that is called on a peer event
type EventCallback func(p interface{}) (interface{}, error)
