package api

import (
	"log"

	userHandler "github.com/andrersp/go-api-template/internal/adapters/api/handlers"

	repository "github.com/andrersp/go-api-template/internal/adapters/repository/postgres"
	"github.com/andrersp/go-api-template/internal/core/service"

	"github.com/go-chi/chi/v5"
)

func RoutersUser(r chi.Router) {

	connection, err := repository.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewUserRepository(connection)

	serviceUser, err := service.NewUserService(
		service.UserServiceWithRDB(repository),
	)

	if err != nil {
		log.Fatal(err)
	}

	handlerUser := userHandler.NewUserHandler(serviceUser)

	r.Post("/", handlerUser.CreateUser)
	r.Get("/", handlerUser.GetUsers)
	r.Get("/{userID}", handlerUser.GetUser)
	r.Post("/login", handlerUser.Login)
}
