package api

import (
	"log"

	handlers "github.com/andrersp/go-api-template/internal/api/handlers/user"
	"github.com/andrersp/go-api-template/internal/api/helpers"
	service "github.com/andrersp/go-api-template/internal/service/user"
	"github.com/go-chi/chi/v5"
)

func RoutersUser(r chi.Router) {
	serviceUser, err := service.NewUserService(
		service.ServiceWithRDB(),
	)

	if err != nil {
		log.Fatal(err)
	}

	validate := helpers.NewCustomValidator()

	hanndlerUser := handlers.NewUserHandler(serviceUser, validate)

	r.Post("/", hanndlerUser.CreateUser)
	r.Get("/", hanndlerUser.GetUsers)
	r.Get("/{userID}", hanndlerUser.GetUser)
}
