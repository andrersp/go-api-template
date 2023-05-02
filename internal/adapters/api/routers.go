package api

import (
	"log"

	userHandler "github.com/andrersp/go-api-template/internal/adapters/api/handlers"
	"github.com/andrersp/go-api-template/internal/core/service"

	"github.com/go-chi/chi/v5"
)

func RoutersUser(r chi.Router) {
	serviceUser, err := service.NewUserService(
		service.UserServiceWithRDB(),
	)

	if err != nil {
		log.Fatal(err)
	}

	hanndlerUser := userHandler.NewUserHandler(serviceUser)

	r.Post("/", hanndlerUser.CreateUser)
	r.Get("/", hanndlerUser.GetUsers)
	r.Get("/{userID}", hanndlerUser.GetUser)
}
