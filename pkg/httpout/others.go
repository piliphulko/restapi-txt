package httpout

import (
	"encoding/json"
	"net/http"
)

func sendErrJson(err error, w http.ResponseWriter, statusCode int) {
	e, _ := json.Marshal(map[string]string{"err": err.Error()})
	w.WriteHeader(statusCode)
	w.Write(e)
}
