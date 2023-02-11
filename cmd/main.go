package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/piliphulko/practiceGo/pkg/datelog"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	var logFatalIF = func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	logFatalIF(datelog.CheckEndWarehousingData(datelog.TypeUser))
	userAdd, userDel, userClose, err := datelog.DataWarehouseDeployment(datelog.TypeUser)
	logFatalIF(err)
	defer userClose()

	rand.Seed(time.Now().UnixNano())
	buf := bytes.NewBuffer(make([]byte, 0, 5))

	for n := 0; n != 10000; n++ {
		buf.Reset()
		for i := 0; i != 4; i++ {
			buf.WriteByte(letterBytes[rand.Intn(len(letterBytes))])
		}
		_, err = fmt.Fprintln(userAdd, datelog.User{Id: rand.Intn(10), Name: string(buf.Bytes())})
		logFatalIF(err)
	}
	for n := 0; n != 10000; n++ {
		buf.Reset()
		for i := 0; i != 4; i++ {
			buf.WriteByte(letterBytes[rand.Intn(len(letterBytes))])
		}
		_, err = fmt.Fprintln(userDel, datelog.User{Id: rand.Intn(10), Name: string(buf.Bytes())})
		logFatalIF(err)
	}
}
