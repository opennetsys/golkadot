package triedb

import (
	"reflect"
	"testing"

	"github.com/c3systems/go-substrate/common/crypto"
)

func TestCheckpoint(t *testing.T) {
	// TODO: table tests

	hash := new(crypto.Blake2b256Hash)
	copy(hash[:], []uint8{0x1})

	chkpt := NewCheckpoint(hash)

	txRoot := chkpt.CreateCheckpoint()
	if !reflect.DeepEqual(txRoot, hash) {
		t.Fail()
	}

	rootHash := chkpt.CommitCheckpoint()
	if !reflect.DeepEqual(rootHash, hash) {
		t.Fail()
	}

	rootHash = chkpt.RevertCheckpoint()
	if !reflect.DeepEqual(rootHash, hash) {
		t.Fail()
	}
}
