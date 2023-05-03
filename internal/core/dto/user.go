package dto

import "github.com/google/uuid"

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"userName"`
	Email    string    `json:"email"`
}

type UserRequest struct {
	UserName        string `json:"userName" validate:"required" example:"username"`
	Email           string `json:"email" validate:"required,email" example:"myemail@mail.com"`
	Password        string `json:"password" validate:"required,min=8,containsany=!@#?*$" example:"yourpassword!@#$"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	UserName string `json:"userName" validate:"required" example:"username"`
	Password string `json:"password" validate:"required" example:"mypassword"`
}

type TokenData struct {
	UserID string `json:"id"`
}
