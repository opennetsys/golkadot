package state

import (
	"github.com/c3systems/go-substrate/client/p2p/pubsub"
	rpctypes "github.com/c3systems/go-substrate/client/rpc/types"
	"github.com/c3systems/go-substrate/client/runtime"
	"github.com/c3systems/go-substrate/common/storage"
)

// ServiceInterface described the methods performed by the state rpc api.
type ServiceInterface interface {
	// Call is used to call a contract at a block's state.
	Call(args *CallArgs, response []byte) error
	// GetStorage returns a storage entry at a specific block's state.
	GetStorage(args *StorageArgs, response *storage.Data) error
	// GetStorageHash returns the hash of a storage entry at a block's state.
	GetStorageHash(args *StorageArgs, response *string) error
	// GetStorageSize returns the size of a storage entry at a block's state.
	GetStorageSize(args *StorageArgs, response *uint64) error
	// GetMetadata returns the runtime metadata as an opaque blob.
	GetMetadata(blockHash *string, response []byte) error
	// GetRuntimeVersion returns the runtime version.
	GetRuntimeVersion(blockHash *string, response *runtime.Version) error
	// QueryStorage queries historical storage entries (by key) starting from a block given as the second parameter.
	//
	// NOTE: This first returned result contains the initial state of storage for all keys.
	// Subsequent values in the vector represent changes to the previous state (diffs).
	QueryStorage(args *QueryStorageArgs, response *storage.ChangeSet) error
	// SubscribeRuntimeVersion subscribes to a new runtime version subscription.
	SubscribeRuntimeVersion(subscriber pubsub.Subscriber, response rpctypes.NilResponse) error
	// UnsubscribeRuntimeVersion unsubscribes from runtime version subscription.
	UnsubscribeRuntimeVersion(id pubsub.SubscriptionID, response *bool) error
	// SubscribeStorage subscribes to a new storage subscription.
	SubscribeStorage(args *SubscribeStorageArgs, response rpctypes.NilResponse) error
	// UnsubscribeStorage unsubscribes from storage subscription.
	UnsubscribeStorage(id pubsub.SubscriptionID, response *bool) error
}
