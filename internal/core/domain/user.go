package domain

import (
	"net/mail"

	apperrors "github.com/andrersp/go-api-template/pkg/app-errors"
	secutiry "github.com/andrersp/go-api-template/pkg/security"
	"github.com/google/uuid"
)

var (
	ErrUserNameEmpyt   = "userName cant be empty"
	ErrInvalidEmail    = "invalid email"
	ErrInvalidPassword = "character number less than 6"
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
		err := apperrors.NewAppError("userName cant be empty")
		return user, err
	}

	if _, err := mail.ParseAddress(email); err != nil {
		err = apperrors.NewAppError("invalid email")
		return user, err
	}

	if password == "" || len(password) < 6 {
		err := apperrors.NewAppError("character number less than 6")
		return user, err

	}

	user.ID = uuid.New()
	user.UserName = userName
	user.Email = email

	hashedPassword, err := secutiry.GenerateHash(password)
	if err != nil {
		err = apperrors.NewAppError("error on create hash password")
	}
	user.Password = hashedPassword

	return user, nil

}
