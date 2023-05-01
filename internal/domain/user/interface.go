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

type UserInterface interface {
	Create(User) error
	Get(uuid.UUID) (User, error)
	GetAll() []User
	Update(User) error
	FindDuplicate(string, string) bool
}
