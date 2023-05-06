package api

import (
	"log"

	"github.com/andrersp/go-api-template/internal/handlers/api/controllers"
	"github.com/andrersp/go-api-template/internal/handlers/api/middlewares"

	"github.com/andrersp/go-api-template/internal/core/service"
	repository "github.com/andrersp/go-api-template/internal/repository/postgres"

	"github.com/go-chi/chi/v5"
)

func RoutersUser(r chi.Router) {

	r.Use(middlewares.JwtMiddleware())

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

	handlerUser := controllers.NewUserHandler(serviceUser)

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

	handlerAcount := controllers.NewAccountHandler(accountService)

	r.Post("/create", handlerAcount.CreateUser)
	r.Post("/login", handlerAcount.Login)
}
