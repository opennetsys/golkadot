package clientdb

import (
	"testing"

	clientchainloader "github.com/c3systems/go-substrate/client/chains/loader"
	clientdbtypes "github.com/c3systems/go-substrate/client/db/types"
	clienttypes "github.com/c3systems/go-substrate/client/types"
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
