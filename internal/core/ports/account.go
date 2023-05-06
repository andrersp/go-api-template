package ports

import (
	"github.com/andrersp/go-api-template/internal/core/domain"
)

type AccountService interface {
	Create(userName, email, password string) error
	Login(string, string) (domain.User, error)
}

type AccountRepository interface {
	Create(domain.User) error
	Login(string) (domain.User, error)
	FindDuplicate(string, string) bool
}
