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
		ID:       u.GetId(),
		UserName: u.GetUserName(),
		Email:    u.GetEmail(),
		Password: u.GetPassword(),
	}
}

func (u UserModel) ToAggregate() user.User {
	user := user.User{}
	user.SetId(u.ID)
	user.SetUserName(u.UserName)
	user.SetEmail(u.Email)
	return user
}
