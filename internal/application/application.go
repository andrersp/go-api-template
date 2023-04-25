package application

import (
	"log"
	"net/http"

	"github.com/andrersp/go-api-template/internal/application/routers"
	"github.com/andrersp/go-api-template/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func init() {
	config.SetConfig()
	err := config.CreateSQLiteConn()
	if err != nil {
		log.Fatal(err)
	}
}

func StartServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)

	r.Mount("/users", routers.RoutersUser())

	http.ListenAndServe(config.API_PORT, r)

}
