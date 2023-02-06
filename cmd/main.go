package main

import (
	"log"

	"github.com/piliphulko/practiceGo/pkg/rememberlog"
)

func main() {
	file, err := rememberlog.CRWfile("new.log")
	if err != nil {
		log.Fatal(err)
	}
}
