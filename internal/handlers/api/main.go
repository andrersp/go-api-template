package api

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/andrersp/go-api-template/docs"
	"github.com/andrersp/go-api-template/internal/config"
	"github.com/andrersp/go-api-template/internal/handlers/api/middlewares"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func StartApiServer() {

	r := chi.NewRouter()
	r.Get("/docs/*", httpSwagger.Handler())

	r.Route("/v1", func(r chi.Router) {
		r.Use(middlewares.LoggerMiddleware())
		r.Use(middlewares.SetHeader("Content-Type", "application/json"))
		r.Route("/account", RoutersAccount)
		r.Route("/users", RoutersUser)

	})

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.API_PORT), r); err != nil {
		log.Fatal(err)

	}

}
