package handlers

import (
	"net/http"

	"github.com/andrersp/go-api-template/internal/adapters/api/helpers"
	"github.com/andrersp/go-api-template/internal/core/dto"
	"github.com/andrersp/go-api-template/internal/core/ports"
)

type loginHandler struct {
	loginService ports.LoginService
}

func NewLoginHandler(loginService ports.LoginService) loginHandler {

	return loginHandler{
		loginService: loginService,
	}
}

// Login godoc
// @Summary Login
// @Description User Login
// @Tags Login
// @Param Payload body dto.LoginRequest true "Payload login"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /login [post]
func (hl loginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload dto.LoginRequest

	errResponse := dto.ErrorResponse{Success: false}

	if err := helpers.DecodeJsonBody(w, r, &payload); err != nil {
		errResponse.Msg = err.Error()
		helpers.Responder(422, w, errResponse)
		return
	}

	userResponse, err := hl.loginService.Login(payload.UserName, payload.Password)
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
