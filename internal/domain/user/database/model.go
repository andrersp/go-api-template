package database

import (
	"github.com/andrersp/go-api-template/internal/domain/user"
	"github.com/google/uuid"
)

type UserModel struct {
	ID       uuid.UUID `gorm:"primaryKey,index"`
	UserName string    `gorm:"size:40"`
	Email    string    `gorm:"size:120"`
	Password string    `gorm:"size:200"`
}

func (UserModel) TableName() string {
	return "users"
}

func NewFromUser(u user.User) UserModel {
	return UserModel{
		ID:       u.ID,
		UserName: u.UserName,
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u UserModel) ToAggregate() user.User {
	user := user.User{}
	user.ID = u.ID
	user.UserName = u.UserName
	user.Email = u.Email
	return user
}
