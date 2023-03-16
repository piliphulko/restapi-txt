package main

import (
	"net/http"

	_ "github.com/piliphulko/restapi-txt/pkg/datatxt"
	"github.com/piliphulko/restapi-txt/pkg/httpout"
	_ "github.com/piliphulko/restapi-txt/pkg/privacy"
)

func main() {
	http.ListenAndServe(":8080", httpout.StartRouter())
}
