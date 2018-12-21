package sync

// Interface defines the methods of the sync service
type Interface interface {
	// On handles events
	On(event EventEnum, cb EventCallback) (interface{}, error)
}
