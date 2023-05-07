package ports

import (
	"github.com/andrersp/go-api-template/internal/core/domain"
	"github.com/google/uuid"
)

type UserSerice interface {
	Get(uuid.UUID) (domain.User, error)
	GetAll() []domain.User
	Update(username string) error
	Create(userName, email, password string) error
	Login(string, string) (domain.User, error)
}

type UserRepository interface {
	Get(uuid.UUID) (domain.User, error)
	GetAll() []domain.User
	Update(domain.User) error
	Create(domain.User) error
	Login(string) (domain.User, error)
	FindDuplicate(string, string) bool
}
