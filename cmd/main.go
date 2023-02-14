package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	dl "github.com/piliphulko/practiceGo/pkg/datalog"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	rand.Seed(time.Now().Unix())

	dl.CheckEndWarehousingData(dl.TypeUser)
	addFile, delFile, endFn, err := dl.DataWarehouseDeployment(dl.TypeUser)
	if err != nil {
		log.Fatal(err)
	}
	defer endFn()
	userData, err := dl.GetMainSlice(dl.TypeUser)
	if err != nil {
		log.Fatal(err)
	}
	AddValueUser := dl.GetAddfunc(addFile, userData)
	DelValueUser := dl.GetDelfunc(delFile, userData)

	if AddValueUser(dl.User{Id: 10, Name: "pip"}) != nil {
		log.Fatal(err)
	}
	if DelValueUser(dl.User{Id: 10, Name: "pip"}) != nil {
		log.Fatal(err)
	}
	if AddValueUser(dl.User{Id: 11, Name: "pip"}) != nil {
		log.Fatal(err)
	}
	f, err := userData.FindValue(dl.User{Id: 11, Name: "pip"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}
