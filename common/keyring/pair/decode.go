package pair

import (
	"errors"
	"log"

	"github.com/opennetsys/golkadot/common/crypto"
	"github.com/opennetsys/golkadot/common/u8util"
)

// Decode ...
func Decode(passphrase *string, encrypted []byte) ([32]byte, [64]byte, error) {
	var (
		naclPub  [32]byte
		naclPriv [64]byte
	)

	if encrypted == nil || len(encrypted) == 0 {
		return naclPub, naclPriv, errors.New("no encrypted data to decode")
	}

	encoded := encrypted
	if passphrase != nil {
		if len(encrypted) < 24 {
			return naclPub, naclPriv, errors.New("encrypted length is less than 24")
		}

		secret := u8util.FixLength([]byte(*passphrase), 256, true)
		if len(secret) != 32 {
			log.Println(secret, len(secret))
			return naclPub, naclPriv, errors.New("secret length is not 32")
		}

		var (
			tmpSecret [32]byte
			tmpNonce  [24]byte
			err       error
		)
		copy(tmpSecret[:], secret)
		copy(tmpNonce[:], encrypted[0:24])
		encoded, err = crypto.NaclDecrypt(encrypted[24:], tmpNonce, tmpSecret)
		if err != nil {
			return naclPub, naclPriv, err
		}
	}

	if encoded == nil || len(encoded) == 0 {
		return naclPub, naclPriv, errors.New("unable to decode")
	}

	// note: check encoded lengths?
	header := encoded[0:DEFAULT_SEED_OFFSET]
	divider := encoded[DEFAULT_DIV_OFFSET : DEFAULT_DIV_OFFSET+len(DEFAULT_PKCS8_DIVIDER)]
	if string(header) != string(DEFAULT_PKCS8_HEADER) {
		return naclPub, naclPriv, errors.New("Invalid Pkcs8 header found in body")
	}
	if string(divider) != string(DEFAULT_PKCS8_DIVIDER) {
		return naclPub, naclPriv, errors.New("Invalid Pkcs8 divider found in body")
	}

	publicKey := encoded[DEFAULT_PUBLIC_OFFSET : DEFAULT_PUBLIC_OFFSET+DEFAULT_KEY_LENGTH]
	seed := encoded[DEFAULT_SEED_OFFSET : DEFAULT_SEED_OFFSET+DEFAULT_KEY_LENGTH]
	secretKey := u8util.Concat(seed, publicKey)

	pub, priv, err := crypto.NewNaclKeyPairFromSeed(seed)
	if err != nil {
		return naclPub, naclPriv, err
	}
	if string(pub[:]) != string(publicKey) {
		return naclPub, naclPriv, errors.New("Pkcs8 decoded publicKeys are not matching")
	}
	if string(priv[:]) != string(secretKey) {
		return naclPub, naclPriv, errors.New("Pkcs8 decoded secret Keys are not matching")
	}

	naclPub = pub
	naclPriv = priv
	return naclPub, naclPriv, nil
}
