package user

import (
	"net/mail"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	UserName string
	Email    string
	Password string
}

func CreateNewUser(userName, email, password string) (User, error) {

	if userName == "" {
		return User{}, ErrEmptyName
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return User{}, ErrInvalidEmail
	}

	if password == "" || len(password) < 6 {
		return User{}, ErrInvalidPassword

	}

	user := User{
		ID:       uuid.New(),
		UserName: userName,
		Email:    email,
		Password: password,
	}

	return user, nil

}
