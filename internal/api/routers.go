package api

import (
	"log"

	controller "github.com/andrersp/go-api-template/internal/api/handlers/controller"
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

	hanndlerUser := controller.NewUserController(serviceUser)

	r.Post("/", hanndlerUser.CreateUser)
	r.Get("/", hanndlerUser.GetUsers)
	r.Get("/{userID}", hanndlerUser.GetUser)
}
