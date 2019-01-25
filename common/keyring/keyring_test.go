package keyring

import (
	"testing"

	"github.com/c3systems/go-substrate/common/keyring/address"
	"github.com/c3systems/go-substrate/common/u8util"
)

func TestKeyRing(t *testing.T) {
	publicKeyOne := [32]byte{47, 140, 97, 41, 216, 22, 207, 81, 195, 116, 188, 127, 8, 195, 230, 62, 209, 86, 207, 120, 174, 251, 74, 101, 80, 217, 123, 135, 153, 121, 119, 238}
	publicKeyTwo := [32]byte{215, 90, 152, 1, 130, 177, 10, 183, 213, 75, 254, 211, 201, 100, 7, 58, 14, 225, 114, 243, 218, 166, 35, 37, 175, 2, 26, 104, 247, 7, 81, 26}
	seedOne := []byte("12345678901234567890123456789012")
	seedTwo := u8util.FromHex("0x9d61b19deffd5a60ba844af492ec2cc44449c5697b326919703bac031cae7f60")

	kr, err := New()
	if err != nil {
		t.Fatal(err)
	}

	_, err = kr.AddFromSeed(seedOne, nil)
	if err != nil {
		t.Fatal(err)
	}

	address.SetDefaultPrefix(address.FortyTwo)

	t.Run("adds the pair", func(t *testing.T) {
		p, err := kr.AddFromSeed(seedTwo, nil)
		if err != nil {
			t.Error(err)
			return
		}
		pk, err := p.PublicKey()
		if err != nil {
			t.Error(err)
			return
		}

		if len(pk) != len(publicKeyTwo) {
			t.Errorf("expected public key length %d, received %d", len(publicKeyTwo), len(pk))
			return
		}
		for idx := range publicKeyTwo {
			if publicKeyTwo[idx] != pk[idx] {
				t.Errorf("err at idx %d; expecteded %v, received %v", idx, publicKeyTwo[idx], pk[idx])
			}
		}
	})

	t.Run("adds from a mnemonic", func(t *testing.T) {
		address.SetDefaultPrefix(address.SixtyEight)

		p, err := kr.AddFromMnemonic("moral movie very draw assault whisper awful rebuild speed purity repeat card", "", nil)
		if err != nil {
			t.Error(err)
			return
		}

		addr, err := p.Address()
		if err != nil {
			t.Error(err)
			return
		}

		expected := "7pDZKLEixRnF6Q5jzr7DsCEiNPt3d6Rknc14SyUcnRwTQK14"

		if addr != expected {
			t.Errorf("expected %s, received %s", expected, addr)
		}
	})

	t.Run("allows publicKeys retrieval", func(t *testing.T) {
		kr1, err := New()
		if err != nil {
			t.Error(err)
			return
		}

		_, err = kr1.AddFromSeed(seedOne, nil)
		if err != nil {
			t.Fatal(err)
			return
		}
		_, err = kr1.AddFromSeed(seedTwo, nil)
		if err != nil {
			t.Error(err)
			return
		}

		pks, err := kr1.GetPublicKeys()
		if err != nil {
			t.Error(err)
			return
		}

		if len(pks) != 2 {
			t.Errorf("expected 2 pks, received %d", len(pks))
			return
		}
		for idx := range pks[0] {
			if pks[0][idx] != publicKeyOne[idx] {
				t.Errorf("public key 1 does not match at %d; expected %v, received %v", idx, publicKeyOne[idx], pks[0][idx])
				return
			}
		}
		for idx := range pks[1] {
			if pks[1][idx] != publicKeyTwo[idx] {
				t.Errorf("public key 2 does not match at %d; expected %v, received %v", idx, publicKeyOne[idx], pks[1][idx])
				return
			}
		}
	})

	t.Run("allows retrieval of a specific item", func(t *testing.T) {
		addr, err := address.Encode(publicKeyOne[:], nil)
		p, err := kr.GetPair([]byte(addr))
		if err != nil {
			t.Error(err)
			return
		}
		if p == nil {
			t.Error("p is nil")
			return
		}

		pk, err := p.PublicKey()
		if err != nil {
			t.Error(err)
			return
		}

		for idx := range publicKeyOne {
			if publicKeyOne[idx] != pk[idx] {
				t.Errorf("err at idx %d; expected %v, received %v", idx, publicKeyOne[idx], pk[idx])
				return
			}
		}
	})

	t.Run("allows adding from JSON", func(t *testing.T) {
		expected := [32]byte{209, 114, 167, 76, 218, 76, 134, 89, 18, 195, 43, 160, 168, 10, 87, 174, 105, 171, 174, 65, 14, 92, 203, 89, 222, 232, 78, 47, 68, 50, 219, 79}

		p, err := kr.AddFromJSON([]byte(`{"Address":"5GoKvZWG5ZPYL1WUovuHW3zJBWBP5eT8CbqjdRY4Q6iMaDtZ","Encoded":"3053020101300506032b657004220420416c696365202020202020202020202020202020202020202020202020202020a123032100d172a74cda4c865912c32ba0a80a57ae69abae410e5ccb59dee84e2f4432db4f","Encoding":{"Content":"PKCS8","Type":"none","Version":"0"},"Meta":{"isTesting":true,"name":"alice"}}`), nil)
		if err != nil {
			t.Error(err)
			return
		}
		if p == nil {
			t.Error("pair is nil")
			return
		}

		pk, err := p.PublicKey()
		if err != nil {
			t.Error(err)
			return
		}

		for idx := range expected {
			if expected[idx] != pk[idx] {
				t.Errorf("err at idx %d; expected %v, received %v", idx, expected[idx], pk[idx])
				return
			}
		}

		password := "password"
		p, err = kr.AddFromJSON([]byte(`{"Address":"5GoKvZWG5ZPYL1WUovuHW3zJBWBP5eT8CbqjdRY4Q6iMaDtZ","Encoded":"7accc3662c519b2def55141e805833876ba0222452b93c9f595fa208e1540631b288f58ebed5e4e7639f977cd4ecf4a5d7947339a71a50576eb6b4a49fab4e8777caf8a788f0051bfb123828430a953391c98b63fca6d925e0ccf0ceaa1824147b0baf26adf47b9a0f0494aad8eda61ee4ad2022be781a6b83d0c1139b","Encoding":{"Content":"PKCS8","Type":"xsalsa20-poly1305","Version":"0"},"Meta":{"isTesting":true,"name":"alice"}}`), &password)
		if err != nil {
			t.Error(err)
			return
		}
		if p == nil {
			t.Error("pair is nil")
			return
		}

		pk, err = p.PublicKey()
		if err != nil {
			t.Error(err)
			return
		}

		for idx := range expected {
			if expected[idx] != pk[idx] {
				t.Errorf("err at idx %d; expected %v, received %v", idx, expected[idx], pk[idx])
				return
			}
		}
	})
}
