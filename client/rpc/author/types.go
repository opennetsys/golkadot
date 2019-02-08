package author

import "github.com/opennetsys/godot/client/p2p/pubsub"

// Service implements the author service
type Service struct{}

// SubmitAndWatchExtrinsicArgs is passed to the SubmitAndWatchExtrinsic method
type SubmitAndWatchExtrinsicArgs struct {
	// Subscriber ...
	Subscriber pubsub.Subscriber
	// Extrinsic ...
	Extrinsic []byte
}
