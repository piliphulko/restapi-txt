package datalog

import (
	"testing"
)

func Test_Unique_UniqueOld(t *testing.T) {
	var slice = []AllTypes{
		User{Login: "am13", Passwort: "12345678"},
		User{Login: "am134", Passwort: "12345678"},
		User{Login: "am135", Passwort: "12345678"},
		User{Login: "am13", Passwort: "12345678"},
		User{Login: "am13", Passwort: "12345678"},
	}
	if len(Unique(slice)) != len(UniqueOld(slice)) {
		t.Error("UniqueNew non work")
	}
}

var bend *[]AllTypes

func Benchmark_funcUniqueOld(b *testing.B) {
	var slice = []AllTypes{
		User{Login: "am13", Passwort: "12345678"},
		User{Login: "am134", Passwort: "12345678"},
		User{Login: "am135", Passwort: "12345678"},
		User{Login: "am13", Passwort: "12345678"},
		User{Login: "am13", Passwort: "12345678"},
	}
	for i := 0; i != 10; i++ {
		slice = append(slice, slice...)
	}
	for i := 0; i < b.N; i++ {
		result := UniqueOld(slice)
		if len(result) == 0 {
			b.Fatal()
		}
		bend = &result
	}
}

func Benchmark_funcUnique(b *testing.B) {
	var slice = []AllTypes{
		User{Login: "am13", Passwort: "12345678"},
		User{Login: "am134", Passwort: "12345678"},
		User{Login: "am135", Passwort: "12345678"},
		User{Login: "am13", Passwort: "12345678"},
		User{Login: "am13", Passwort: "12345678"},
	}
	for i := 0; i != 10; i++ {
		slice = append(slice, slice...)
	}
	for i := 0; i < b.N; i++ {
		result := Unique(slice)
		if len(result) == 0 {
			b.Fatal()
		}
		bend = &result
	}
}
