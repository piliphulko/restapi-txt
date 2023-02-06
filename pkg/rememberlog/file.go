package rememberlog

import (
	"os"
)

func CRWfile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.FileMode(0664))
}
