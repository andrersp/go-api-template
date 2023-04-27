package api

import (
	"net/http"

	"github.com/andrersp/go-api-template/internal/api/middlewares"
	"github.com/andrersp/go-api-template/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartApiServer() {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(middlewares.SetHeader("Content-Type", "application/json"))
		r.Route("/users", RoutersUser)
	})

	http.ListenAndServe(config.API_PORT, r)

}
