package httpout

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/piliphulko/practiceGo/pkg/datalog"
)

func ErrorNoJWT(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(ErrNoJWT.Error()))
}

func CreateJWT(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	var v struct {
		Login    string
		Passwort string
	}
	err := json.NewDecoder(r.Body).Decode(&v)
	fmt.Println(v)
	if err != nil {
		sendErrJson(err, w, http.StatusBadRequest)
		return
	}
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"Login": v.Login})
	e, _ := json.Marshal(map[string]string{"JWT": tokenString})
	w.WriteHeader(http.StatusCreated)
	w.Write(e)
}

func GetDate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	var v struct {
		Login    string
		Passwort string
	}
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		sendErrJson(err, w, http.StatusBadRequest)
		return
	}
	if users.FindValue(datalog.User{Login: v.Login, Passwort: v.Passwort}) {
		sendErrJson(errors.New("such user exists"), w, http.StatusBadRequest)
		return
	}
	if err := addUser(datalog.User{Login: v.Login, Passwort: v.Passwort}); err != nil {
		sendErrJson(err, w, http.StatusBadRequest)
		return
	}
	fmt.Println(v)
	w.WriteHeader(http.StatusCreated)
}
