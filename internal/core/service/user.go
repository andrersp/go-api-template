package service

import (
	"errors"

	"github.com/andrersp/go-api-template/internal/core/domain"
	"github.com/andrersp/go-api-template/internal/core/ports"
	secutiry "github.com/andrersp/go-api-template/pkg/security"

	"github.com/google/uuid"
)

var (
	ErrLogin = errors.New("error on username or password")
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

func (us UserService) Create(userName, email, password string) (err error) {

	user, err := domain.NewUser(userName, email, password)
	if err != nil {
		return
	}

	if ok := us.userRepo.FindDuplicate(user.UserName, user.Email); ok {
		return errors.New("duplicate username or email")
	}

	err = us.userRepo.Create(user)
	if err != nil {
		return err
	}

	return
}

func (us UserService) Login(userName, password string) (userResponse domain.User, err error) {

	userResponse, err = us.userRepo.Login(userName)

	if err != nil {
		err = ErrLogin
		return
	}

	if !secutiry.CheckPasswordHash(userResponse.Password, password) {
		err = ErrLogin
	}
	return
}
