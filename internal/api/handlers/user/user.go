package handlers

import (
	"fmt"
	"net/http"

	"github.com/andrersp/go-api-template/internal/api/handlers/user/dto"
	"github.com/andrersp/go-api-template/internal/api/helpers"
	erroresponse "github.com/andrersp/go-api-template/internal/pkg/error-response"

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

func (hu UserHander) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user dto.DtoCreateUserRequest

	err := helpers.DecodeJsonBody(w, r, &user)

	if err != nil {
		helpers.ErrorResponder(422, w, err)
	}

	helpers.SuccessResponder(200, w, user)

	fmt.Println(user)
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
	helpers.SuccessResponder(200, w, response)
}

func (hu UserHander) GetUser(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "userID")

	userID, err := uuid.Parse(paramID)
	if err != nil {
		err := erroresponse.NewErrorResponse("PARAM_ERROR", "")
		helpers.ErrorResponder(400, w, err)
		return
	}

	fmt.Println(userID)

	user, err := hu.serviceUser.GetUser(userID)

	if err != nil {

		err := erroresponse.NewErrorResponse("RECORD_NOT_FOUND", "")
		helpers.ErrorResponder(400, w, err)
		return

	}

	helpers.SuccessResponder(200, w, user)
}
