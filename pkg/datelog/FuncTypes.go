package datelog

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
