package handlers

import (
	"net/http"

	"github.com/andrersp/go-api-template/internal/adapters/api/helpers"
	"github.com/andrersp/go-api-template/internal/core/dto"
	"github.com/andrersp/go-api-template/internal/core/ports"
	customvalidator "github.com/andrersp/go-api-template/internal/pkg/custom-validator"

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

	helpers.Responder(200, w, userResponse)
}

// Login godoc
// @Summary Login
// @Description User Login
// @Tags Login
// @Param Payload body dto.LoginRequest true "Payload login"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /login [post]
func (hu UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload dto.LoginRequest

	errResponse := dto.ErrorResponse{Success: false}

	if err := helpers.DecodeJsonBody(w, r, &payload); err != nil {
		errResponse.Msg = err.Error()
		helpers.Responder(422, w, errResponse)
		return
	}

	userResponse, err := hu.serviceUser.Login(payload.UserName, payload.Password)
	if err != nil {
		errResponse.Msg = err.Error()
		helpers.Responder(422, w, errResponse)
		return
	}

	loginResponse, err := helpers.CreateToken(userResponse.ID)

	if err != nil {
		errResponse.Msg = err.Error()
		helpers.Responder(500, w, errResponse)
		return
	}

	helpers.Responder(200, w, loginResponse)

}
