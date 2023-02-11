package datelog

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func logFataIF(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func errInfoType(err error, typeData int) error {
	return fmt.Errorf("%w | type: %s", err, DetailTypes[typeData].NameType)
}

func GetDataTypesFromFile(fileName string, typeData int) ([]AllTypes, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, os.FileMode(0755))
	if err != nil {
		return nil, errInfoType(err, typeData)
	}
	scanner := bufio.NewScanner(file)
	slice := make([]AllTypes, 0)
	for scanner.Scan() {
		sliceType, err := DetailTypes[typeData].ScanType(scanner.Text())
		if err != nil {
			return nil, errInfoType(err, typeData)
		}
		slice = append(slice, sliceType)
	}
	if scanner.Err() != nil {
		return nil, errInfoType(err, typeData)
	}
	if file.Close() != nil {
		return nil, errInfoType(err, typeData)
	}
	return slice, errInfoType(err, typeData)
}

func DataWarehouseDeployment(typeData int) (*os.File, *os.File, func(), error) {
	var (
		fm           = os.FileMode(0755)
		settingsFile = os.O_CREATE | os.O_RDWR | os.O_APPEND
	)
	if err := deploymentMkdir(DetailTypes, typeData); err != nil {
		return nil, nil, nil, errInfoType(err, typeData)
	}
	fAdd, err := os.OpenFile(DetailTypes[typeData].LocationAddFile, settingsFile, fm)
	if err != nil {
		return nil, nil, nil, errInfoType(err, typeData)
	}
	fDel, err := os.OpenFile(DetailTypes[typeData].LocationDelFile, settingsFile, fm)
	if err != nil {
		return nil, nil, nil, errInfoType(err, typeData)
	}
	closeData := func() {
		logFataIF(fAdd.Close())
		logFataIF(fDel.Close())
		WarehousingData(typeData)
	}
	return fAdd, fDel, closeData, nil
}

func WarehousingData(typeData int) {
	typesMain, err := GetDataTypesFromFile(DetailTypes[typeData].LocationMainFile, typeData)
	logFataIF(err)
	typesAdd, err := GetDataTypesFromFile(DetailTypes[typeData].LocationAddFile, typeData)
	logFataIF(err)
	typesDel, err := GetDataTypesFromFile(DetailTypes[typeData].LocationDelFile, typeData)
	logFataIF(err)
	logFataIF(os.Rename(DetailTypes[typeData].LocationMainFile, DetailTypes[typeData].LocationStockMainFile))
	fMain, err := os.OpenFile(DetailTypes[typeData].LocationMainFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0755))
	logFataIF(err)
	for _, v := range cleanAddDel(typesMain, typesAdd, typesDel) {
		fmt.Fprintln(fMain, v)
	}
	logFataIF(fMain.Close())
	os.Remove(DetailTypes[typeData].LocationAddFile)
	os.Remove(DetailTypes[typeData].LocationDelFile)
	os.Remove(DetailTypes[typeData].LocationStockMainFile)
}

func deploymentMkdir(m map[int]typeDetail, typeData int) error {
	mkdirAll := func(path string) error {
		if err := os.MkdirAll(path, 0750); err != nil && !os.IsExist(err) {
			return err
		}
		return nil
	}
	var slice []string
	slice = append(slice, filepath.Dir(m[typeData].LocationMainFile),
		filepath.Dir(m[typeData].LocationAddFile),
		filepath.Dir(m[typeData].LocationDelFile),
		filepath.Dir(m[typeData].LocationStockMainFile))
	for _, v := range slice {
		if err := mkdirAll(v); err != nil {
			return errInfoType(err, typeData)
		}
	}
	return nil
}

func CheckEndWarehousingData(typeData int) error {
	_, err0 := os.Stat(DetailTypes[typeData].LocationStockMainFile)
	_, err1 := os.Stat(DetailTypes[typeData].LocationAddFile)
	_, err2 := os.Stat(DetailTypes[typeData].LocationDelFile)
	switch {
	case (errors.Is(err0, os.ErrNotExist)) && !(errors.Is(err1, os.ErrNotExist)) && !(errors.Is(err2, os.ErrNotExist)):
		fmt.Printf("did not happen: %s\n", DetailTypes[typeData].NameType)
		WarehousingData(typeData)
	case !(errors.Is(err0, os.ErrNotExist)) && !(errors.Is(err1, os.ErrNotExist)) && !(errors.Is(err2, os.ErrNotExist)):
		fmt.Printf("there was an incorrect completion of writing to the main file: %s\n", DetailTypes[typeData].NameType)
		if err := os.Remove(DetailTypes[typeData].LocationMainFile); err != nil {
			return errInfoType(err, typeData)
		}
		if err := os.Rename(DetailTypes[typeData].LocationStockMainFile, DetailTypes[typeData].LocationMainFile); err != nil {
			return errInfoType(err, typeData)
		}
		WarehousingData(typeData)
	case (!(errors.Is(err1, os.ErrNotExist)) && (errors.Is(err2, os.ErrNotExist))) ||
		((errors.Is(err1, os.ErrNotExist)) && !(errors.Is(err2, os.ErrNotExist))):
		fmt.Printf("some file is missing (Add or Del): %s\n", DetailTypes[typeData].NameType)
		return errors.New("error file Add or Del")
	default:
		return nil
	}
	return nil
}
