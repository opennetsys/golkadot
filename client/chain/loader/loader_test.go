package chainloader

import (
	"testing"

	clienttypes "github.com/opennetsys/go-substrate/client/types"
	"github.com/opennetsys/go-substrate/common/u8util"
)

func TestLoader(t *testing.T) {
	t.Run("default chain (dev)", func(t *testing.T) {
		loader := NewLoader(&clienttypes.ConfigClient{})

		root := loader.CalculateGenesisRoot()
		got := u8util.ToHex(root, -1, true)
		want := "0xf51fa1968b1f74ae72b91c8e4a73658633b6f663a30d7f4f1c2e4436c717d4e8"

		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("chain json path (dev)", func(t *testing.T) {
		loader := NewLoader(&clienttypes.ConfigClient{
			Chain: "../json/dev.json",
		})

		root := loader.CalculateGenesisRoot()
		got := u8util.ToHex(root, -1, true)
		want := "0xf51fa1968b1f74ae72b91c8e4a73658633b6f663a30d7f4f1c2e4436c717d4e8"

		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}
