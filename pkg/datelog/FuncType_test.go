package datelog

import (
	"testing"
)

func TestUnique(t *testing.T) {
	var slice = []AllTypes{
		User{Id: 1, Name: "a"},
		User{Id: 1, Name: "a"},
		User{Id: 1, Name: "b"},
		User{Id: 1, Name: "b"},
		User{Id: 1, Name: "c"},
	}
	if len(Unique(slice)) != len(UniqueNew(slice)) {
		t.Error("UniqueNew non work")
	}
}

var bend *[]AllTypes

func Benchmark_funcUnique(b *testing.B) {
	var slice = []AllTypes{
		User{Id: 1, Name: "a"},
		User{Id: 1, Name: "a"},
		User{Id: 1, Name: "b"},
		User{Id: 1, Name: "b"},
		User{Id: 1, Name: "c"},
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

func Benchmark_funcUniqueNew(b *testing.B) {
	var slice = []AllTypes{
		User{Id: 1, Name: "a"},
		User{Id: 1, Name: "a"},
		User{Id: 1, Name: "b"},
		User{Id: 1, Name: "b"},
		User{Id: 1, Name: "c"},
	}
	for i := 0; i != 10; i++ {
		slice = append(slice, slice...)
	}
	for i := 0; i < b.N; i++ {
		result := UniqueNew(slice)
		if len(result) == 0 {
			b.Fatal()
		}
		bend = &result
	}
}
