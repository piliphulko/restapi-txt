package datelog

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetDataTypesFromFile(fileName string, typeData int) ([]AllTypes, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, os.FileMode(0755))
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	slice := make([]AllTypes, 0)
	for scanner.Scan() {
		sliceType, err := DetailTypes[typeData].ScanType(scanner.Text())
		if err != nil {
			return nil, err
		}
		slice = append(slice, sliceType)
	}
	if scanner.Err() != nil {
		return nil, err
	}
	if file.Close() != nil {
		return nil, err
	}
	return slice, nil
}

func DataWarehouseDeployment(typeData int) (*os.File, *os.File, func(), error) {
	var (
		fm           = os.FileMode(0755)
		settingsFile = os.O_CREATE | os.O_RDWR | os.O_APPEND
		logFataIF    = func(err error) {
			if err != nil {
				log.Fatal(err)
			}
		}
	)

	err := os.Mkdir("data", fm)
	if err != nil && !os.IsExist(err) {
		return nil, nil, nil, err
	}
	err = os.Mkdir(DetailTypes[typeData].DirectoryFiles, fm)
	if err != nil && !os.IsExist(err) {
		return nil, nil, nil, err
	}
	fAdd, err := os.OpenFile(DetailTypes[typeData].LocationAddFile, settingsFile, fm)
	if err != nil {
		return nil, nil, nil, err
	}
	fDel, err := os.OpenFile(DetailTypes[typeData].LocationDelFile, settingsFile, fm)
	if err != nil {
		return nil, nil, nil, err
	}
	closeData := func() {
		logFataIF(fAdd.Close())
		logFataIF(fDel.Close())
		WarehousingData(typeData)
	}
	return fAdd, fDel, closeData, err
}

func WarehousingData(typeData int) {
	var logFataIF = func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
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
