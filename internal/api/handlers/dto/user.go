package dto

import (
	"github.com/google/uuid"
)

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"userName"`
	Email    string    `json:"email"`
}

type UserRequest struct {
	UserName        string `json:"userName" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,containsany=!@#?*$" example:"yourpassword!@#$"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}
