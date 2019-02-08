package clientdb

import (
	"testing"

	clientchainloader "github.com/opennetsys/godot/client/chain/loader"
	clientdbtypes "github.com/opennetsys/godot/client/db/types"
	clienttypes "github.com/opennetsys/godot/client/types"
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
