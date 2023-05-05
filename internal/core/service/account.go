package service

import (
	"errors"

	"github.com/andrersp/go-api-template/internal/core/domain"
	"github.com/andrersp/go-api-template/internal/core/dto"
	"github.com/andrersp/go-api-template/internal/core/ports"
	secutiry "github.com/andrersp/go-api-template/internal/pkg/security"
)

var (
	ErrLogin = errors.New("error on username or password")
)

type AccountServiceConfig func(ls *accountService) error

type accountService struct {
	accountRepo ports.AccountRepository
}

func NewAccountService(cfgs ...AccountServiceConfig) (ports.AccountService, error) {

	ls := &accountService{}

	for _, cfg := range cfgs {
		err := cfg(ls)
		if err != nil {
			return nil, err
		}
	}
	return ls, nil
}

func AccountServiceWithRDB(repository ports.AccountRepository) AccountServiceConfig {

	return func(ls *accountService) error {
		ls.accountRepo = repository
		return nil
	}
}

func (as accountService) Create(userRequest dto.UserRequest) (err error) {

	user := domain.User{
		UserName: userRequest.UserName,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
	if err := user.Validate(); err != nil {
		return err
	}

	if ok := as.accountRepo.FindDuplicate(user.UserName, user.Email); ok {
		return errors.New("duplicate username or email")
	}

	err = as.accountRepo.Create(user)
	if err != nil {
		return err
	}

	return
}

func (as accountService) Login(userName, password string) (userResponse domain.User, err error) {

	userResponse, err = as.accountRepo.Login(userName)

	if err != nil {
		err = ErrLogin
		return
	}

	if !secutiry.CheckPasswordHash(userResponse.Password, password) {
		err = ErrLogin
	}
	return
}
