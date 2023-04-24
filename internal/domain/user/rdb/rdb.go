package rdb

import (
	"github.com/andrersp/go-api-template/internal/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RDBUser struct {
	db *gorm.DB
}

func NewRDBUser(conn *gorm.DB) user.UserRepository {

	return &RDBUser{
		db: conn,
	}
}

func (ru *RDBUser) CreateUser(user user.User) error {

	userModel := NewFromUser(user)

	err := ru.db.Model(&User{}).Create(userModel).Error

	return err
}

func (ru *RDBUser) GetUser(ID uuid.UUID) (user.User, error) {

	var user User

	err := ru.db.First(&user, ID).Error

	return user.ToAggregate(), err
}

func (ru *RDBUser) UpdateUser(user user.User) error {

	return nil
}
