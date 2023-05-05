package api

import (
	"log"

	"github.com/andrersp/go-api-template/internal/adapters/api/handlers"

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

	handlerUser := handlers.NewUserHandler(serviceUser)

	r.Post("/", handlerUser.CreateUser)
	r.Get("/", handlerUser.GetUsers)
	r.Get("/{userID}", handlerUser.GetUser)

}

func RoutersLogin(r chi.Router) {

	connection, err := repository.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewLoginRepository(connection)

	loginService, err := service.NewLoginService(
		service.LoginServiceWithRDB(repository),
	)

	if err != nil {
		log.Fatal(err)
	}

	handlerLogin := handlers.NewLoginHandler(loginService)

	r.Post("/", handlerLogin.Login)
}
