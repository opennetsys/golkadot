package author

import (
	"github.com/c3systems/go-substrate/p2p/pubsub"
)

type ServiceInterface interface {
	// SubmitExtrinsic submits a hex-encoded extrinsic for inclusion in block.
	SubmitExtrinsic(extrinsic []byte) (string, error)
	// PendingExtrinsics returns all pending extrinsics, potentially grouped by sender.
	PendingExtrinsics() ([][]byte, error)
	// SubmitAndWatchExtrinsic submits an extrinsic to watch.
	SubmitAndWatchExtrinsic(subscriber pubsub.Subscriber, extrinsic []byte) error
	// UnwatchExtrinsic unsubscribes from extrinsic watching.
	UnwatchExtrinsic(id pubsub.SubscriptionID) (bool, error)
}
