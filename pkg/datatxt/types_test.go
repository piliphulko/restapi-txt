package datatxt

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_typeTestWriteAndReadTable(t *testing.T) {
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
			dataGet, err := DetailTypes[typeTest].ScanType(buf.String())
			assert.Nil(t, err)
			assert.Equal(t, data[i], dataGet)
		})
	}
}

// if diff := cmp.Diff(data[i], getData); diff != "" {
//	t.Error(diff)
//}
