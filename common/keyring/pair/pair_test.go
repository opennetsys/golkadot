package pair

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestToFromJSON(t *testing.T) {
	password := "password"
	pMap, err := testKeyringPairs()
	if err != nil {
		t.Fatal(err)
	}
	for i, tt := range []struct {
		name     string
		password *string
	}{
		{"alice", nil},
		{"alice", &password},
		{"bob", nil},
		{"bob", &password},
		{"charlie", nil},
		{"charlie", &password},
		{"dave", nil},
		{"dave", &password},
		{"eve", nil},
		{"eve", &password},
		{"ferdie", nil},
		{"ferdie", &password},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			p, ok := pMap[tt.name]
			if !ok {
				t.Fatal("err getting key")
			}

			jsn, err := p.ToJSON(tt.password)
			if err != nil {
				t.Fatal(err)
			}
			//t.Log(string(jsn))

			tmpP, err := NewPairFromJSON(jsn, tt.password)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(p, tmpP) {
				t.Errorf("expected %v\nreceived %v", p, tmpP)
			}
		})
	}
}

func TestPair(t *testing.T) {
	sig := []byte{80, 191, 198, 147, 225, 207, 75, 88, 126, 39, 129, 109, 191, 38, 72, 181, 75, 254, 81, 143, 244, 79, 237, 38, 236, 141, 28, 252, 134, 26, 169, 234, 79, 33, 153, 158, 151, 34, 175, 188, 235, 20, 35, 135, 83, 120, 139, 211, 233, 130, 1, 208, 201, 215, 73, 80, 56, 98, 185, 196, 11, 8, 193, 14}

	pMap, err := testKeyringPairs()
	if err != nil {
		t.Fatal(err)
	}

	alice, ok := pMap["alice"]
	if !ok {
		t.Fatal("err getting alice")
	}
	jsn, err := alice.ToJSON(nil)
	if err != nil {
		t.Fatal(err)
	}
	password := "password"
	jsn, err = alice.ToJSON(&password)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(jsn))

	t.Run("has a publicKey", func(t *testing.T) {
		apk, err := alice.PublicKey()
		if err != nil {
			t.Error(err)
			return
		}

		expected := []byte{209, 114, 167, 76, 218, 76, 134, 89, 18, 195, 43, 160, 168, 10, 87, 174, 105, 171, 174, 65, 14, 92, 203, 89, 222, 232, 78, 47, 68, 50, 219, 79}

		// note: why does reflect.DeepEqual fail, here?
		if len(apk) != len(expected) {
			t.Errorf("expected pub key len %d, received %d", len(expected), len(apk))
			return
		}
		for idx := range expected {
			if expected[idx] != apk[idx] {
				t.Errorf("err at idx %d; expected %v, received %v", idx, expected[idx], apk[idx])
				return
			}
		}
	})

	t.Run("allows signing", func(t *testing.T) {
		aSig, err := alice.Sign([]byte{0x61, 0x62, 0x63, 0x64})
		if err != nil {
			t.Error(err)
			return
		}

		if !reflect.DeepEqual(aSig, sig) {
			t.Errorf("expected %v\nreceived %v", sig, aSig)
			return
		}
	})

	t.Run("validates a correctly signed message", func(t *testing.T) {
		ok, err := alice.Verify([]byte{0x61, 0x62, 0x63, 0x64}, sig)
		if err != nil {
			t.Error(err)
			return
		}

		if !ok {
			t.Errorf("expected %v\nreceived %v", true, ok)
			return
		}
	})

	t.Run("fails a correctly signed message (message changed)", func(t *testing.T) {
		ok, err := alice.Verify([]byte{0x61, 0x62, 0x63, 0x64, 0x65}, sig)
		if err != nil {
			t.Error(err)
			return
		}

		if ok {
			t.Errorf("expected %v\nreceived %v", false, ok)
			return
		}
	})
}
