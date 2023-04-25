package routers

import (
	"log"

	handlers "github.com/andrersp/go-api-template/internal/handlers/user"
	service "github.com/andrersp/go-api-template/internal/service/user"
	"github.com/go-chi/chi/v5"
)

func RoutersUser() chi.Router {

	serviceUser, err := service.NewUserService(
		service.ServiceWithRDB(),
	)

	if err != nil {
		log.Fatal(err)
	}

	hanndlerUser := handlers.NewUserHandler(*serviceUser)
	r := chi.NewRouter()

	r.Get("/", hanndlerUser.GetUsers)
	return r

}
