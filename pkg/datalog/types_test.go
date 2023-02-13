package datalog

import (
	"bytes"
	"fmt"
	"testing"

	_ "github.com/google/go-cmp/cmp"
	_ "github.com/stretchr/testify/require"
)

func Test_typeTestWriteAndRead(t *testing.T) {
	var buf = bytes.Buffer{}
	var data = []testType{
		{1, "a", true, 1.15},
		{20, "abc", false, -51.777},
		{-30, "MMM", false, 1000.99},
	}
	for i := range data {
		t.Run("struct test", func(t *testing.T) {
			buf.Reset()
			fmt.Fprint(&buf, data[i])
			_, err := DetailTypes[typeTest].ScanType(buf.String())
			if err != nil {
				t.Error(err)
			}
		})
	}
}

// if diff := cmp.Diff(data[i], getData); diff != "" {
//	t.Error(diff)
//}
