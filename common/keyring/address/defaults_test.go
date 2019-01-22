package address

import "testing"

func TestSetDefaultPrefix(t *testing.T) {
	SetDefaultPrefix(SixtyEight)

	result, err := Encode([]byte{1}, nil)
	if err != nil {
		t.Error(err)
	}

	if result != "Pqt7" {
		t.Errorf("expected: Pqt7, received: %s", result)
	}
}
