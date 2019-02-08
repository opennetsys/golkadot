package pair

import (
	"github.com/opennetsys/godot/common/crypto"
	keytypes "github.com/opennetsys/godot/common/keyring/types"
)

var seeds = map[string][]byte{
	"alice":   padSeed("Alice"),
	"bob":     padSeed("Bob"),
	"charlie": padSeed("Charlie"),
	"dave":    padSeed("Dave"),
	"eve":     padSeed("Eve"),
	"ferdie":  padSeed("Ferdie"),
}

func padSeed(seed string) []byte {
	b := []byte(seed)
	for i := 0; len(b) < 32; i++ {
		b = append(b, []byte(" ")...)
	}

	return b
}

// TestKeyRing ...
//func TestKeyRing() (*keyring.KeyRing, error) {
//kr, err := keyring.New()
//if err != nil {
//return nil, err
//}
//for name := range seeds {
//meta := make(keytypes.Meta)
//meta["isTesting"] = true
//meta["name"] = name

//_, err := kr.AddFromSeed(seeds[name], &meta)
//if err != nil {
//return nil, err
//}
//}

//return kr, nil
//}

func testKeyringPairs() (MapPair, error) {
	m := make(MapPair)

	//keyring := TestKeyRing()
	//pairs := keyring.GetPairs()
	for name := range seeds {
		meta := make(keytypes.Meta)
		meta["isTesting"] = true
		meta["name"] = name

		pub, priv, err := crypto.NewNaclKeyPairFromSeed(seeds[name])
		if err != nil {
			return nil, err
		}

		pair, err := NewPair(pub, priv, meta, nil)
		if err != nil {
			return nil, err
		}

		m[name] = pair
	}

	return m, nil
}
