package httpout

import (
	"context"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

func TakeJWTandLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := jwtauth.VerifyRequest(tokenAuth, r, jwtauth.TokenFromHeader)
		if err != nil {
			http.Redirect(w, r, "/error/jwt", 302)
		} else {
			claims := token.PrivateClaims()
			login := claims["login"].(string)
			ctx := context.WithValue(r.Context(), "login", login)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
