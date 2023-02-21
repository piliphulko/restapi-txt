package main

import (
	"fmt"
	"log"

	_ "github.com/piliphulko/practiceGo/pkg/datalog"
	"github.com/piliphulko/practiceGo/pkg/privacy"
)

func main() {
	p, err := privacy.CreatePassword("12345678")
	if err != nil {
		log.Fatal(err)
	}
	CryptoKey, err := privacy.GetHashCryptoKeyFromPassword(p)
	if err != nil {
		log.Fatal(err)
	}
	privacy.Keystore.InserKey(CryptoKey)
	text, err := privacy.EncryptString("some text")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)
	text, err = privacy.DecryptString(text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)
}
