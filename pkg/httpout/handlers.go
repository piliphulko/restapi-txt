package httpout

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/piliphulko/restapi-txt/pkg/datatxt"
)

func ErrorNoJWT(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(ErrNoJWT.Error()))
}

func CreateJWT(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	var v user
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		sendErrJsonAndLog(ErrLoginPasswort, w, http.StatusUnauthorized)
		return
	}
	if !users.FindValue(datatxt.User{Login: v.Login, Passwort: v.Passwort}) {
		sendErrJsonAndLog(ErrLoginPasswort, w, http.StatusBadRequest)
		return
	}
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{
		"login": v.Login,
		"exp":   time.Now().Add(1 * time.Minute).Unix()})
	e, _ := json.Marshal(map[string]string{"JWT": tokenString})
	w.WriteHeader(http.StatusCreated)
	w.Write(e)
}

func GetDate(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value("login").(string)
	w.Write([]byte(login))
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	var v user
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		sendErrJsonAndLog(err, w, http.StatusBadRequest)
		return
	}
	if err := addUser(datatxt.User{Login: v.Login, Passwort: v.Passwort}); err != nil {
		sendErrJsonAndLog(err, w, http.StatusBadRequest) // possible error: datatxt.ErrValueExist
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	var v user
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		sendErrJsonAndLog(err, w, http.StatusBadRequest)
		return
	}
	if err := delUser(datatxt.User{Login: v.Login, Passwort: v.Passwort}); err != nil {
		sendErrJsonAndLog(err, w, http.StatusBadRequest) // possible error: datatxt.ErrNoSuchValue
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
