package handlers

import (
	"fmt"
	"net/http"

	service "github.com/andrersp/go-api-template/internal/service/user"
)

type UserHander struct {
	serviceUser service.ServiceUser
}

func NewUserHandler(serviceUser service.ServiceUser) UserHander {

	return UserHander{
		serviceUser: serviceUser,
	}
}

func (hu UserHander) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := hu.serviceUser.GetUsers()

	fmt.Println(users)
	w.Write([]byte("Usuarios"))
}
