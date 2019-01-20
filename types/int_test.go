package types

import (
	"testing"
)

func TestInt(t *testing.T) {
	t.Run("provides a toBn interface", func(t *testing.T) {
		if NewInt(-1234).Int64() != -1234 {
			t.Fail()
		}
	})
}
