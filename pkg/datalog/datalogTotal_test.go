package datalog

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func MainTest(m *testing.M) {
	locationTest := "datatest"
	mRun := m.Run()
	os.RemoveAll(locationTest)
	os.Exit(mRun)
}

func Test_ToatalPackage(t *testing.T) {
	assert.Nil(t, CheckEndWarehousingData(typeTest))
}
