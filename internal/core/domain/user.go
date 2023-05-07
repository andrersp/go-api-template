package domain

import (
	"errors"
	"net/mail"

	secutiry "github.com/andrersp/go-api-template/pkg/security"
	"github.com/google/uuid"
)

var (
	ErrUserNameEmpyt   = errors.New("userName cant be empty")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("character number less than 6")
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey,index"`
	UserName string    `gorm:"size:40;unique"`
	Email    string    `gorm:"size:120;unique"`
	Password string    `gorm:"size:200"`
}

func NewUser(userName, email, password string) (User, error) {

	user := User{}

	if userName == "" {
		return user, ErrUserNameEmpyt
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return user, ErrInvalidEmail
	}

	if password == "" || len(password) < 6 {
		return user, ErrInvalidPassword

	}

	user.ID = uuid.New()
	user.UserName = userName
	user.Email = email

	user.Password, _ = secutiry.GenerateHash(password)

	return user, nil

}
