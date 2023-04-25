package user

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrEmptyName       = errors.New("name cant be empty")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("passowd cant be empty")
	ErrUserNotFound    = errors.New("user not found")
)

type UserRepository interface {
	CreateUser(User) error
	GetUser(uuid.UUID) (User, error)
	GetUsers() []User
	UpdateUser(User) error
}
