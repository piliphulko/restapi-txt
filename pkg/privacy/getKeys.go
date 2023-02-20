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

func (ku *keysUsing) InserKey(key string) {
	key = hex.EncodeToString([]byte(key))
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	ku.block = block
}
