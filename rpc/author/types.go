package author

import "github.com/c3systems/go-substrate/p2p/pubsub"

// Service implements the author service
type Service struct{}

// SubmitAndWatchExtrinsicArgs is passed to the SubmitAndWatchExtrinsic method
type SubmitAndWatchExtrinsicArgs struct {
	Subscriber pubsub.Subscriber
	Extrinsic  []byte
}
