package assert

import (
	"testing"
)

func TestAssert(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if r != "this should panic" {
				t.Error("expected error")
			}
		} else {
			t.Error("expected panic")
		}
	}()

	Assert(false, "this should panic")
	Assert(true, "this should not panic")
}
