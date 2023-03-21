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
// the password must be generated using the function CreatePassword

func (ku *keysUsing) InserKey(keyHex string) {
	key, err := GetHashCryptoKeyFromPassword(keyHex)
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
