package secutiry

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var (
	ErrUnauthorized = errors.New("UNAUTHORIZED")
	ErrForbiden     = errors.New("FORBIDEN")
)

type jwtLoginClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

type TokenData struct {
	UserID uuid.UUID `json:"id"`
}
