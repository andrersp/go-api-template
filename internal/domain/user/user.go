package user

import (
	"net/mail"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey,index"`
	UserName string    `gorm:"size:40"`
	Email    string    `gorm:"size:120"`
	Password string    `gorm:"size:200"`
}

func (u *User) Validate() error {

	if u.UserName == "" {
		return ErrEmptyName
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

	return nil

}
