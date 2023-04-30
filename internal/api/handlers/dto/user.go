package dto

import (
	"github.com/google/uuid"
)

type DtoUserResponse struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"userName"`
	Email    string    `json:"email"`
}

type DtoUserRequest struct {
	UserName        string `json:"userName" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,containsany=!@#?*$"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}
