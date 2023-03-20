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

type MutexAllTypes struct {
	rwm sync.RWMutex
	all []AllTypes
}
