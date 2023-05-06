package secutiry

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrUnauthorized = errors.New("UNAUTHORIZED")
	ErrForbiden     = errors.New("FORBIDEN")
)

type jwtLoginClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

type TokenData struct {
	UserID string `json:"id"`
}
