package datelog

import "fmt"

type AllTypes interface {
	String() string
}

type User struct {
	Id   int
	Name string
}

func (u User) String() string {
	return fmt.Sprintf(DetailTypes[TypeUser].SampleFMT, u.Id, u.Name)
}

const (
	_ = iota
	TypeUser
)

type typeDetail struct {
	SampleFMT             string
	LocationMainFile      string
	LocationAddFile       string
	LocationDelFile       string
	LocationStockMainFile string
	ScanType              func(string) (AllTypes, error)
}

var DetailTypes = map[int]typeDetail{
	TypeUser: {
		SampleFMT:             "Id: %d Name: %s",
		LocationMainFile:      "data/User/mainUser.log",
		LocationAddFile:       "data/User/addUser.log",
		LocationDelFile:       "data/User/delUser.log",
		LocationStockMainFile: "data/User/stockMainUser.log",
		ScanType: func(s string) (AllTypes, error) {
			var (
				id   int
				name string
			)
			if _, err := fmt.Sscanf(s, "Id: %d Name: %s", &id, &name); err != nil {
				return nil, err
			}
			return User{id, name}, nil
		},
	},
}
