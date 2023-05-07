package controllers

import (
	"fmt"
	"net/http"

	"github.com/andrersp/go-api-template/internal/core/ports"
	"github.com/andrersp/go-api-template/internal/handlers/api/controllers/schemas"
	"github.com/andrersp/go-api-template/internal/handlers/api/helpers"
	apperrors "github.com/andrersp/go-api-template/pkg/app-errors"
	secutiry "github.com/andrersp/go-api-template/pkg/security"

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

// Users godoc
// @Summary Get Users
// @Description Get List of users
// @Tags Users
// @Security ApiKeyAuth
// @Success 200 {array} schemas.UserResponse
// @Failure 400 {object} apperrors.AppError
// @Router /users [get]
func (hu UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	tokenData, err := secutiry.GetTokenData(r.Context())
	if err != nil {
		err := apperrors.NewAppError(err.Error())
		helpers.Responder(500, w, err)
		return

	}

	fmt.Println(tokenData.UserID)

	response := make([]schemas.UserResponse, 0)

	users := hu.serviceUser.GetAll()

	for _, user := range users {
		response = append(response, schemas.UserResponse{
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
// @Security ApiKeyAuth
// @Param userID path string true "User id" Format(uuid)
// @Success 200 {object} schemas.UserResponse
// @Failure 400 {object} apperrors.AppError
// @Router /users/{userID} [get]
func (hu UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "userID")

	userResponse := schemas.UserResponse{}

	userID, err := uuid.Parse(paramID)
	if err != nil {
		err := apperrors.NewAppError(err.Error())
		helpers.Responder(400, w, err)
		return
	}

	user, err := hu.serviceUser.Get(userID)

	if err != nil {
		err := apperrors.NewAppError(err.Error())
		helpers.Responder(400, w, err)
		return

	}
	userResponse.Email = user.Email
	userResponse.UserName = user.UserName
	userResponse.ID = user.ID

	helpers.Responder(200, w, userResponse)
}
