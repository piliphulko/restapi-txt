package privacy

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/argon2"
)

type argon2Arguments struct {
	password []byte
	salt     []byte
	time     uint32
	memory   uint32
	threads  uint8
	keyLen   uint32
}

var defaultArgon2 = argon2Arguments{
	time:    3,
	memory:  64 * 1024 * 2,
	threads: 4,
	keyLen:  32,
}
var (
	ErrPasswortLong8 = errors.New("password must be 8 characters long and more")
)

func CreatePassword(password string) (string, error) {
	if len(password) < 8 {
		return "", ErrPasswortLong8
	}
	salt := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}
	//ar2Pas := argon2.IDKey([]byte(password), salt, defaultArgon2.time, defaultArgon2.memory, defaultArgon2.threads, defaultArgon2.keyLen)
	saltHex := hex.EncodeToString(salt)
	passwordHex := hex.EncodeToString([]byte(password))
	encodedRepresentation := fmt.Sprintf("argon2id salt=%s password=%s time=%d memory=%d threads=%d keyLen=%d",
		saltHex, passwordHex, defaultArgon2.time, defaultArgon2.memory, defaultArgon2.threads, defaultArgon2.keyLen)
	return hex.EncodeToString([]byte(encodedRepresentation)), nil
}

func GetHashCryptoKeyFromPassword(passwordHex string) (string, error) {
	password, err := hex.DecodeString(passwordHex)
	if err != nil {
		return "", err
	}
	var (
		argon          = argon2Arguments{}
		saltEncode     string
		passwordEncode string
	)
	if _, err := fmt.Sscanf(string(password), "argon2id salt=%s password=%s time=%d memory=%d threads=%d keyLen=%d",
		&saltEncode, &passwordEncode, &argon.time, &argon.memory, &argon.threads, &argon.keyLen); err != nil {
		return "", err
	}
	s, err := hex.DecodeString(saltEncode)
	if err != nil {
		return "", err
	}
	p, err := hex.DecodeString(passwordEncode)
	if err != nil {
		return "", err
	}

	argon.salt = s
	argon.password = p
	argon.keyLen = defaultArgon2.keyLen
	cryptoKey := argon2.IDKey(argon.password, argon.salt, argon.time, argon.memory, argon.threads, argon.keyLen)

	return hex.EncodeToString(cryptoKey), nil
}
