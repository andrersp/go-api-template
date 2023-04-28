package handlers

import (
	"fmt"
	"net/http"

	"github.com/andrersp/go-api-template/internal/api/handlers/user/dto"
	erroresponse "github.com/andrersp/go-api-template/internal/pkg/error-response"
	"github.com/andrersp/go-api-template/internal/pkg/responder"
	service "github.com/andrersp/go-api-template/internal/service/user"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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
	responder.Success(200, w, response)
}

func (hu UserHander) GetUser(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "userID")

	userID, err := uuid.Parse(paramID)
	if err != nil {
		err := erroresponse.NewErrorResponse("PARAM_ERROR")
		responder.Error(400, w, err)
		return
	}

	fmt.Println(userID)

	user, err := hu.serviceUser.GetUser(userID)

	if err != nil {

		err := erroresponse.NewErrorResponse("RECORD_NOT_FOUND")
		responder.Error(400, w, err)
		return

	}

	responder.Success(200, w, user)
}
