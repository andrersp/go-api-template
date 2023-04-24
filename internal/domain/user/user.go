package user

import (
	"net/mail"

	template "github.com/andrersp/go-api-template"

	"github.com/google/uuid"
)

type User struct {
	user *template.User
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

	user := template.User{
		ID:       uuid.New(),
		UserName: userName,
		Email:    email,
		Password: password,
	}

	return User{
		user: &user,
	}, nil

}

func (u *User) GetId() uuid.UUID {
	return u.user.ID
}

func (u *User) GetUserName() string {
	return u.user.UserName
}

func (u *User) GetEmail() string {
	return u.user.Email
}

func (u *User) GetPassword() string {
	return u.user.Password
}

func (u *User) SetId(ID uuid.UUID) {
	if u.user == nil {
		u.user = &template.User{}
	}
	u.user.ID = ID
}

func (u *User) SetUserName(userName string) {
	if u.user == nil {
		u.user = &template.User{}
	}
	u.user.UserName = userName
}

func (u *User) SetEmail(email string) {
	if u.user == nil {
		u.user = &template.User{}
	}

	u.user.Email = email

}
