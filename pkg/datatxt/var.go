package datatxt

import (
	"errors"
	"fmt"

	"github.com/piliphulko/restapi-txt/pkg/privacy"
)

var (
	ErrNoSuchValue = errors.New("no such value")
	ErrValueExist  = errors.New("Error value exists")
)

const (
	typeTest = iota
	TypeUser
)

// DetailTypes map where spelled out the main details of the implementation of types in the package

var DetailTypes = map[int]typeDetail{
	typeTest: {
		NameType:              "testType",
		SampleFMT:             "tInt: %d tString: %s tBool: %t tFloat: %g",
		LocationMainFile:      "datatest/main.txt",
		LocationAddFile:       "datatest/add.txt",
		LocationDelFile:       "datatest/del.txt",
		LocationStockMainFile: "datatest/stockMain.txt",
		ScanType: func(s string) (AllTypes, error) {
			var (
				tIntV    int
				tStringV string
				tBoolV   bool
				tFloatV  float64
			)
			if _, err := fmt.Sscanf(s, "tInt: %d tString: %s tBool: %t tFloat: %g", &tIntV, &tStringV, &tBoolV, &tFloatV); err != nil {
				return nil, err
			}
			return testType{tIntV, tStringV, tBoolV, tFloatV}, nil
		},
		Cipher: privacy.CipherUse{Cipher: true},
	},
	TypeUser: {
		NameType:              "User",
		SampleFMT:             "Login: %s Passwort: %s",
		LocationMainFile:      "data/User/mainUser.txt",
		LocationAddFile:       "data/User/addUser.txt",
		LocationDelFile:       "data/User/delUser.txt",
		LocationStockMainFile: "data/User/stockMainUser.txt",
		ScanType: func(s string) (AllTypes, error) {
			var (
				login    string
				passwort string
			)
			if _, err := fmt.Sscanf(s, "Login: %s Passwort: %s", &login, &passwort); err != nil {
				return nil, err
			}
			return User{login, passwort}, nil
		},
	},
}
