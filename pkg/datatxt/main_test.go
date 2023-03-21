package datatxt

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/piliphulko/restapi-txt/pkg/privacy"
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

var (
	AddValue         func(AllTypes) error
	DelValue         func(AllTypes) error
	addFile, delFile *os.File
)

func TestMain(m *testing.M) {
	os.RemoveAll(locationTest)
	changeToTestPath(DetailTypes, typeTest, "datatest/TotalTest")
	UseCipherInType(typeTest)
	CheckEndWarehousingData(typeTest)
	privacy.Keystore.InserKey("key12345")
	addFile, delFile, _, err := DataWarehouseDeployment(typeTest)
	if err != nil {
		panic(err)
	}
	testData, err := GetMainSlice(typeTest)
	if err != nil {
		panic(err)
	}
	AddValue = NewGetAddfunc(addFile, testData, typeTest)
	DelValue = NewGetDelfunc(delFile, testData, typeTest)

	mR := m.Run()
	addFile.Close()
	delFile.Close()
	os.RemoveAll(locationTest)
	os.Exit(mR)
}

func Test_TotalTest(t *testing.T) {

	assert.Nil(t, AddValue(testType{tInt: 10, tString: "abc", tBool: true, tFloat: 25.25}))
	assert.ErrorContains(t, AddValue(testType{tInt: 10, tString: "abc", tBool: true, tFloat: 25.25}), ErrValueExist.Error())
	assert.Nil(t, DelValue(testType{tInt: 10, tString: "abc", tBool: true, tFloat: 25.25}))
	assert.ErrorContains(t, DelValue(testType{tInt: 10, tString: "abc", tBool: true, tFloat: 25.25}), ErrNoSuchValue.Error())

	testFind := testType{tInt: 11, tString: "abc", tBool: false, tFloat: 25.25}
	assert.Nil(t, AddValue(testFind))
}

var bendE *error

func Benchmark_all(b *testing.B) {
	var i int
	for i = 0; i != 10000; i++ {
		require.Nil(b, AddValue(testType{tInt: i, tString: "abc", tBool: true, tFloat: 25.25}))
		require.Nil(b, DelValue(testType{tInt: i, tString: "abc", tBool: true, tFloat: 25.25}))
	}
	addFile.Close()
	delFile.Close()
	err := CheckEndWarehousingData(typeTest)
	bendE = &err
}

//Benchmark_all-8   	       1	2794215200 ns/op	1648823224 B/op	  329862 allocs/op
// cipher in read
//Benchmark_all-8   	       1	2983876800 ns/op	1648828000 B/op	  329927 allocs/op // new 7%
// cipher in write
//Benchmark_all-8   	       1	2852895400 ns/op	1648824872 B/op	  329877 allocs/op // new 2%
// cipher in read + write
//Benchmark_all-8   	       1	2870447200 ns/op	1648821768 B/op	  329859 allocs/op // new 3%???
//implementation cipher:
//Benchmark_all-8   	       1	3008773400 ns/op	1671801944 B/op	  690103 allocs/op --
