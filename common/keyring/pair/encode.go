package pair

import (
	"crypto/rand"

	"github.com/opennetsys/golkadot/common/crypto"
	"github.com/opennetsys/golkadot/common/u8util"
)

// Encode ...
func Encode(secretKey [64]byte, passphrase *string) ([]byte, error) {
	encoded := u8util.Concat(DEFAULT_PKCS8_HEADER, secretKey[0:32], DEFAULT_PKCS8_DIVIDER, secretKey[32:64])

	if passphrase == nil {
		return encoded, nil
	}

	nonce := [24]byte{}
	_, err := rand.Read(nonce[:])
	if err != nil {
		return nil, err
	}

	secret := [32]byte{}
	tmp := u8util.FixLength([]byte(*passphrase), 256, true)
	copy(secret[:], tmp)

	encrypted, err := crypto.NaclEncrypt(encoded, nonce, secret)
	if err != nil {
		return nil, err
	}

	return u8util.Concat(nonce[:], encrypted), nil
}
