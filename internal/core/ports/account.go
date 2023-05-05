package ports

import (
	"github.com/andrersp/go-api-template/internal/core/domain"
	"github.com/andrersp/go-api-template/internal/core/dto"
)

type AccountService interface {
	Create(dto.UserRequest) error
	Login(string, string) (domain.User, error)
}

type AccountRepository interface {
	Create(domain.User) error
	Login(string) (domain.User, error)
	FindDuplicate(string, string) bool
}
