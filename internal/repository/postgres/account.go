package repository

import (
	"github.com/andrersp/go-api-template/internal/core/domain"
	"github.com/andrersp/go-api-template/internal/core/ports"
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(conn *gorm.DB) ports.AccountRepository {
	return &accountRepository{
		db: conn,
	}
}

func (ar accountRepository) Create(user domain.User) error {

	err := ar.db.Model(&domain.User{}).Create(user).Error

	return err
}

func (ar accountRepository) FindDuplicate(userName, email string) (exist bool) {

	var user domain.User
	err := ar.db.
		Select("user_name", "email").
		Where("user_name = ?", userName).
		Or("email = ?", email).First(&user).Error
	return err == nil

}

func (lr accountRepository) Login(userName string) (user domain.User, err error) {

	err = lr.db.Where("user_name = ?", userName).First(&user).Error
	return

}
