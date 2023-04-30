package handlers

import (
	"fmt"
	"net/http"

	"github.com/andrersp/go-api-template/internal/api/handlers/dto"
	"github.com/andrersp/go-api-template/internal/api/helpers"

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
// @Success 201 {object} helpers.SuccessResponse
// @Failure 400 {object} helpers.ErrorResponse
// @Router /users [post]
func (hu UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

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

	userDomain := userDomain.User{
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}

	_, err = hu.serviceUser.CreateUser(userDomain)
	if err != nil {
		helpers.ErrorResponder(400, w, err)
		return
	}

	helpers.SuccessResponder(201, w, nil)

}

// Users godoc
// @Summary Get Users
// @Description Get List of users
// @Tags Users
// @Success 200 {array} dto.DtoUserResponse
// @Failure 400 {object} helpers.ErrorResponse
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
	helpers.SuccessResponder(200, w, response)
}

func (hu UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "userID")

	userID, err := uuid.Parse(paramID)
	if err != nil {

		helpers.ErrorResponder(400, w, err)
		return
	}

	fmt.Println(userID)

	user, err := hu.serviceUser.GetUser(userID)

	if err != nil {

		helpers.ErrorResponder(400, w, err)
		return

	}

	helpers.SuccessResponder(200, w, user)
}
