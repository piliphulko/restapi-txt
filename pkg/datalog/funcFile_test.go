package datalog

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var locationTest = "datatest"

func TestMain(m *testing.M) {
	mR := m.Run()
	os.RemoveAll(locationTest)
	os.Exit(mR)
}
func Test_deploymentMkdir(t *testing.T) {
	assert.Nil(t, deploymentMkdir(DetailTypes, typeTest))
	assert.True(t, os.IsExist(os.Mkdir(filepath.Dir(DetailTypes[typeTest].LocationAddFile), 0750)))
	assert.True(t, os.IsExist(os.Mkdir(filepath.Dir(DetailTypes[typeTest].LocationDelFile), 0750)))
	assert.True(t, os.IsExist(os.Mkdir(filepath.Dir(DetailTypes[typeTest].LocationMainFile), 0750)))
}

func Test_GetDataTypesFromFile(t *testing.T) { //FAIL
	os.MkdirAll(locationTest, 0750)
	locationFile := locationTest + "/GetDataTypesFromFile.log"
	file, err := os.OpenFile(locationFile, os.O_CREATE|os.O_RDONLY, os.FileMode(0755))
	require.Nil(t, err)
	var data = []testType{
		{1, "a", true, 1.15},
		{20, "abc", false, -51.777},
		{-30, "MMM", false, 1000.99},
	}
	for i := 0; i != len(data); i++ {
		fmt.Fprintln(file, data[i])
		fmt.Println(data[i], i)
	}
	require.Nil(t, file.Close())
	dataGet, err := GetDataTypesFromFile(locationFile, typeTest)
	require.Nil(t, err)
	for i := 0; i != len(data); i++ {
		assert.Equal(t, data[i], dataGet[i])
	}
}

func Test_DataWarehouseDeployment_Plus_warehousingData(t *testing.T) {
	assert.Nil(t, os.MkdirAll(filepath.Dir(DetailTypes[typeTest].LocationMainFile), 0750))
	start, err := GetDataTypesFromFile(DetailTypes[typeTest].LocationMainFile, typeTest)
	require.Nil(t, err)
	startLen := len(start)
	addFile, delFile, sortFile, err := DataWarehouseDeployment(typeTest)
	require.Nil(t, err)
	fmt.Fprintln(addFile, testType{10, "tt", true, 66.35})
	fmt.Fprintln(delFile, testType{10, "tt", true, 66.35})
	fmt.Fprintln(addFile, testType{10, "tt", false, 66.35})
	sortFile()
	finish, err := GetDataTypesFromFile(DetailTypes[typeTest].LocationMainFile, typeTest)
	require.Nil(t, err)
	finishLen := len(finish)
	require.Equal(t, startLen+1, finishLen)
}

func Test_CheckEndWarehousingData(t *testing.T) {
	addFile, delFile, _, err := DataWarehouseDeployment(typeTest)
	require.Nil(t, err)
	fmt.Fprintln(addFile, testType{25, "plpG", true, 750.85})
	fmt.Fprintln(delFile, testType{25, "plpG", true, 750.85})
	fmt.Fprintln(addFile, testType{25, "plpG", false, 750.85})
	require.Nil(t, addFile.Close())
	require.Nil(t, delFile.Close())
	start, err := GetDataTypesFromFile(DetailTypes[typeTest].LocationMainFile, typeTest)
	require.Nil(t, err)
	startLen := len(start)

	require.Nil(t, CheckEndWarehousingData(typeTest))

	finish, err := GetDataTypesFromFile(DetailTypes[typeTest].LocationMainFile, typeTest)
	require.Nil(t, err)
	finishLen := len(finish)
	require.Equal(t, startLen+1, finishLen)

	addFileNew, delFileNew, _, err := DataWarehouseDeployment(typeTest)
	require.Nil(t, err)
	fmt.Fprintln(delFileNew, testType{25, "plpG", false, 750.85})
	require.Nil(t, addFileNew.Close())
	require.Nil(t, delFileNew.Close())

	require.Nil(t, CheckEndWarehousingData(typeTest))

	finish1, err := GetDataTypesFromFile(DetailTypes[typeTest].LocationMainFile, typeTest)
	require.Nil(t, err)
	finishLen = len(finish1)
	require.Equal(t, startLen, finishLen)

	require.Nil(t, os.WriteFile(DetailTypes[typeTest].LocationStockMainFile, []byte("tInt: 0 tString: a tBool: false tFloat: 0.10"), 0750))
	require.Nil(t, os.WriteFile(DetailTypes[typeTest].LocationAddFile, []byte("tInt: 0 tString: a tBool: false tFloat: 0.10"), 0750))
	require.Nil(t, os.WriteFile(DetailTypes[typeTest].LocationDelFile, []byte("tInt: 0 tString: a tBool: false tFloat: 0.10"), 0750))

	require.Nil(t, CheckEndWarehousingData(typeTest))

	finish2, err := GetDataTypesFromFile(DetailTypes[typeTest].LocationMainFile, typeTest)
	require.Nil(t, err)
	finishLen = len(finish2)
	require.Equal(t, 1, finishLen)
}
