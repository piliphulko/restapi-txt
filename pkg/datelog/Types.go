package datelog

import "fmt"

type User struct {
	Id   int
	Name string
}

const (
	_ = iota
	TypeUser
)

var sampleTypes = map[int]string{
	TypeUser: "Id: %d Name: %s",
}

func (u User) String() string {
	return fmt.Sprintf(sampleTypes[TypeUser], u.Id, u.Name)
}

type AllTypes interface {
	String() string
}
