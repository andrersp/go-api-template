package service

import (
	"github.com/andrersp/go-api-template/internal/core/domain"
	"github.com/andrersp/go-api-template/internal/core/ports"

	"github.com/google/uuid"
)

type ServiceUserConfiguration func(us *UserService) error

type UserService struct {
	userRepo ports.UserRepository
}

func NewUserService(cfgs ...ServiceUserConfiguration) (ports.UserSerice, error) {

	us := &UserService{}

	for _, cfg := range cfgs {
		err := cfg(us)
		if err != nil {
			return nil, err
		}
	}

	return us, nil
}

func UserServiceWithRDB(repository ports.UserRepository) ServiceUserConfiguration {
	return func(us *UserService) error {

		us.userRepo = repository

		return nil

	}
}

func (us UserService) Get(ID uuid.UUID) (user domain.User, err error) {

	user, err = us.userRepo.Get(ID)

	return

}

func (us UserService) GetAll() []domain.User {

	return us.userRepo.GetAll()

}

func (us *UserService) Update(userName string) error {
	return nil
}
