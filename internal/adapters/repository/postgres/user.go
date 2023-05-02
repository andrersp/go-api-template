package repository

import (
	"github.com/andrersp/go-api-template/internal/core/domain"
	"github.com/andrersp/go-api-template/internal/core/ports"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) ports.UserRepository {

	return &UserRepository{
		db: conn,
	}
}

func (ru *UserRepository) Create(user domain.User) error {

	err := ru.db.Model(&domain.User{}).Create(user).Error

	return err
}

func (ru *UserRepository) Get(ID uuid.UUID) (user domain.User, err error) {

	err = ru.db.First(&user, ID).Error

	return
}

func (ru *UserRepository) Update(user domain.User) error {

	return nil
}

func (ru *UserRepository) GetAll() []domain.User {

	var users []domain.User
	ru.db.Find(&users)

	return users
}

func (ru *UserRepository) FindDuplicate(userName, email string) (exist bool) {

	var user domain.User
	err := ru.db.
		Select("user_name", "email").
		Where("user_name = ?", userName).
		Or("email = ?", email).First(&user).Error
	return err == nil

}
