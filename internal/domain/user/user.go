package user

import (
	"errors"
	"net/mail"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey,index"`
	UserName string    `gorm:"size:40;unique"`
	Email    string    `gorm:"size:120;unique"`
	Password string    `gorm:"size:200"`
}

func (u *User) Validate() error {

	if u.UserName == "" {
		return errors.New("userName cant be empty")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("invalid email")
	}

	if u.Password == "" || len(u.Password) < 6 {
		return errors.New("character number less than 6")

	}

	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	return nil

}
