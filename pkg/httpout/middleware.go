package httpout

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

var ErrNoJWT = errors.New("Pass jwt token in header\nAuthorization BEARER [jwt]\nif not then create it")

func TakeJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := jwtauth.VerifyRequest(tokenAuth, r, jwtauth.TokenFromHeader)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/error/jwt", 302)
			next.ServeHTTP(w, r)
		} else {
			claims := token.PrivateClaims()
			userid := claims["Login"]
			fmt.Println(userid)
			next.ServeHTTP(w, r)
		}
	})
}

func TakeLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
