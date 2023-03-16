package httpout

import (
	"log"

	"github.com/piliphulko/restapi-txt/pkg/datatxt"
)

var (
	users   = &datatxt.MutexAllTypes{}
	addUser func(datatxt.AllTypes) error
	delUser func(datatxt.AllTypes) error
)

func init() {
	err := datatxt.CheckEndWarehousingData(datatxt.TypeUser)
	if err != nil {
		log.Fatal(err)
	}
	addFileUser, delFileUser, _, err := datatxt.DataWarehouseDeployment(datatxt.TypeUser)
	if err != nil {
		log.Fatal(err)
	}
	users, err = datatxt.GetMainSlice(datatxt.TypeUser)
	if err != nil {
		log.Fatal(err)
	}
	addUser = datatxt.GetAddfunc(addFileUser, users)
	delUser = datatxt.GetDelfunc(delFileUser, users)
}
