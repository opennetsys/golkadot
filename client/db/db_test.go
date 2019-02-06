package clientdb

import (
	"testing"

	clientchainloader "github.com/opennetsys/go-substrate/client/chain/loader"
	clientdbtypes "github.com/opennetsys/go-substrate/client/db/types"
	clienttypes "github.com/opennetsys/go-substrate/client/types"
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
}
