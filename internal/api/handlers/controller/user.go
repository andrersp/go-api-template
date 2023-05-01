package handlers

import (
	"net/http"

	"github.com/andrersp/go-api-template/internal/api/handlers/dto"
	"github.com/andrersp/go-api-template/internal/api/helpers"
	customvalidator "github.com/andrersp/go-api-template/internal/pkg/custom-validator"

	userDomain "github.com/andrersp/go-api-template/internal/domain/user"
	service "github.com/andrersp/go-api-template/internal/service/user"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserController struct {
	serviceUser service.ServiceUser
}

func NewUserController(serviceUser service.ServiceUser) UserController {

	return UserController{
		serviceUser: serviceUser,
	}
}

// CreateUser godoc
// @Summary Create user
// @Description Create a new user
// @Tags Users
// @Param payload body dto.DtoUserRequest true "User payload"
// @Success 201 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /users [post]
func (hu UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user dto.DtoUserRequest
	errResponse := dto.ErrorResponse{Success: false}

	if err := helpers.DecodeJsonBody(w, r, &user); err != nil {
		errResponse.Msg = err.Error()
		helpers.Responder(422, w, errResponse)
		return
	}

	validator := customvalidator.Get()

	if err := validator.Validate(user); err != nil {
		errResponse.Msg = err.Error()
		helpers.Responder(422, w, errResponse)
		return

	}

	userDomain := userDomain.User{
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}

	_, err := hu.serviceUser.CreateUser(userDomain)
	if err != nil {
		errResponse.Msg = err.Error()
		helpers.Responder(422, w, errResponse)
		return
	}

	helpers.Responder(201, w, nil)

}

// Users godoc
// @Summary Get Users
// @Description Get List of users
// @Tags Users
// @Success 200 {array} dto.DtoUserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /users [get]
func (hu UserController) GetUsers(w http.ResponseWriter, r *http.Request) {

	response := make([]dto.DtoUserResponse, 0)

	users := hu.serviceUser.GetUsers()

	for _, user := range users {
		response = append(response, dto.DtoUserResponse{
			ID:       user.ID,
			UserName: user.UserName,
			Email:    user.Email,
		})

	}
	helpers.Responder(200, w, response)
}

func (hu UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "userID")

	errResponse := dto.ErrorResponse{Success: false}

	userID, err := uuid.Parse(paramID)
	if err != nil {
		errResponse.Msg = err.Error()
		helpers.Responder(400, w, errResponse)
		return
	}

	user, err := hu.serviceUser.GetUser(userID)

	if err != nil {
		errResponse.Msg = err.Error()
		helpers.Responder(400, w, errResponse)
		return

	}

	helpers.Responder(200, w, user)
}
