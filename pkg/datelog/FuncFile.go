package datelog

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func ReadTypeFromAFile(name string, typeDate int) ([]AllTypes, error) {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_RDONLY|os.O_APPEND, os.FileMode(0755))
	if err != nil {
		return nil, err
	}
	defer func() {
		if file.Close() != nil {
			log.Fatal(err)
		}
	}()
	var (
		scanner    = bufio.NewScanner(file)
		sliceTyper = make([]AllTypes, 0)
	)
	switch typeDate {
	case TypeUser:
		var (
			id   int
			name string
		)
		for scanner.Scan() {
			if _, err := fmt.Sscanf(scanner.Text(), sampleTypes[TypeUser], &id, &name); err != nil {
				return nil, err
			}
			sliceTyper = append(sliceTyper, User{id, name})
		}
	default:
		return nil, errors.New("error scan type - non")
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return sliceTyper, err
}

func FileAddDeleteDateType(fileName string, typeUser int) (*os.File, *os.File, func(), error) {
	var (
		returnFunc func()
		fatalErr   = func(err error) {
			if err != nil {
				log.Fatal(err)
			}
		}
	)
	switch typeUser {
	case TypeUser:
		//fAdd, err := os.CreateTemp("", "tUserAdd_*.log")
		//fDelete, err = os.CreateTemp("", "tUserDelete_*.log")
		fAdd, err := os.OpenFile("tUserAdd.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.FileMode(0755))
		if err != nil {
			return nil, nil, nil, err
		}
		fDelete, err := os.OpenFile("tUserDelete.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.FileMode(0755))
		if err != nil {
			return nil, nil, nil, err
		}
		//if os.WriteFile("locationUserFiles.txt",
		//	[]byte(fmt.Sprintf("User type\nAdd: %s\nDel: %s", fAdd.Name(), fDelete.Name())), os.FileMode(0755)) != nil {
		//	log.Fatal(err)
		//}
		returnFunc = func() {
			fatalErr(fAdd.Close())
			fatalErr(fDelete.Close())
			usersMain, err := ReadTypeFromAFile(fileName, TypeUser)
			fatalErr(err)
			usersNew, err := ReadTypeFromAFile(fAdd.Name(), TypeUser)
			fatalErr(err)
			usersDelete, err := ReadTypeFromAFile(fDelete.Name(), TypeUser)
			fatalErr(err)
			os.Remove(fileName)
			file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0755))
			fatalErr(err)
			for _, v := range cleanAddDel(usersMain, usersNew, usersDelete) {
				fmt.Fprintln(file, v)
			}
			fatalErr(file.Close())
			os.Remove(fAdd.Name())
			os.Remove(fDelete.Name())
		}
		return fAdd, fDelete, returnFunc, nil
	default:
		return nil, nil, nil, errors.New("no such type")
	}
}
