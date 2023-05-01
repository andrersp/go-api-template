package database

import (
	userDomain "github.com/andrersp/go-api-template/internal/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) userDomain.UserInterface {

	return &UserRepository{
		db: conn,
	}
}

func (ru *UserRepository) Create(user userDomain.User) error {

	err := ru.db.Model(&userDomain.User{}).Create(user).Error

	return err
}

func (ru *UserRepository) Get(ID uuid.UUID) (user userDomain.User, err error) {

	err = ru.db.First(&user, ID).Error

	return
}

func (ru *UserRepository) Update(user userDomain.User) error {

	return nil
}

func (ru *UserRepository) GetAll() []userDomain.User {

	var users []userDomain.User
	ru.db.Find(&users)

	return users
}

func (ru *UserRepository) FindDuplicate(userName, email string) (exist bool) {

	var user userDomain.User
	err := ru.db.
		Select("user_name", "email").
		Where("user_name = ?", userName).
		Or("email = ?", email).First(&user).Error
	return err == nil

}
