package service

import (
	"errors"

	"github.com/andrersp/go-api-template/internal/core/domain"
	"github.com/andrersp/go-api-template/internal/core/dto"
	"github.com/andrersp/go-api-template/internal/core/ports"

	"github.com/google/uuid"
)

var (
	ErrLogin = errors.New("error on userame or passwor")
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

func (us UserService) Create(userRequest dto.UserRequest) (err error) {

	user := domain.User{
		UserName: userRequest.UserName,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
	if err := user.Validate(); err != nil {
		return err
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

func (us UserService) Get(ID uuid.UUID) (userResponse dto.UserResponse, err error) {

	user, err := us.userRepo.Get(ID)
	if err != nil {
		return
	}

	userResponse.Email = user.Email
	userResponse.UserName = user.UserName
	userResponse.ID = user.ID
	return

}

func (us UserService) GetAll() []dto.UserResponse {

	usersResponde := make([]dto.UserResponse, 0)
	for _, user := range us.userRepo.GetAll() {
		usersResponde = append(usersResponde, dto.UserResponse{
			UserName: user.UserName,
			Email:    user.Email,
			ID:       user.ID,
		})
	}

	return usersResponde

}

func (us *UserService) Update(userRequest dto.UserRequest) error {
	return nil
}
