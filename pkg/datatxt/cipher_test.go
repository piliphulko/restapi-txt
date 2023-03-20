package datatxt_test

import (
	"os"
	"testing"

	"github.com/piliphulko/restapi-txt/pkg/datatxt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var locationTest = "datatest"

var (
	AddValue         func(datatxt.AllTypes) error
	DelValue         func(AllTypes) error
	addFile, delFile *os.File
)

func TestMain(m *testing.M) {
	os.RemoveAll(locationTest)
	changeToTestPath(DetailTypes, typeTest, "datatest/TotalTest")
	CheckEndWarehousingData(typeTest)
	addFile, delFile, _, err := DataWarehouseDeployment(typeTest)
	if err != nil {
		panic(err)
	}
	testData, err := GetMainSlice(typeTest)
	if err != nil {
		panic(err)
	}
	AddValue = GetAddfunc(addFile, testData)
	DelValue = GetDelfunc(delFile, testData)

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
