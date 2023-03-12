package httpout

import (
	_ "net/http"

	_ "github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

const secretJWT = "secret"

var tokenAuth = jwtauth.New("HS256", []byte(secretJWT), nil)
