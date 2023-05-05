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

	r.Get("/", handlerUser.GetUsers)
	r.Get("/{userID}", handlerUser.GetUser)

}

func RoutersAccount(r chi.Router) {

	connection, err := repository.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewAccountRepository(connection)

	accountService, err := service.NewAccountService(
		service.AccountServiceWithRDB(repository),
	)

	if err != nil {
		log.Fatal(err)
	}

	handlerAcount := handlers.NewAccountHandler(accountService)

	r.Post("/create", handlerAcount.CreateUser)
	r.Post("/login", handlerAcount.Login)
}
