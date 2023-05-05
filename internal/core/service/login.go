package service

import (
	"github.com/andrersp/go-api-template/internal/core/domain"
	"github.com/andrersp/go-api-template/internal/core/ports"
	secutiry "github.com/andrersp/go-api-template/internal/pkg/security"
)

type LoginServiceConfig func(ls *loginService) error

type loginService struct {
	loginRepo ports.LoginRepository
}

func NewLoginService(cfgs ...LoginServiceConfig) (ports.LoginService, error) {

	ls := &loginService{}

	for _, cfg := range cfgs {
		err := cfg(ls)
		if err != nil {
			return nil, err
		}
	}
	return ls, nil
}

func LoginServiceWithRDB(repository ports.LoginRepository) LoginServiceConfig {

	return func(ls *loginService) error {
		ls.loginRepo = repository
		return nil
	}
}

func (ls loginService) Login(userName, password string) (userResponse domain.User, err error) {

	userResponse, err = ls.loginRepo.Login(userName)

	if err != nil {
		err = ErrLogin
		return
	}

	if !secutiry.CheckPasswordHash(userResponse.Password, password) {
		err = ErrLogin
	}
	return
}
