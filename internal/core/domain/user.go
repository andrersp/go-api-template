package domain

import (
	"errors"
	"net/mail"

	secutiry "github.com/andrersp/go-api-template/internal/pkg/security"
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

func (u *User) Validate() error {

	if u.UserName == "" {
		return ErrUserNameEmpyt
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return ErrInvalidEmail
	}

	if u.Password == "" || len(u.Password) < 6 {
		return ErrInvalidPassword

	}

	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	u.Password, _ = secutiry.GenerateHash(u.Password)

	return nil

}
