package httpout

import (
	"log"

	"github.com/piliphulko/restapi-txt/pkg/datatxt"
	"github.com/piliphulko/restapi-txt/pkg/privacy"
)

var (
	users   = &datatxt.MutexAllTypes{}
	addUser func(datatxt.AllTypes) error
	delUser func(datatxt.AllTypes) error
)

func init() {
	var key = "abc12345rdtwer67"
	privacy.Keystore.InserKey(key)

	datatxt.UseCipherInType(datatxt.TypeUser)
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
	addUser = datatxt.NewGetAddfunc(addFileUser, users, datatxt.TypeUser)
	delUser = datatxt.NewGetDelfunc(delFileUser, users, datatxt.TypeUser)
}
