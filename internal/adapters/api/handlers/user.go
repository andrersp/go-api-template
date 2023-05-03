package handlers

import (
	"net/http"

	"github.com/andrersp/go-api-template/internal/adapters/api/helpers"
	"github.com/andrersp/go-api-template/internal/core/dto"
	"github.com/andrersp/go-api-template/internal/core/ports"
	customvalidator "github.com/andrersp/go-api-template/internal/pkg/custom-validator"
	secutiry "github.com/andrersp/go-api-template/internal/pkg/security"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserHandler struct {
	serviceUser ports.UserSerice
}

func NewUserHandler(serviceUser ports.UserSerice) UserHandler {

	return UserHandler{
		serviceUser: serviceUser,
	}
}

// CreateUser godoc
// @Summary Create user
// @Description Create a new user
// @Tags Users
// @Param payload body dto.UserRequest true "User payload"
// @Success 201 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /users [post]
func (hu UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user dto.UserRequest
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

	err := hu.serviceUser.Create(user)
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
// @Success 200 {array} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /users [get]
func (hu UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	response := make([]dto.UserResponse, 0)

	users := hu.serviceUser.GetAll()

	for _, user := range users {
		response = append(response, dto.UserResponse{
			ID:       user.ID,
			UserName: user.UserName,
			Email:    user.Email,
		})

	}
	helpers.Responder(200, w, response)
}

// GetUser godoc
// @Summary Get User
// @Description Get user by id
// @Tags Users
// @Param userID path string true "User id" Format(uuid)
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /users/{userID} [get]
func (hu UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "userID")

	errResponse := dto.ErrorResponse{Success: false}
	userResponse := dto.UserResponse{}

	userID, err := uuid.Parse(paramID)
	if err != nil {
		errResponse.Msg = err.Error()
		helpers.Responder(400, w, errResponse)
		return
	}

	user, err := hu.serviceUser.Get(userID)

	if err != nil {
		errResponse.Msg = err.Error()
		helpers.Responder(400, w, errResponse)
		return

	}
	userResponse.Email = user.Email
	userResponse.UserName = user.UserName
	userResponse.ID = user.ID

	secutiry.CreateToken(user.ID)

	helpers.Responder(200, w, userResponse)
}
