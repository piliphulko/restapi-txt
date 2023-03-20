package httpout

import (
	"errors"

	"github.com/go-chi/jwtauth/v5"
)

const secretJWT = "secret"

var tokenAuth = jwtauth.New("HS256", []byte(secretJWT), nil)

var (
	ErrNoJWT         = errors.New("Pass jwt token in header\nAuthorization BEARER [jwt]\nif not then create it")
	ErrLoginPasswort = errors.New("login and password is wrong")
)
