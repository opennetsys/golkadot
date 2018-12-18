package state

import (
	"github.com/c3systems/go-substrate/common/storage"
	"github.com/c3systems/go-substrate/p2p/pubsub"
	"github.com/c3systems/go-substrate/runtime"
)

// ServiceInterface described the methods performed by the state rpc api.
type ServiceInterface interface {
	// Call is used to call a contract at a block's state.
	Call(method string, data []byte, blockHash *string) ([]byte, error)
	// GetStorage returns a storage entry at a specific block's state.
	GetStorage(key storage.Key, blockHash *string) (*storage.Data, error)
	// GetStorageHash returns the hash of a storage entry at a block's state.
	GetStorageHash(key storage.Key, blockHash *string) (string, error)
	// GetStorageSize returns the size of a storage entry at a block's state.
	GetStorageSize(key storage.Key, blockHash *string) (uint64, error)
	// GetMetadata returns the runtime metadata as an opaque blob.
	GetMetadata(blockHash *string) ([]byte, error)
	// GetRuntimeVersion returns the runtime version.
	GetRuntimeVersion(blockHash *string) (runtime.Version, error)
	// QueryStorage queries historical storage entries (by key) starting from a block given as the second parameter.
	//
	// NOTE: This first returned result contains the initial state of storage for all keys.
	// Subsequent values in the vector represent changes to the previous state (diffs).
	QueryStorage(keys []*storage.Key, FromBlockHash string, ToBlockHash *string) (*storage.ChangeSet, error)
	// SubscribeRuntimeVersion subscribes to a new runtime version subscription.
	SubscribeRuntimeVersion(subscriber pubsub.Subscriber) error
	// UnsubscribeRuntimeVersion unsubscribes from runtime version subscription.
	UnsubscribeRuntimeVersion(id pubsub.SubscriptionID) (bool, error)
	// SubscribeStorage subscribes to a new storage subscription.
	SubscribeStorage(subscriber pubsub.Subscriber, keys []*storage.Key) error
	// UnsubscribeStorage unsubscribes from storage subscription.
	UnsubscribeStorage(id pubsub.SubscriptionID) (bool, error)
}
