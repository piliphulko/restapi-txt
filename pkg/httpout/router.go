package httpout

import (
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

	r.Get("/error/jwt", ErrorNoJWT)

	r.Route("/{Login}", func(r chi.Router) {
		r.Use(TakeJWT)
		r.Get("/", GetDate)
	})
	return r
}
