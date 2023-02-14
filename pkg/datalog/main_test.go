package datalog

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func changeToTestPath(m map[int]typeDetail, typeData int, newDir string) {
	change := func(oldLocation, newDir string) string {
		s := strings.Split(filepath.Dir(oldLocation), `\`)
		s[0] = newDir
		return strings.Join(s, "/") + "/" + filepath.Base(oldLocation)
	}
	getStruct := m[typeData]
	delete(m, typeData)
	getStruct.LocationMainFile = change(getStruct.LocationMainFile, newDir)
	getStruct.LocationAddFile = change(getStruct.LocationAddFile, newDir)
	getStruct.LocationDelFile = change(getStruct.LocationDelFile, newDir)
	getStruct.LocationStockMainFile = change(getStruct.LocationStockMainFile, newDir)
	m[typeData] = getStruct
}

var locationTest = "datatest"

func TestMain(m *testing.M) {
	mR := m.Run()
	os.RemoveAll(locationTest)
	os.Exit(mR)
}

func Test_TotalTest(t *testing.T) {
	changeToTestPath(DetailTypes, typeTest, "datatest/TotalTest")

	CheckEndWarehousingData(typeTest)
	addFile, delFile, endFn, err := DataWarehouseDeployment(typeTest)
	require.Nil(t, err)
	defer endFn()
	testData, err := GetMainSlice(typeTest)
	require.Nil(t, err)
	AddValue := GetAddfunc(addFile, testData)
	DelValue := GetDelfunc(delFile, testData)

	assert.Nil(t, AddValue(testType{tInt: 10, tString: "abc", tBool: true, tFloat: 25.25}))
	assert.ErrorContains(t, AddValue(testType{tInt: 10, tString: "abc", tBool: true, tFloat: 25.25}), ErrValueExist.Error())
	assert.Nil(t, DelValue(testType{tInt: 10, tString: "abc", tBool: true, tFloat: 25.25}))

	testFind := testType{tInt: 11, tString: "abc", tBool: false, tFloat: 25.25}
	assert.Nil(t, AddValue(testFind))
}
