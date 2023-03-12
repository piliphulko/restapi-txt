package httpout

import (
	"log"

	"github.com/piliphulko/practiceGo/pkg/datalog"
)

var (
	users   = &datalog.MutexAllTypes{}
	addUser func(datalog.AllTypes) error
	delUser func(datalog.AllTypes) error
)

func init() {
	err := datalog.CheckEndWarehousingData(datalog.TypeUser)
	if err != nil {
		log.Fatal(err)
	}
	addFileUser, delFileUser, _, err := datalog.DataWarehouseDeployment(datalog.TypeUser)
	if err != nil {
		log.Fatal(err)
	}
	users, err = datalog.GetMainSlice(datalog.TypeUser)
	if err != nil {
		log.Fatal(err)
	}
	addUser = datalog.GetAddfunc(addFileUser, users)
	delUser = datalog.GetDelfunc(delFileUser, users)
}
