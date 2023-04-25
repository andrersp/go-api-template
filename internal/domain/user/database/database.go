package database

import (
	"github.com/andrersp/go-api-template/internal/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBUser struct {
	db *gorm.DB
}

func NewDBUser(conn *gorm.DB) user.UserRepository {

	return &DBUser{
		db: conn,
	}
}

func (ru *DBUser) CreateUser(user user.User) error {

	userModel := NewFromUser(user)

	err := ru.db.Model(&UserModel{}).Create(userModel).Error

	return err
}

func (ru *DBUser) GetUser(ID uuid.UUID) (user.User, error) {

	var user UserModel

	err := ru.db.First(&user, ID).Error

	return user.ToAggregate(), err
}

func (ru *DBUser) UpdateUser(user user.User) error {

	return nil
}

func (ru *DBUser) GetUsers() (users []user.User) {

	var dbUsers []UserModel
	ru.db.Find(&dbUsers)

	for _, user := range dbUsers {
		users = append(users, user.ToAggregate())
	}
	return
}
