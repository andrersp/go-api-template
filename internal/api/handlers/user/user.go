package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/andrersp/go-api-template/internal/api/handlers/user/dto"
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

	var response []dto.DtoUserResponse

	users := hu.serviceUser.GetUsers()

	for _, user := range users {
		response = append(response, dto.DtoUserResponse{
			ID:       user.ID,
			UserName: user.UserName,
		})

	}

	json.NewEncoder(w).Encode(response)

	// fmt.Println(users)
	// w.Write([]byte("Usuarios"))
}
