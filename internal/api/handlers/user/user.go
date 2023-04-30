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

// CreateUser godoc
// @Summary Create user
// @Description Create a new user
// @Tags Users
// @Param payload body dto.DtoUserRequest true "User payload"
// @Success 200
// @Failure 400 {object} erroresponse.ErrorResponse
// @Router /users [post]
func (hu UserHander) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user dto.DtoUserRequest

	err := helpers.DecodeJsonBody(w, r, &user)

	if err != nil {
		helpers.ErrorResponder(422, w, err)
	}

	validator := helpers.NewValidator()

	if err = validator.Validate(user); err != nil {
		helpers.ErrorResponder(400, w, err)
		return

	}

	helpers.SuccessResponder(200, w, user)

	fmt.Println(user)
}

// Users godoc
// @Summary Get Users
// @Description Get List of users
// @Tags Users
// @Success 200 {array} dto.DtoUserResponse
// @Failure 400 {object} erroresponse.ErrorResponse
// @Router /users [get]
func (hu UserHander) GetUsers(w http.ResponseWriter, r *http.Request) {

	response := make([]dto.DtoUserResponse, 0)

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
