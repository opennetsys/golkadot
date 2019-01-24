package pair

import "testing"

func TestDecode(t *testing.T) {
	pMap, err := testKeyringPairs()
	if err != nil {
		t.Fatal(err)
	}
	alice, ok := pMap["alice"]
	if !ok {
		t.Fatal("err getting alice")
	}
	password := "testing"

	// 1. it should fail decoding with no encoded data
	if err := alice.DecodePkcs8(nil, nil); err == nil {
		t.Error("expected err")
	}

	// 2. should pass when given data
	enc, err := alice.EncodePkcs8(&password)
	if err != nil {
		t.Fatal(err)
	}

	if err := alice.DecodePkcs8(&password, enc); err != nil {
		t.Error(err)
	}
}
