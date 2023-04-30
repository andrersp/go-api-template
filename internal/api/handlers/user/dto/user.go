package dto

import (
	"github.com/google/uuid"
)

type DtoUserResponse struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"userName"`
}

type DtoUserRequest struct {
	UserName        string `json:"userName" validate:"required"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
