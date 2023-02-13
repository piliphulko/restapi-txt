package datalog

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//	func testFillingFile() {
//		DetailTypes[typeTest]
//	}
var locationTest = "datatest"

func TestMain(m *testing.M) {
	mR := m.Run()
	os.RemoveAll(locationTest)
	os.Exit(mR)
}
func Test_funcdeploymentMkdir(t *testing.T) {
	assert.Nil(t, deploymentMkdir(DetailTypes, typeTest))
	assert.True(t, os.IsExist(os.Mkdir(filepath.Dir(DetailTypes[typeTest].LocationAddFile), 0750)))
	assert.True(t, os.IsExist(os.Mkdir(filepath.Dir(DetailTypes[typeTest].LocationDelFile), 0750)))
	assert.True(t, os.IsExist(os.Mkdir(filepath.Dir(DetailTypes[typeTest].LocationMainFile), 0750)))
}

func Test_funcGetDataTypesFromFile(t *testing.T) { //FAIL
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
