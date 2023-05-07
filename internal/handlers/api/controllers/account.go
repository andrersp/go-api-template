package controllers

import (
	"net/http"

	"github.com/andrersp/go-api-template/internal/core/ports"
	"github.com/andrersp/go-api-template/internal/handlers/api/controllers/schemas"
	"github.com/andrersp/go-api-template/internal/handlers/api/helpers"

	apperrors "github.com/andrersp/go-api-template/pkg/app-errors"
	customvalidator "github.com/andrersp/go-api-template/pkg/custom-validator"
	secutiry "github.com/andrersp/go-api-template/pkg/security"
)

type accountHandler struct {
	accountService ports.AccountService
}

func NewAccountHandler(accountService ports.AccountService) accountHandler {

	return accountHandler{
		accountService: accountService,
	}
}

// CreateUser godoc
// @Summary Create user
// @Description Create a new user
// @Tags Account
// @Param payload body schemas.CreateUserRequest true "User payload"
// @Success 201 {object} schemas.SuccessResponse
// @Failure 400 {object} apperrors.AppError
// @Router /account/create [post]
func (ah accountHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user schemas.CreateUserRequest

	if err := helpers.DecodeJsonBody(w, r, &user); err != nil {
		err := apperrors.NewAppError(err.Error())
		helpers.Responder(422, w, err)
		return
	}

	validator := customvalidator.Get()

	if err := validator.Validate(user); err != nil {
		err := apperrors.NewAppError(err.Error())
		helpers.Responder(422, w, err)
		return

	}

	err := ah.accountService.Create(user.UserName, user.Email, user.Password)
	if err != nil {
		err := apperrors.NewAppError(err.Error())
		helpers.Responder(422, w, err)
		return
	}

	helpers.Responder(201, w, nil)

}

// Login godoc
// @Summary Login
// @Description User Login
// @Tags Account
// @Param Payload body schemas.LoginRequest true "Payload login"
// @Success 200 {object} schemas.LoginResponse
// @Failure 400 {object} apperrors.AppError
// @Router /account/login [post]
func (hl accountHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload schemas.LoginRequest

	if err := helpers.DecodeJsonBody(w, r, &payload); err != nil {
		err := apperrors.NewAppError(err.Error())
		helpers.Responder(422, w, err)
		return
	}

	userResponse, err := hl.accountService.Login(payload.UserName, payload.Password)
	if err != nil {
		err := apperrors.NewAppError(err.Error())
		helpers.Responder(422, w, err)
		return
	}

	token, err := secutiry.CreateToken(userResponse.ID)

	if err != nil {
		err := apperrors.NewAppError(err.Error())
		helpers.Responder(500, w, err)
		return
	}
	loginResponse := schemas.LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
	}

	helpers.Responder(200, w, loginResponse)

}
