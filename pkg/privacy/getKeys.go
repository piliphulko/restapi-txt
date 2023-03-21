package privacy

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
)

type keysUsing struct {
	block cipher.Block
}

// Keystore The variable into which you need to insert the key
// through the method InserKey

var Keystore keysUsing

// InserKey inserting a key into a variable Keystore
// the password

func (ku *keysUsing) InserKey(key string) {
	block, err := aes.NewCipher([]byte(hex.EncodeToString([]byte(key))))
	if err != nil {
		log.Fatal(err)
	}
	ku.block = block
}

func (ku *keysUsing) inserKey(keyHex string) {
	key, err := getHashCryptoKeyFromPassword(keyHex)
	if err != nil {
		log.Fatal(err)
	}
	keyHash, err := hex.DecodeString(key)
	if err != nil {
		log.Fatal(err)
	}
	block, err := aes.NewCipher(keyHash)
	if err != nil {
		log.Fatal(err)
	}
	ku.block = block
}
