package main

import (
	"fmt"
	"log"

	"github.com/piliphulko/practiceGo/pkg/datelog"
)

func main() {
	var (
		mainFileUser = "file.log"
		//userAdd = []datelog.User{}
		//userDel = []datelog.User{}
		fileUserAdd, fileUserDel, endUserDate, err = datelog.FileAddDeleteDateType(mainFileUser, datelog.TypeUser)
	)
	if err != nil {
		log.Fatal(err)
	}
	defer endUserDate()

	//userAdd = append(userAdd, datelog.User{5, "lip"})
	//userDel = append(userDel, datelog.User{5, "lip"})
	fmt.Fprintln(fileUserAdd, datelog.User{5, "lip"})
	fmt.Fprintln(fileUserDel, datelog.User{5, "lip"})
	fmt.Fprintln(fileUserAdd, datelog.User{5, "lip"})

	//fileUserAdd.Write([]byte("do you work?"))
	//time.Sleep(1 * time.Minute)
}
