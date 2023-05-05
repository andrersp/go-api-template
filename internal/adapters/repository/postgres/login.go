package repository

import (
	"github.com/andrersp/go-api-template/internal/core/domain"
	"github.com/andrersp/go-api-template/internal/core/ports"
	"gorm.io/gorm"
)

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(conn *gorm.DB) ports.LoginRepository {
	return &loginRepository{
		db: conn,
	}
}

func (lr loginRepository) Login(userName string) (user domain.User, err error) {

	err = lr.db.Where("user_name = ?", userName).First(&user).Error
	return

}
