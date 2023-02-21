package privacy

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var letterBytes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func newWord(length int) string {
	word := make([]byte, length)
	for i := 0; i != length; i++ {
		word[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(word)
}

func TestTotal(t *testing.T) {
	_, err := EncryptString("")
	require.ErrorContains(t, err, ErrNoKey.Error())
	_, err = DecryptString("")
	require.ErrorContains(t, err, ErrNoKey.Error())

	_, err = CreatePassword("1234567")
	require.ErrorContains(t, err, ErrPasswortLong8.Error())
	p, err := CreatePassword("12345678")
	require.Nil(t, err)

	CryptoKey, err := GetHashCryptoKeyFromPassword(p)
	require.Nil(t, err)

	Keystore.InserKey(CryptoKey)
	text, err := EncryptString("some text")
	require.Nil(t, err)

	text, err = DecryptString(text)
	require.Nil(t, err)
	require.Equal(t, text, "some text")
}

func TestCreateGet(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i != 100; i++ {
		t.Run("", func(t *testing.T) {
			a, err := CreatePassword(newWord(rand.Intn(5) + 8))
			require.Nil(t, err)
			_, err = GetHashCryptoKeyFromPassword(a)
			require.Nil(t, err)
		})
	}
}
