package clientdb

import (
	"testing"

	clientchainloader "github.com/opennetsys/golkadot/client/chain/loader"
	clientdbtypes "github.com/opennetsys/golkadot/client/db/types"
	clienttypes "github.com/opennetsys/golkadot/client/types"
)

func TestClientDB(t *testing.T) {
	db := NewDB(&clienttypes.ConfigClient{
		DB: &clientdbtypes.Config{
			Compact:  false,
			Snapshot: false,
			IsTrieDB: false,
			Type:     "memory",
		},
	}, &clientchainloader.Loader{})
	_ = db

	// TODO
}
