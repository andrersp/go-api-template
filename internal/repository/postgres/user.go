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

func (ru UserRepository) Get(ID uuid.UUID) (user domain.User, err error) {

	err = ru.db.First(&user, ID).Error

	return
}

func (ru UserRepository) Update(user domain.User) error {

	return nil
}

func (ru UserRepository) GetAll() []domain.User {

	var users []domain.User
	ru.db.Find(&users)

	return users
}
