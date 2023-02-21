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

var (
	Keystore keysUsing
)

func (ku *keysUsing) InserKey(keyHex string) {
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		log.Fatal(err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	ku.block = block
}
