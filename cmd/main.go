package main

import (
	"fmt"
	"log"

	"github.com/piliphulko/practiceGo/pkg/datelog"
)

func main() {
	var logFatalIF = func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	userAdd, userDel, userClose, err := datelog.DataWarehouseDeployment(datelog.TypeUser)
	logFatalIF(err)
	defer userClose()
	_, err = fmt.Fprintln(userAdd, datelog.User{Id: 2, Name: "pip"})
	logFatalIF(err)
	_, err = fmt.Fprintln(userDel, datelog.User{Id: 2, Name: "pip"})
	logFatalIF(err)
	_, err = fmt.Fprintln(userAdd, datelog.User{Id: 2, Name: "pip"})
	logFatalIF(err)
}
