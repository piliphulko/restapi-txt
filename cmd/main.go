package main

import (
	"fmt"
	"log"

	_ "github.com/piliphulko/practiceGo/pkg/datalog"
	"github.com/piliphulko/practiceGo/pkg/privacy"
)

func main() {
	privacy.Keystore.InserKey("0123456789abcdem")
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
