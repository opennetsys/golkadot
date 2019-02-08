package clientchain

import (
	"testing"

	"github.com/opennetsys/golkadot/common/u8util"
)

func TestChains(t *testing.T) {
	var chain *Chain
	// TODO: fix triedb package
	/*
		chain := NewChain(&clienttypes.ConfigClient{
			Chain: "dev",
			DB: &clientdbtypes.Config{
				Type: "memory",
			},
		})
	*/

	t.Run("creates a correct genesis block (parentHash)", func(t *testing.T) {
		t.Skip()
		got := u8util.ToHex(chain.Genesis.Block.Header.ParentHash[:], -1, true)
		want := "0x0000000000000000000000000000000000000000000000000000000000000000"
		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("creates a correct genesis block (stateRoot)", func(t *testing.T) {
		t.Skip()
		got := u8util.ToHex(chain.Genesis.Block.Header.StateRoot[:], -1, true)
		want := "0xf51fa1968b1f74ae72b91c8e4a73658633b6f663a30d7f4f1c2e4436c717d4e8"
		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("creates a correct genesis block (extrinsicsRoot)", func(t *testing.T) {
		t.Skip()
		got := u8util.ToHex(chain.Genesis.Block.Header.ExtrinsicsRoot[:], -1, true)
		want := "0x03170a2e7597b7b7e3d84c05391d139a62b157e78786d8c082f29dcf4c111314"
		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("creates a correct block hash", func(t *testing.T) {
		t.Skip()
		got := u8util.ToHex(chain.Genesis.Block.Hash[:], -1, true)
		want := "0x3e66d3b17a316ecf9bf3fc35e4131500d1dbb41de3ee210f7a680a2c5327b490"
		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}
