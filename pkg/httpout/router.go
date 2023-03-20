package httpout

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func StartRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Post("/create/account", CreateAccount)
	r.Post("/create/jwt", CreateJWT)

	r.Post("/ok", func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		fmt.Println((buf.String()))
	})

	r.Delete("/delete/account", DeleteAccount)

	r.Get("/error/jwt", ErrorNoJWT)

	r.Route("/abc", func(r chi.Router) {
		r.Use(TakeJWTandLogin)
		r.Get("/date", GetDate)
	})
	return r
}
