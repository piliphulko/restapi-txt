package privacy

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

var (
	ErrNoKey          = errors.New("the key is not inserted into the variable: privacy.Keystore")
	ErrCipherTooShort = errors.New("ciphertext too short")
)

func EncryptString(text string) (string, error) {
	if Keystore.block == nil {
		return "", ErrNoKey
	}
	//src := []byte(text)
	//dst := make([]byte, hex.EncodedLen(len(text)))
	//hex.Encode(dst, src)

	chiphertext := make([]byte, aes.BlockSize+len(text))
	iv := chiphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(Keystore.block, iv)
	stream.XORKeyStream(chiphertext[aes.BlockSize:], []byte(text))

	src := []byte(chiphertext)
	dst := make([]byte, hex.EncodedLen(len(chiphertext)))
	hex.Encode(dst, src)

	return string(dst), nil
}

func DecryptString(text string) (string, error) {
	if Keystore.block == nil {
		return "", ErrNoKey
	}
	src := []byte(text)
	ciphertext := make([]byte, hex.DecodedLen(len(text)))
	hex.Decode(ciphertext, src)

	if len(ciphertext) < aes.BlockSize {
		return "", ErrCipherTooShort
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(Keystore.block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}
