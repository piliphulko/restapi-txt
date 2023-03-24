package httpout

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func sendErrJsonAndLog(err error, w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	Logger.Error("func sendErrJsonAndLog", zap.Error(err))
	e, _ := json.Marshal(map[string]string{"err": err.Error()})
	w.WriteHeader(statusCode)
	w.Write(e)
}

type user struct {
	Login    string `json:"login"`
	Passwort string `json:"passwort"`
}
