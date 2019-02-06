package chain

import (
	"github.com/opennetsys/go-substrate/client/p2p/pubsub"
	rpctypes "github.com/opennetsys/go-substrate/client/rpc/types"
	"github.com/opennetsys/go-substrate/client/runtime"
)

// ServiceInterface defines the methods implemented by the chain service
type ServiceInterface interface {
	// GetHeader returns the header of a relay chain block.
	GetHeader(blockHash *string, response *runtime.Header) error
	// GetBlock returns the header and body of a relay chain block.
	GetBlock(blockHash *string, response *runtime.SignedBlock) error
	// GetBlockHash returns the hash of the n-th block in the canon chain.
	//
	// By default returns latest block hash.
	GetBlockHash(blockHash *string, response *string) error
	// GetFinalizedHead returns the hash of the last finalised block in the canon chain.
	GetFinalizedHead(args rpctypes.NilArgs, response *string) error
	// SubscribeNewHead creates a new head subscription.
	SubscribeNewHead(subscriber pubsub.Subscriber, response rpctypes.NilResponse) error
	// UnsubscribeNewHead unsubscribes from new head subscription.
	UnsubscribeNewHead(id pubsub.SubscriptionID, response *bool) error
	// SubscribeFinalizedHeads returns a new head subscription
	SubscribeFinalizedHeads(subscriber pubsub.Subscriber, response rpctypes.NilResponse) error
	// UnsubscribeFinalizedHeads from new head subscription.
	UnsubscribeFinalizedHeads(id pubsub.SubscriptionID, response *bool) error
}
