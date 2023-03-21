package datatxt

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/piliphulko/restapi-txt/pkg/privacy"
)

func errInfoType(err error, typeData int) error {
	return fmt.Errorf("%w | type: %s", err, DetailTypes[typeData].NameType)
}

func logFataIF(err error, typeData int) {
	if err != nil {
		log.Fatal(fmt.Errorf("%w | type: %s", err, DetailTypes[typeData].NameType))
	}
}

// GetDataFromFile takes two variables: file name and data type written to this file
// returns a slice of the read data and an error
//
// If there is no file, then it is created. after reading the data, the file is closed inside the function

func GetDataTypesFromFile(fileName string, typeData int) ([]AllTypes, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, os.FileMode(0755))
	if err != nil {
		return nil, errInfoType(err, typeData)
	}
	var (
		scanner  = bufio.NewScanner(file)
		slice    = make([]AllTypes, 0)
		readText = func(text string) (string, error) { // -7% fast
			return text, nil
		}
	)
	if DetailTypes[typeData].Cipher.Cipher {
		readText = func(text string) (string, error) {
			return privacy.DecryptString(text)
		}
	}
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		text, err := readText(text)
		fmt.Println(text)
		if err != nil {
			return nil, errInfoType(err, typeData)
		}
		sliceType, err := DetailTypes[typeData].ScanType(text)
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
	return slice, nil
}

// DataWarehouseDeployment takes a data type and returns:
//  1. file to add data
//  2. file to delete data
//  3. the function that closes files and writes all added data to the main storage
//  4. errors

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
		logFataIF(fAdd.Close(), typeData)
		logFataIF(fDel.Close(), typeData)
		warehousingData(typeData)
	}
	return fAdd, fDel, closeData, nil
}

// warehousingData a function that collects data from files removed and added to the main storage

func warehousingData(typeData int) {
	typesMain, err := GetDataTypesFromFile(DetailTypes[typeData].LocationMainFile, typeData)
	logFataIF(err, typeData)
	typesAdd, err := GetDataTypesFromFile(DetailTypes[typeData].LocationAddFile, typeData)
	logFataIF(err, typeData)
	typesDel, err := GetDataTypesFromFile(DetailTypes[typeData].LocationDelFile, typeData)
	logFataIF(err, typeData)
	logFataIF(os.Rename(DetailTypes[typeData].LocationMainFile, DetailTypes[typeData].LocationStockMainFile), typeData)
	fMain, err := os.OpenFile(DetailTypes[typeData].LocationMainFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0755))
	logFataIF(err, typeData)
	buf := bytes.Buffer{}
	for _, v := range cleanAddDel(typesMain, typesAdd, typesDel) {
		fmt.Fprint(&buf, v)
		logFataIF(WriteFromBuffer(fMain, buf, typeData), typeData)
		buf.Reset()
	}
	logFataIF(fMain.Close(), typeData)
	os.Remove(DetailTypes[typeData].LocationAddFile)
	os.Remove(DetailTypes[typeData].LocationDelFile)
	os.Remove(DetailTypes[typeData].LocationStockMainFile)
}

// deploymentMkdir creates the necessary directory for files

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

// CheckEndWarehousingData check the correct completion of data processing in the main storage
// If data processing has not occurred and it can be completed, then the function performs processing

func CheckEndWarehousingData(typeData int) error {
	_, err0 := os.Stat(DetailTypes[typeData].LocationStockMainFile)
	_, err1 := os.Stat(DetailTypes[typeData].LocationAddFile)
	_, err2 := os.Stat(DetailTypes[typeData].LocationDelFile)
	switch {
	case (errors.Is(err0, os.ErrNotExist)) && !(errors.Is(err1, os.ErrNotExist)) && !(errors.Is(err2, os.ErrNotExist)):
		fmt.Printf("type sorting to file did not happen, attempt to fix: %s\n", DetailTypes[typeData].NameType)
		warehousingData(typeData)
		fmt.Printf("type sorting to file successfully completed: %s\n", DetailTypes[typeData].NameType)
	case !(errors.Is(err0, os.ErrNotExist)) && !(errors.Is(err1, os.ErrNotExist)) && !(errors.Is(err2, os.ErrNotExist)):
		fmt.Printf("there was an incorrect completion of writing to the main file: %s\n", DetailTypes[typeData].NameType)
		if err := os.Remove(DetailTypes[typeData].LocationMainFile); err != nil {
			return errInfoType(err, typeData)
		}
		if err := os.Rename(DetailTypes[typeData].LocationStockMainFile, DetailTypes[typeData].LocationMainFile); err != nil {
			return errInfoType(err, typeData)
		}
		warehousingData(typeData)
	case (!(errors.Is(err1, os.ErrNotExist)) && (errors.Is(err2, os.ErrNotExist))) ||
		((errors.Is(err1, os.ErrNotExist)) && !(errors.Is(err2, os.ErrNotExist))):
		fmt.Printf("some file is missing (Add or Del): %s\n", DetailTypes[typeData].NameType)
		return errors.New("error file Add or Del")
	default:
		return nil
	}
	return nil
}
