package datatxt

import (
	"fmt"
	"sync"
)

type AllTypes interface {
	String() string
}
type testType struct {
	tInt    int
	tString string
	tBool   bool
	tFloat  float64
}

type User struct {
	Login    string
	Passwort string
}

func (u User) String() string {
	return fmt.Sprintf(DetailTypes[TypeUser].SampleFMT, u.Login, u.Passwort)
}
func (tt testType) String() string {
	return fmt.Sprintf(DetailTypes[typeTest].SampleFMT, tt.tInt, tt.tString, tt.tBool, tt.tFloat)
}

const (
	typeTest = iota
	TypeUser
)

type typeDetail struct {
	NameType              string
	SampleFMT             string // writer template
	LocationMainFile      string
	LocationAddFile       string
	LocationDelFile       string
	LocationStockMainFile string
	ScanType              func(string) (AllTypes, error) // reader template
}

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

type MutexAllTypes struct {
	rwm sync.RWMutex
	all []AllTypes
}
