package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/piliphulko/practiceGo/pkg/datalog"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	var logFatalIF = func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	logFatalIF(datalog.CheckEndWarehousingData(datalog.TypeUser))
	userAdd, userDel, userClose, err := datalog.DataWarehouseDeployment(datalog.TypeUser)
	logFatalIF(err)
	defer userClose()

	rand.Seed(time.Now().UnixNano())
	buf := bytes.NewBuffer(make([]byte, 0, 5))

	for n := 0; n != 10000; n++ {
		buf.Reset()
		for i := 0; i != 4; i++ {
			buf.WriteByte(letterBytes[rand.Intn(len(letterBytes))])
		}
		_, err = fmt.Fprintln(userAdd, datalog.User{Id: rand.Intn(10), Name: string(buf.Bytes())})
		logFatalIF(err)
	}
	for n := 0; n != 10000; n++ {
		buf.Reset()
		for i := 0; i != 4; i++ {
			buf.WriteByte(letterBytes[rand.Intn(len(letterBytes))])
		}
		_, err = fmt.Fprintln(userDel, datalog.User{Id: rand.Intn(10), Name: string(buf.Bytes())})
		logFatalIF(err)
	}
}
