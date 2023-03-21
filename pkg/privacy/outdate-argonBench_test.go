package privacy

import (
	"math/rand"
	"testing"
	"time"
)

var compilerGoof interface{}

func BenchmarkArgonParams(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		pas, _ := createPassword(newWord(8))
		cryptoKey, _ := getHashCryptoKeyFromPassword(pas)
		compilerGoof = cryptoKey
	}
}

// 	time:    1, memory:  64 * 1024, threads: 4, || 31752570 ns/op	67121678 B/op	      93 allocs/op | 32ms
//  time:    3, memory:  64 * 1024, threads: 4, || 75329893 ns/op	67125828 B/op	     138 allocs/op | 75ms
//  time:    3, memory:  64 * 1024 * 2, threads: 4, || 153524186 ns/op	134234173 B/op	     138 allocs/op | 154ms
