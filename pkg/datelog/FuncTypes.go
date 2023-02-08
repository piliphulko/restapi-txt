package datelog

import (
	"fmt"
	"io"
)

func Unique(slice []AllTypes) []AllTypes {
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
	return Unique(append(slice[:n], slice[n+1:]...))
}

func cleanAddDel(types, typesAdd, typesDel []AllTypes) []AllTypes {
	type allTypesCOUNT struct {
		AllTypes
		count int
	}
	typesCOUNT := []allTypesCOUNT{}
	returnTypes := []AllTypes{}
	for _, v := range Unique(append(types, typesAdd...)) {
		typesCOUNT = append(typesCOUNT, allTypesCOUNT{v, 0})
	}
	for _, v := range typesCOUNT {
		fmt.Println(v, v.count)
		for _, vA := range append(types, typesAdd...) {
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

func WriteType(fileName io.Writer, slice []AllTypes, vType AllTypes) ([]AllTypes, error) {
	if _, err := fmt.Fprintln(fileName, vType); err != nil {
		return slice, err
	}
	slice = append(slice, vType)
	return slice, nil
}
