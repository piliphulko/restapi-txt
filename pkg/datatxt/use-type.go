package datatxt

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/piliphulko/restapi-txt/pkg/privacy"
)

func UniqueOld(slice []AllTypes) []AllTypes {
	var n int
	for n1 := 0; n1 < len(slice); n1++ {
		for n2 := len(slice) - 1; n2 > 0; n2-- {
			if n1 == n2 {
			} else if slice[n1] == slice[n2] {
				n = n1
				goto finish
			}
		}
	}
	return slice
finish:
	return UniqueOld(append(slice[:n], slice[n+1:]...))
}

// Unique returns all unique values slice

func Unique(slice []AllTypes) []AllTypes {
	mapSlice := map[AllTypes]bool{}
	returnSlice := []AllTypes{}
	for _, v := range slice {
		mapSlice[v] = true
	}
	for v := range mapSlice {
		returnSlice = append(returnSlice, v)
	}
	return returnSlice
}

// cleanAddDel takes three slices:
//  1. main data
//  2. added data
//  3. delete data
// the function implements adding and deleting data and returns a slice of all data that is created and not deleted

func cleanAddDel(typesMain, typesAdd, typesDel []AllTypes) []AllTypes {
	type allTypesCOUNT struct {
		AllTypes
		count int
	}
	var (
		returnTypes = []AllTypes{}
		typesCOUNT  = make(chan allTypesCOUNT)
	)
	go func() {
		for _, v := range Unique(append(typesMain, typesAdd...)) {
			typesCOUNT <- allTypesCOUNT{v, 0}
		}
		close(typesCOUNT)
	}()
	for v := range typesCOUNT {
		for _, vA := range append(typesMain, typesAdd...) {
			if v.AllTypes == vA {
				v.count += 1
			}
		}
		for _, vD := range typesDel {
			if v.AllTypes == vD {
				v.count -= 1
			}
		}
		if v.count >= 1 {
			returnTypes = append(returnTypes, v.AllTypes)
		}
	}
	return returnTypes
}

func WriteFromBuffer(w io.Writer, buf bytes.Buffer, typeData int) error {
	if DetailTypes[typeData].Cipher.Cipher {
		text, err := privacy.EncryptString(buf.String())
		if err != nil {
			return err
		}
		if _, err := fmt.Fprintln(w, text); err != nil {
			return err
		}
	} else {
		if _, err := fmt.Fprintln(w, buf.String()); err != nil {
			return err
		}
	}
	return nil
}

func GetMainSlice(typeData int) (*MutexAllTypes, error) {
	slice, err := GetDataTypesFromFile(DetailTypes[typeData].LocationMainFile, typeData)
	if err != nil {
		return nil, err
	}
	var slicePlusMutex = MutexAllTypes{all: slice}
	return &slicePlusMutex, nil
}

func (mat *MutexAllTypes) FindValue(value AllTypes) bool {
	for _, v := range mat.all {
		if value == v {
			return true
		}
	}
	return false
}

func GetAddfunc(addFile *os.File, dataMain *MutexAllTypes) func(AllTypes) error {
	fn := func(addValue AllTypes) error {
		if dataMain.FindValue(addValue) {
			return ErrValueExist
		}
		dataMain.rwm.Lock()
		defer dataMain.rwm.Unlock()
		if _, err := fmt.Fprintln(addFile, addValue); err != nil {
			return err
		}
		dataMain.all = append(dataMain.all, addValue)
		return nil
	}
	return fn
}

func cutDel(value AllTypes, slice []AllTypes) ([]AllTypes, error) {
	for i, v := range slice {
		if value == v {
			return append(slice[:i], slice[i+1:]...), nil
		}
	}
	return slice, ErrNoSuchValue
}

func GetDelfunc(delFile *os.File, dataMain *MutexAllTypes) func(AllTypes) error {
	fn := func(delValue AllTypes) error {
		dataMain.rwm.Lock()
		defer dataMain.rwm.Unlock()
		newSlice, err := cutDel(delValue, dataMain.all)
		if err != nil {
			return err // [ErrNoSuchValue]
		}
		if _, err := fmt.Fprintln(delFile, delValue); err != nil {
			return err
		}
		dataMain.all = newSlice
		return nil
	}
	return fn
}
