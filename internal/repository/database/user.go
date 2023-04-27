package database

import (
	"github.com/andrersp/go-api-template/internal/domain/user"
	userDomain "github.com/andrersp/go-api-template/internal/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBUser struct {
	db *gorm.DB
}

func NewDBUser(conn *gorm.DB) user.UserInterface {

	return &DBUser{
		db: conn,
	}
}

func (ru *DBUser) CreateUser(user userDomain.User) error {

	err := ru.db.Model(&userDomain.User{}).Create(user).Error

	return err
}

func (ru *DBUser) GetUser(ID uuid.UUID) (user userDomain.User, err error) {

	err = ru.db.First(&userDomain.User{}, ID).Error

	return
}

func (ru *DBUser) UpdateUser(user user.User) error {

	return nil
}

func (ru *DBUser) GetUsers() []userDomain.User {

	var users []userDomain.User
	ru.db.Find(&users)

	return users
}
