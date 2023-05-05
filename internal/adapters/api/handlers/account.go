package handlers

import (
	"net/http"

	"github.com/andrersp/go-api-template/internal/adapters/api/helpers"
	"github.com/andrersp/go-api-template/internal/core/dto"
	"github.com/andrersp/go-api-template/internal/core/ports"
	customvalidator "github.com/andrersp/go-api-template/internal/pkg/custom-validator"
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
// @Param payload body dto.UserRequest true "User payload"
// @Success 201 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /account/create [post]
func (ah accountHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

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

	err := ah.accountService.Create(user)
	if err != nil {
		errResponse.Msg = err.Error()
		helpers.Responder(422, w, errResponse)
		return
	}

	helpers.Responder(201, w, nil)

}

// Login godoc
// @Summary Login
// @Description User Login
// @Tags Account
// @Param Payload body dto.LoginRequest true "Payload login"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /account/login [post]
func (hl accountHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload dto.LoginRequest

	errResponse := dto.ErrorResponse{Success: false}

	if err := helpers.DecodeJsonBody(w, r, &payload); err != nil {
		errResponse.Msg = err.Error()
		helpers.Responder(422, w, errResponse)
		return
	}

	userResponse, err := hl.accountService.Login(payload.UserName, payload.Password)
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
