package triedb

import (
	"reflect"
	"testing"
)

func TestCheckpoint(t *testing.T) {
	// TODO: table tests

	chkpt := NewCheckpoint([]uint8{0x1})

	txRoot := chkpt.CreateCheckpoint()
	if !reflect.DeepEqual(txRoot, []uint8{0x1}) {
		t.Fail()
	}

	rootHash := chkpt.CommitCheckpoint()
	if !reflect.DeepEqual(rootHash, []uint8{0x1}) {
		t.Fail()
	}

	rootHash = chkpt.RevertCheckpoint()
	if !reflect.DeepEqual(rootHash, []uint8{0x1}) {
		t.Fail()
	}
}
