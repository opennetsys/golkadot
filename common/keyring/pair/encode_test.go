package pair

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	password := "testing"
	pMap, err := testKeyringPairs()
	if err != nil {
		t.Fatal(err)
	}

	for i, tt := range []struct {
		in  *string
		out int
	}{
		{
			nil,
			len(DEFAULT_PKCS8_DIVIDER) + len(DEFAULT_PKCS8_HEADER) + 64,
		},
		{
			&password,
			125,
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			alice, ok := pMap["alice"]
			if !ok {
				t.Fatal("err getting alice")
			}

			out, err := alice.EncodePkcs8(tt.in)
			if err != nil {
				t.Error(err)
			}

			if len(out) != tt.out {
				t.Errorf("expected %d, received %d", tt.out, len(out))
			}
		})
	}
}
