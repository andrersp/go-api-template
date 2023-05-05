package ports

import "github.com/andrersp/go-api-template/internal/core/domain"

type LoginRepository interface {
	Login(string) (domain.User, error)
}

type LoginService interface {
	Login(string, string) (domain.User, error)
}
