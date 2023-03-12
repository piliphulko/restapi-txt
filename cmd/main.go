package main

import (
	"net/http"

	_ "github.com/piliphulko/practiceGo/pkg/datalog"
	"github.com/piliphulko/practiceGo/pkg/httpout"
	_ "github.com/piliphulko/practiceGo/pkg/privacy"
)

func main() {
	http.ListenAndServe(":8080", httpout.StartRouter())
}
