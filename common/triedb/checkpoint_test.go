package triedb

import (
	"reflect"
	"testing"
)

func TestCheckpoint(t *testing.T) {
	// TODO: table tests

	hash := []uint8{0x1}
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
